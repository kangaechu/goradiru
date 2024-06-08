package goradiru

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"
)

// ProgramJSON はJSONパース用の型
type ProgramJSON struct {
	ID                   int       `json:"id"`
	Title                string    `json:"title"`
	RadioBroadcast       string    `json:"radio_broadcast"`
	Schedule             string    `json:"schedule"`
	CornerName           string    `json:"corner_name"`
	ThumbnailURL         string    `json:"thumbnail_url"`
	SeriesDescription    string    `json:"series_description"`
	SeriesURL            string    `json:"series_url"`
	ShareTextTitle       string    `json:"share_text_title"`
	ShareTextURL         string    `json:"share_text_url"`
	ShareTextDescription string    `json:"share_text_description"`
	Episodes             []Episode `json:"episodes"`
}

// Programs は番組の一覧
type Programs []Program

// Program は番組
type Program struct {
	ID       string    // 番組ID
	Title    string    // 番組名
	Episodes []Episode // エピソード
}

// Download program
func (p Program) Download(dps *DownloadedPrograms) (err error) {
	cpus := runtime.NumCPU() // CPUの数
	log.Println("Parallels: ", cpus)
	semaphore := make(chan int, cpus)
	for _, episode := range p.Episodes {
		var wg sync.WaitGroup
		err := episode.download(p.ID, p.Title, &wg, &semaphore, dps)
		if err != nil {
			return err
		}
		wg.Wait()
	}

	return nil
}

// Episode 番組のエピソード（各話）
type Episode struct {
	ID              int    `json:"id"`
	ProgramTitle    string `json:"program_title"`
	OnairDate       string `json:"onair_date"`
	ClosedAt        string `json:"closed_at"`
	StreamURL       string `json:"stream_url"`
	AaContentsID    string `json:"aa_contents_id"`
	ProgramSubTitle string `json:"program_sub_title"`
}

// Episodeをダウンロード
func (e *Episode) download(programID string, programTitle string, _ *sync.WaitGroup, semaphore *chan int, dps *DownloadedPrograms) (err error) {
	if dps.isAlreadyDownloaded(e) {
		log.Printf("download skipped %s", fmtTitle(e))
	} else {
		*semaphore <- 1
		log.Printf("download started %s", fmtTitle(e))

		err = downloadEpisode(e)
		if err != nil {
			return err
		}
		dps.addDownloadedEpisode(e, programID, programTitle)
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

	program.ID = strconv.Itoa(programJSON.ID)
	program.Title = programJSON.Title
	program.Episodes = programJSON.Episodes

	return program, nil
}

// CreateProgram は番組情報をダウンロードし、Programを生成します
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
