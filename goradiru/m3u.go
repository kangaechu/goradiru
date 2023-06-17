package goradiru

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/sys/unix"
)

func downloadEpisode(episode *Episode) (err error) {
	config := GetConfig()

	m3u8URL := episode.URL

	progDir := config.ProgDir
	fileType := config.FileType

	metadata := generateMetadata(episode)

	// 書き込み先ディレクトリの確認
	episodePath := fmtFileName(episode, progDir, fileType)
	if !isWritableDir(episodePath) {
		err := os.MkdirAll(progDir, 0755)
		if err != nil {
			return err
		}
	}
	if fileType == "m4a" {
		err = convertM3u8ToM4A(m3u8URL, episodePath, metadata)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Not implemented such file type:" + fileType)
	}
	return nil
}

func convertM3u8ToM4A(masterM3u8Path string, filename string, metadata []string) error {
	f, err := newFFMPEG()
	if err != nil {
		return err
	}

	f.setArgs(
		"-y", "-http_seekable", "0",
		"-i", masterM3u8Path,
		"-absf", "aac_adtstoasc",
		"-acodec", "copy",
	)

	f.setArgs(metadata...)

	msg, err := f.execute(filename)
	if err != nil {
		fmt.Println(string(msg))
		return err
	}
	return nil
}

// 　出力ファイル名のフルパスを返す
func fmtFileName(episode *Episode, baseDir string, fileType string) string {
	filename := fmtTitle(episode) + "." + fileType
	return filepath.Join(baseDir, filename)
}

// ファイル・ディレクトリの存在・書き込み確認
func isWritableDir(filename string) bool {
	stat, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if !stat.IsDir() {
		return false
	}
	err = unix.Access(filename, unix.W_OK)
	if err != nil {
		return false
	}
	return true
}

func generateMetadata(episode *Episode) (metadata []string) {
	title := fmtTitle(episode)
	metadata = []string{
		"-metadata", "title=" + title,
		"-metadata", "genre=radio",
		"-metadata", "artist=NHK",
	}
	return metadata
}

func fmtTitle(episode *Episode) string {
	var title string
	// program series episode
	// 全部違う : こども科学電話相談
	// seriesがnull, programがepisodeに含まれている : カルチャーラジオ 東京03
	// ProgramとSeriesが同じ場合は2回出力しない
	programTitle := episode.Program.Title
	seriesTitle := episode.Series.Title
	episodeTitle := episode.Title
	episodeDate := episode.Start.Format("20060102-1504")

	series := episode.Series.Title
	if series == "" {
		if programTitle == episodeTitle {
			title = episodeTitle + " " + episodeDate
		} else if strings.Contains(episodeTitle, programTitle) {
			title = episodeTitle
		} else {
			title = programTitle + " " + episodeTitle
		}
	} else {
		title = programTitle + " " + seriesTitle + " " + episodeTitle
	}
	title = strings.Replace(title, "　", " ", -1)
	title = strings.Replace(title, "【", " ", -1)
	title = strings.Replace(title, "】", " ", -1)
	title = strings.Replace(title, "“", "", -1)
	title = strings.Replace(title, "”", "", -1)
	trimSpace := regexp.MustCompile(`\s+`)
	title = trimSpace.ReplaceAllString(title, " ")
	title = strings.TrimSpace(title)
	return title
}
