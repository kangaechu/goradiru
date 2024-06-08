package goradiru

import (
	"testing"
)

func TestCreateProgram(t *testing.T) {
	jsonStr := `
{
	"id": 458,
	"title": "ポルトガル語講座",
	"radio_broadcast": "R2",
	"corner_name": "",
	"onair_date": "2024年6月8日(土)放送",
	"thumbnail_url": "https://www.nhk.or.jp/prog/img/2769/g2769.jpg",
	"series_site_id": "2769",
	"corner_site_id": "01"
}
`
	program, err := createProgramFromJSONBytes([]byte(jsonStr))
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	{
		expected := "ポルトガル語講座"
		actual := program.Title

		if expected != actual {
			t.Errorf("title of program error Expected: %v Actual: %v", expected, actual)
		}
	}
}
