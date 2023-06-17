package goradiru

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
)

// ProgramJSON はJSONパース用の型
type ProgramJSON struct {
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

// Programs は番組の一覧
type Programs []Program

// Program は番組
type Program struct {
	ID     string   // 番組ID
	Title  string   // 番組名
	Series []Series // シリーズ
}

// Download program
func (p Program) Download(dps *DownloadedPrograms) (err error) {
	cpus := runtime.NumCPU() // CPUの数
	log.Println("Parallels: ", cpus)
	semaphore := make(chan int, cpus)
	for _, series := range p.Series {
		var wg sync.WaitGroup
		err := series.download(&wg, &semaphore, dps)
		if err != nil {
			return err
		}
		wg.Wait()
	}

	return nil
}

// Series は 番組のシリーズ
type Series struct {
	ID       string    // ID
	Title    string    // シリーズ名
	SubTitle string    // サブタイトル
	Episodes []Episode // エピソード
}

// Seriesをダウンロード
func (s Series) download(wg *sync.WaitGroup, semaphore *chan int, dps *DownloadedPrograms) (err error) {
	for _, episode := range s.Episodes {
		wg.Add(1)

		go func(ep Episode) {
			defer wg.Done()
			err := ep.download(wg, semaphore, dps)
			if err != nil {
				log.Fatal(err)
			}

		}(episode)
		time.Sleep(100 * time.Millisecond) // 順序よく並ぶように入れている
	}
	return nil
}

// Episode シリーズのエピソード（各話）
type Episode struct {
	ID       string    // ID
	Title    string    // エピソード名
	SubTitle string    // サブタイトル
	URL      string    // url
	Station  string    // 放送局 (aa_vinfo2の , より前)
	Start    time.Time // 開始時刻 (aa_vinfo4 の _ より前)
	End      time.Time // 終了時刻 (aa_vinfo4 の _ より後)
	Program  *Program
	Series   *Series
}

// Episodeをダウンロード
func (e *Episode) download(_ *sync.WaitGroup, semaphore *chan int, dps *DownloadedPrograms) (err error) {
	if dps.isAlreadyDownloaded(e) {
		log.Printf("download skipped %s", fmtTitle(e))
	} else {
		*semaphore <- 1
		log.Printf("download started %s", fmtTitle(e))

		err = downloadEpisode(e)
		if err != nil {
			return err
		}
		dps.addDownloadedEpisode(e)
		<-*semaphore
		log.Printf("download completed %s", fmtTitle(e))
	}
	return nil
}

// download は番組情報をダウンロードします
func download(programURL string) (jsonBytes []byte, err error) {
	res, err := http.Get(programURL) // nolint: gosec
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err != nil {
			err = res.Body.Close()
		}
	}()
	jsonBytes, err = io.ReadAll(res.Body)
	if err != nil {
		return jsonBytes, err
	}
	return jsonBytes, nil
}

// 番組情報のJSONからProgramを生成
func createProgramFromJSONBytes(jsonBytes []byte) (program Program, err error) {
	var programJSON ProgramJSON

	// JSON parse
	if err := json.Unmarshal(jsonBytes, &programJSON); err != nil {
		return program, err
	}

	series := make([]Series, len(programJSON.Main.DetailList))
	for si, s := range programJSON.Main.DetailList {

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
	program.ID = programJSON.Main.SiteID
	program.Title = programJSON.Main.ProgramName
	program.Series = series

	return program, nil
}

func CreateProgram(programURL string) (program Program, err error) {
	jsonBytes, err := download(programURL)
	if err != nil {
		return program, err
	}
	program, err = createProgramFromJSONBytes(jsonBytes)
	if err != nil {
		return program, err
	}
	return program, nil
}
