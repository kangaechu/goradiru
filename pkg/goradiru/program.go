package goradiru

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// ProgramJson はJSONパース用の型
type ProgramJson struct {
	Main struct {
		SiteID      string `json:"site_id"`
		ProgramName string `json:"program_name"`
		DetailList  []struct {
			HeadlineID  string `json:"headline_id"`
			Headline    string `json:"headline"`
			HeadlineSub string `json:"headline_sub"`
			FileList    []struct {
				FileID       string `json:"file_id"`
				FileTitle    string `json:"file_title"`
				FileTitleSub string `json:"file_title_sub"`
				FileName     string `json:"file_name"`
				AaVinfo2     string `json:"aa_vinfo2"`
				AaVinfo4     string `json:"aa_vinfo4"`
			} `json:"file_list"`
		} `json:"detail_list"`
	} `json:"main"`
}

// Programのリスト
type Programs []Program

// Program は番組
type Program struct {
	Id     string   // 番組ID
	Title  string   // 番組名
	Series []Series // シリーズ
}

// Programをダウンロード
func (p Program) Download() (err error) {
	for _, series := range p.Series {
		err := series.download()
		if err != nil {
			return err
		}
	}
	return nil
}

// Series は 番組のシリーズ
type Series struct {
	Id       string    // ID
	Title    string    // シリーズ名
	SubTitle string    // サブタイトル
	Episodes []Episode // エピソード
}

// Seriesをダウンロード
func (s Series) download() (err error) {
	for _, episode := range s.Episodes {
		err := episode.download()
		if err != nil {
			return err
		}
	}
	return nil
}

// Episode シリーズのエピソード（各話）
type Episode struct {
	Id       string    // ID
	Title    string    // エピソード名
	SubTitle string    // サブタイトル
	Url      string    // url
	Station  string    // 放送局 (aa_vinfo2の , より前)
	Start    time.Time // 開始時刻 (aa_vinfo4 の _ より前)
	End      time.Time // 終了時刻 (aa_vinfo4 の _ より後)
	Program  *Program
	Series   *Series
}

// Episodeをダウンロード
func (e *Episode) download() (err error) {
	dp := GetDownloadedPrograms()
	if dp.isAlreadyDownloaded(e) {
		log.Printf("download skipped %s_%s_%s", e.Program.Title, e.Series.Title, e.Title)
	} else {
		log.Printf("download started %s_%s_%s", e.Program.Title, e.Series.Title, e.Title)

		err = downloadEpisode(e)
		if err != nil {
			return err
		}
		log.Printf("download completed %s_%s_%s", e.Program.Title, e.Series.Title, e.Title)
		err = dp.addDownloadedEpisode(e)
		if err != nil {
			return err
		}
	}
	return nil
}

// download は番組情報をダウンロードします
func download(programUrl string) (jsonBytes []byte, err error) {
	res, err := http.Get(programUrl) // nolint: gosec
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err != nil {
			err = res.Body.Close()
		}
	}()
	jsonBytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return jsonBytes, err
	}
	return jsonBytes, nil
}

// 番組情報のJSONからProgramを生成
func createProgramFromJsonBytes(jsonBytes []byte) (program Program, err error) {
	var programJson ProgramJson

	// JSON parse
	if err := json.Unmarshal(jsonBytes, &programJson); err != nil {
		return program, err
	}

	series := make([]Series, len(programJson.Main.DetailList))
	for si, s := range programJson.Main.DetailList {

		var episodes = make([]Episode, len(s.FileList))
		for ei, e := range s.FileList {
			// set Station
			station := strings.Split(e.AaVinfo2, ",")[0]
			// set Start/End
			var start, end time.Time
			if e.AaVinfo4 != "9999-99-99T99:99:99+09:00_9999-99-99T99:99:99+09:00" {
				startEnd := strings.Split(e.AaVinfo4, "_")
				timeFmtStr := "2006-01-02T15:04:05-07:00"
				start, err = time.Parse(timeFmtStr, startEnd[0])
				if err != nil {
					return program, err
				}
				end, err = time.Parse(timeFmtStr, startEnd[1])
				if err != nil {
					return program, err
				}
			}
			episodes[ei] = Episode{e.FileID, e.FileTitle, e.FileTitleSub, e.FileName, station, start, end, &program, &series[si]}
		}
		series[si] = Series{s.HeadlineID, s.Headline, s.HeadlineSub, episodes}
	}
	program.Id = programJson.Main.SiteID
	program.Title = programJson.Main.ProgramName
	program.Series = series

	return program, nil
}

func CreateProgram(programUrl string) (program Program, err error) {
	jsonBytes, err := download(programUrl)
	if err != nil {
		return program, err
	}
	program, err = createProgramFromJsonBytes(jsonBytes)
	if err != nil {
		return program, err
	}
	return program, nil
}
