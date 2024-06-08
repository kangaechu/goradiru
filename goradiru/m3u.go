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

	m3u8URL := episode.StreamURL

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
	programTitle := episode.ProgramTitle
	episodeDate := episode.OnairDate

	title = programTitle + " " + episodeDate
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
