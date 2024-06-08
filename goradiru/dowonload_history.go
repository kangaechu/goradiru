package goradiru

import (
	"io"
	"os"
	"sort"
	"strconv"

	"gopkg.in/yaml.v2"
)

type DownloadedPrograms []DownloadedProgram

type DownloadedProgram struct {
	ProgramID    string `yaml:"ProgramID"`
	ProgramTitle string `yaml:"ProgramTitle"`
	EpisodeID    string `yaml:"EpisodeID"`
	EpisodeTitle string `yaml:"EpisodeTitle"`
}

func LoadDownloadedPrograms(downloadedHistoryConfFile string) (dps *DownloadedPrograms) {
	var err error
	file, err := os.Open(downloadedHistoryConfFile)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("error on closing downloaded history")
		}
	}(file)
	if err != nil {
		// ファイルがない場合は空のDownloadedProgramsを返す
		dps = new(DownloadedPrograms)
		return dps
	}
	readBytes, err := io.ReadAll(file)
	if err != nil {
		panic("error on reading downloaded history")
	}
	err = yaml.Unmarshal(readBytes, &dps)
	if err != nil {
		panic("error on reading downloaded history")
	}
	return dps
}

func (dps DownloadedPrograms) Save() error {
	config := GetConfig()

	file, err := os.OpenFile(config.DownloadedHistoryConfFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	sort.Sort(ByEpisodeID(dps))
	bytesYaml, err := yaml.Marshal(dps)
	if err != nil {
		return err
	}
	_, err = file.Write(bytesYaml)
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil

}

// すでにダウンロードされたものか確認する
func (dps DownloadedPrograms) isAlreadyDownloaded(episode *Episode) bool {
	for _, dp := range dps {
		if dp.EpisodeID == strconv.Itoa(episode.ID) {
			return true
		}
	}
	return false
}

// Downloadされたものに追加する
func (dps DownloadedPrograms) addDownloadedEpisode(episode *Episode, programID string, programTitle string) {
	dps = append(dps, DownloadedProgram{
		programID,
		programTitle,
		strconv.Itoa(episode.ID),
		episode.ProgramTitle,
	})
}

type ByEpisodeID DownloadedPrograms

func (a ByEpisodeID) Len() int      { return len(a) }
func (a ByEpisodeID) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByEpisodeID) Less(i, j int) bool {
	if a[i].ProgramID < a[j].ProgramID {
		return true
	}
	if a[i].ProgramID > a[j].ProgramID {
		return false
	}
	return a[i].EpisodeID < a[j].EpisodeID
}
