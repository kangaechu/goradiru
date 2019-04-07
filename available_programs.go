package goradiru

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AvailableProgramJson struct {
	AvailablePrograms []struct {
		SiteID      string `json:"site_id"`
		ProgramName string `json:"program_name"`
		MediaCode   string `json:"media_code"`
		CornerID    string `json:"corner_id"`
		CornerName  string `json:"corner_name"`
		ThumbnailP  string `json:"thumbnail_p"`
		ThumbnailC  string `json:"thumbnail_c"`
		DetailJSON  string `json:"detail_json"`
		OpenTime    string `json:"open_time"`
		CloseTime   string `json:"close_time"`
		OnairDate   string `json:"onair_date"`
		LinkURL     string `json:"link_url"`
	} `json:"data_list"`
}

func getAvailablePrograms() error {
	indexUrl := "https://www.nhk.or.jp/radioondemand/json/index/index.json"
	availableProgramJson := AvailableProgramJson{}

	res, err := http.Get(indexUrl)
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

	if err := json.Unmarshal(jsonBytes, &availableProgramJson); err != nil {
		return err
	}

	for _, p := range availableProgramJson.AvailablePrograms {
		fmt.Println("  - Name:", p.ProgramName)
		fmt.Println("    Url:", p.DetailJSON)
	}
	return nil
}
