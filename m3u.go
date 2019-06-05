package goradiru

import (
	"errors"
	"github.com/grafov/m3u8"
	"golang.org/x/sys/unix"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func downloadEpisode(episode *Episode) (err error) {
	config := GetConfig()

	m3u8Url := getM3u8MasterPlaylist(episode.Url)

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
		err = convertM3u8ToM4A(m3u8Url, episodePath, metadata)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Not implemented such file type:" + fileType)
	}
	return nil
}

func convertM3u8ToM4A(masterM3u8Path string, filename string, metadata []string) error {
	f, err := newFFMPEG(masterM3u8Path)
	if err != nil {
		return err
	}

	f.setArgs(
		"-protocol_whitelist", "file,crypto,http,https,tcp,tls",
		"-movflags", "faststart",
		"-c", "copy",
		"-y",
		"-bsf:a", "aac_adtstoasc",
	)

	f.setArgs(metadata...)

	_, err = f.execute(filename)
	if err != nil {
		return err
	}
	return nil
}

//func convertM4AToMP3(m4apath string, title string) error {
//	f, err := newFFMPEG(m4apath)
//	if err != nil {
//		return err
//	}
//
//	f.setArgs(
//		"-y",
//		"-acodec", "libmp3lame",
//		"-ab", "256k",
//	)
//
//	var name = title + ".mp3"
//	fmt.Println(name)
//
//	_, err = f.execute(name)
//	return err
//}

func getM3u8MasterPlaylist(m3u8FilePath string) string {
	resp, err := http.Get(m3u8FilePath) // nolint: gosec
	if err != nil {
		log.Fatal(err)
	}
	f := resp.Body

	p, t, err := m3u8.DecodeFrom(f, true)
	if err != nil {
		log.Fatal(err)
	}

	if t != m3u8.MASTER {
		log.Fatalf("not support file type [%v]", t)
	}

	return p.(*m3u8.MasterPlaylist).Variants[0].URI
}

//　出力ファイル名のフルパスを返す
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
	// ProgramとSeriesが同じ場合は2回出力しない
	series := episode.Series.Title
	if episode.Program.Title == episode.Series.Title {
		series = ""
	}
	if series == "" {
		title = episode.Program.Title + "_" + episode.Title
	} else {
		title = episode.Program.Title + "_" + episode.Series.Title + "_" + episode.Title
	}
	title = strings.Replace(title, "【", " ", -1)
	title = strings.Replace(title, "】", " ", -1)
	title = strings.Replace(title, "“", "", -1)
	title = strings.Replace(title, "”", "", -1)
	title = strings.TrimSpace(title)
	return title
}
