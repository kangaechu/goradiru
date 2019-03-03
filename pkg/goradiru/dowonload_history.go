package goradiru

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// シングルトン
var sharedDownloaded = &DownloadedPrograms{ /* 初期化 */ }

func GetDownloadedPrograms() (dps *DownloadedPrograms) {
	if dps == nil {
		err := Load()
		if err != nil {
			panic("error on reading downloaded history")
		}
	}
	return sharedDownloaded
}

func Load() (err error) {
	config := GetConfig()
	dps, err := loadYaml(config.DownloadedHistoryConfFile)
	if err != nil {
		return err
	}
	sharedDownloaded = &dps
	return nil
}

type DownloadedPrograms []DownloadedProgram

type DownloadedProgram struct {
	ProgramID    string `yaml:"ProgramID"`
	ProgramTitle string `yaml:"ProgramTitle"`
	EpisodeID    string `yaml:"EpisodeID"`
	EpisodeTitle string `yaml:"EpisodeTitle"`
}

func loadYaml(configFile string) (dps DownloadedPrograms, err error) {
	file, err := os.Open(configFile)
	defer file.Close()
	if err != nil {
		// ファイルがない場合は空のDownloadedProgramsを返す
		return dps, nil
	}
	readBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return dps, err
	}
	err = yaml.Unmarshal(readBytes, &dps)
	if err != nil {
		return dps, err
	}
	return dps, nil

}

// すでにダウンロードされたものか確認する
func (dps DownloadedPrograms) isAlreadyDownloaded(episode *Episode) bool {
	for _, dp := range dps {
		if dp.EpisodeID == episode.Id {
			return true
		}
	}
	return false
}

// Downloadされたものに追加する
func (dps *DownloadedPrograms) addDownloadedEpisode(episode *Episode) (err error) {
	config := GetConfig()
	*dps = append(*dps, DownloadedProgram{episode.Program.Id, episode.Program.Title, episode.Id, episode.Title})
	// open yaml file for write
	file, err := os.OpenFile(config.DownloadedHistoryConfFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	bytesYaml, err := yaml.Marshal(dps)
	if err != nil {
		return err
	}
	_, err = file.Write(bytesYaml)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}
