package goradiru

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AvailableProgram struct {
	ID             int    `json:"id"` // 458
	Title          string `json:"title"`
	RadioBroadcast string `json:"radio_broadcast"` // "R1" / "R2" / "FM"
	CornerName     string `json:"corner_name"`
	OnAirDate      string `json:"onair_date"` // 最新の放送日 (2024年6月8日(土)放送)
	ThumbnailURL   string `json:"thumbnail_url"`
	SeriesSiteID   string `json:"series_site_id"` // 2769
	CornerSiteID   string `json:"corner_site_id"` // 01
}

type AvailableProgramJSON struct {
	Corners []AvailableProgram `json:"corners"`
}

func getAvailablePrograms() error {
	indexURL := "https://www.nhk.or.jp/radio-api/app/v1/web/ondemand/corners/new_arrivals"
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
	jsonBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonBytes, &availableProgramJSON); err != nil {
		return err
	}

	for _, p := range availableProgramJSON.Corners {
		fmt.Println("  - Name:", p.Title)
		fmt.Println("    URL:", generateProgramURL(p.SeriesSiteID, p.CornerSiteID))
	}
	return nil
}

func generateProgramURL(seriesSiteID string, cornerSiteID string) string {
	return fmt.Sprintf("https://www.nhk.or.jp/radio-api/app/v1/web/ondemand/series?site_id=%s&corner_site_id=%s", seriesSiteID, cornerSiteID)
}
