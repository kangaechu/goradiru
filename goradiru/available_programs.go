package goradiru

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AvailableProgramJSON struct {
	AvailablePrograms []struct {
		SiteID          string `json:"site_id"`
		ProgramName     string `json:"program_name"`
		ProgramNameKana string `json:"program_name_kana"`
		MediaCode       string `json:"media_code"`
		CornerID        string `json:"corner_id"`
		CornerName      string `json:"corner_name"`
		ThumbnailP      string `json:"thumbnail_p"`
		ThumbnailC      string `json:"thumbnail_c"`
		OpenTime        string `json:"open_time"`
		CloseTime       string `json:"close_time"`
		OnairDate       string `json:"onair_date"`
		LinkURL         string `json:"link_url"`
		StartTime       string `json:"start_time"`
		UpdateTime      string `json:"update_time"`
		Dev             string `json:"dev"`
		DetailJSON      string `json:"detail_json"`
	} `json:"data_list"`
}

func getAvailablePrograms() error {
	indexURL := "https://www.nhk.or.jp/radioondemand/json/index_v3/index.json"
	availableProgramJSON := AvailableProgramJSON{}

	res, err := http.Get(indexURL)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err != nil {
			err = res.Body.Close()
		}
	}()
	jsonBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonBytes, &availableProgramJSON); err != nil {
		return err
	}

	for _, p := range availableProgramJSON.AvailablePrograms {
		fmt.Println("  - Name:", p.ProgramName)
		fmt.Println("    URL:", p.DetailJSON)
	}
	return nil
}
