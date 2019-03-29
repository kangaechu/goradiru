package goradiru

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type DownloadedPrograms []DownloadedProgram

type DownloadedProgram struct {
	ProgramID    string `yaml:"ProgramID"`
	ProgramTitle string `yaml:"ProgramTitle"`
	EpisodeID    string `yaml:"EpisodeID"`
	EpisodeTitle string `yaml:"EpisodeTitle"`
}

func LoadDownloadedPrograms(downloadedHistoryConfFile string) (dps *DownloadedPrograms) {
	if dps == nil {
		var err error
		file, err := os.Open(downloadedHistoryConfFile)
		defer file.Close()
		if err != nil {
			// ファイルがない場合は空のDownloadedProgramsを返す
			dps = new(DownloadedPrograms)
			return dps
		}
		readBytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic("error on reading downloaded history")
		}
		err = yaml.Unmarshal(readBytes, &dps)
		if err != nil {
			panic("error on reading downloaded history")
		}
		log.Println("READ unlocked")
	}
	return dps
}

func (dps DownloadedPrograms) Save() error {
	config := GetConfig()

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
		if dp.EpisodeID == episode.Id {
			return true
		}
	}
	return false
}

// Downloadされたものに追加する
func (dps *DownloadedPrograms) addDownloadedEpisode(episode *Episode) {
	*dps = append(*dps, DownloadedProgram{
		episode.Program.Id,
		episode.Program.Title,
		episode.Id,
		episode.Title,
	})
}
