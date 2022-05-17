package goradiru

import (
	"testing"
)

func TestCreateProgram(t *testing.T) {
	jsonStr := `
{
    "main": {
        "site_id": "1929",
        "program_name": "カルチャーラジオ 文学の世界",
        "mode": 0,
        "media_type": "radio",
        "media_code": "06",
        "media_name": "NHKラジオ第2",
        "site_detail": "この番組は、古今東西の名作や受賞やベストセラーで話題になっている作品をいち早く取り上げ鑑賞すると共に、著者の生き様や作品が出来るまでの知られざるエピソードを講師の解説により丁寧に辿っていきます。",
        "navi": "documentary",
        "navi_name": "ドキュメンタリー\/教養",
        "cast": null,
        "thumbnail_p": "https:\/\/www.nhk.or.jp\/radioondemand\/json\/1929\/img\/program_g_313.jpg",
        "thumbnail_c": "https:\/\/www.nhk.or.jp\/radioondemand\/json\/1929\/img\/corner_313_20.jpg",
        "site_logo": null,
        "week": null,
        "schedule": "毎週木曜 午後8時30分 | 再放送 毎週木曜 午前10時",
        "official_url": "http:\/\/www4.nhk.or.jp\/P1929\/",
        "share_url": "http:\/\/nhk.jp\/radio\/?p=1929_01",
        "corner_id": "01",
        "corner_name": null,
        "corner_detail": null,
        "noindex_flag": false,
        "program_index": false,
        "detail_list": [{
            "headline_id": "11",
            "headline": "英訳で知る“百人一首”の世界",
            "headline_sub": "翻訳家、詩人…ピーター・マクミラン",
            "headline_image": null,
            "file_list": [{
                "seq": 1,
                "file_id": "19618",
                "file_title": "第1回【百人一首の世界。外国人の視点から。】",
                "file_title_sub": null,
                "file_name": "https:\/\/nhks-vh.akamaihd.net\/i\/radioondemand\/r\/1929\/s\/stream_1929_3718bbf8dad3ce2a49fa252943a4962d.mp4\/master.m3u8",
                "open_time": "2018-10-05T15:00:00+09:00",
                "close_time": "2018-11-30T15:00:00+09:00",
                "aa_contents_id": "[radio]vod;カルチャーラジオ 文学の世界;r2,130;2018100473478;2018-10-04T20:30:00+09:00_2018-10-04T21:00:00+09:00",
                "aa_measurement_id": "vod",
                "aa_vinfo1": "カルチャーラジオ 文学の世界",
                "aa_vinfo2": "r2,130",
                "aa_vinfo3": "2018100473478",
                "aa_vinfo4": "2018-10-04T20:30:00+09:00_2018-10-04T21:00:00+09:00",
                "onair_date": "2018年10月4日(木)放送",
                "share_url": "http:\/\/nhk.jp\/radio\/?p=1929_01_19618"
            }, {
                "seq": 2,
                "file_id": "19857",
                "file_title": "第2回【百人一首翻訳の難しさと挑戦】",
                "file_title_sub": null,
                "file_name": "https:\/\/nhks-vh.akamaihd.net\/i\/radioondemand\/r\/1929\/s\/stream_1929_98a332676504a88332d94c67b3e3e5ff.mp4\/master.m3u8",
                "open_time": "2018-10-12T15:00:00+09:00",
                "close_time": "2018-12-07T15:00:00+09:00",
                "aa_contents_id": "[radio]vod;カルチャーラジオ 文学の世界;r2,130;2018101175166;2018-10-11T20:30:00+09:00_2018-10-11T21:00:00+09:00",
                "aa_measurement_id": "vod",
                "aa_vinfo1": "カルチャーラジオ 文学の世界",
                "aa_vinfo2": "r2,130",
                "aa_vinfo3": "2018101175166",
                "aa_vinfo4": "2018-10-11T20:30:00+09:00_2018-10-11T21:00:00+09:00",
                "onair_date": "2018年10月11日(木)放送",
                "share_url": "http:\/\/nhk.jp\/radio\/?p=1929_01_19857"
            }]
        }, {
            "headline_id": "12",
            "headline": "心理学者河合俊雄が読み解く村上春樹の“物語”",
            "headline_sub": "京都大学こころの未来研究センター教授…河合俊雄",
            "headline_image": null,
            "file_list": [{
                "seq": 1,
                "file_id": "22814",
                "file_title": "第1回【物語・夢のリアリティ：『夢を見るために毎朝僕は目覚めるのです』】",
                "file_title_sub": null,
                "file_name": "https:\/\/nhks-vh.akamaihd.net\/i\/radioondemand\/r\/1929\/s\/stream_1929_1cad02ae589d54f565ae7e66699b572d.mp4\/master.m3u8",
                "open_time": "2019-01-11T15:00:00+09:00",
                "close_time": "2019-03-08T15:00:00+09:00",
                "aa_contents_id": "[radio]vod;カルチャーラジオ 文学の世界;r2,130;2019011066978;2019-01-10T20:30:00+09:00_2019-01-10T21:00:00+09:00",
                "aa_measurement_id": "vod",
                "aa_vinfo1": "カルチャーラジオ 文学の世界",
                "aa_vinfo2": "r2,130",
                "aa_vinfo3": "2019011066978",
                "aa_vinfo4": "2019-01-10T20:30:00+09:00_2019-01-10T21:00:00+09:00",
                "onair_date": "2019年1月10日(木)放送",
                "share_url": "http:\/\/nhk.jp\/radio\/?p=1929_01_22814"
            }, {
                "seq": 2,
                "file_id": "23055",
                "file_title": "第2回【現代のリアリティ・バラバラとデタッチメント：『風の歌を聴け』】",
                "file_title_sub": null,
                "file_name": "https:\/\/nhks-vh.akamaihd.net\/i\/radioondemand\/r\/1929\/s\/stream_1929_c7b49cef9fa0c8cdd8fe46922fb1af8f.mp4\/master.m3u8",
                "open_time": "2019-01-18T15:00:00+09:00",
                "close_time": "2019-03-15T15:00:00+09:00",
                "aa_contents_id": "[radio]vod;カルチャーラジオ 文学の世界;r2,130;2019011768631;2019-01-17T20:30:00+09:00_2019-01-17T21:00:00+09:00",
                "aa_measurement_id": "vod",
                "aa_vinfo1": "カルチャーラジオ 文学の世界",
                "aa_vinfo2": "r2,130",
                "aa_vinfo3": "2019011768631",
                "aa_vinfo4": "2019-01-17T20:30:00+09:00_2019-01-17T21:00:00+09:00",
                "onair_date": "2019年1月17日(木)放送",
                "share_url": "http:\/\/nhk.jp\/radio\/?p=1929_01_23055"
            }]
        }]
    }
}	`
	program, err := createProgramFromJSONBytes([]byte(jsonStr))
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	{
		expected := "カルチャーラジオ 文学の世界"
		actual := program.Title

		if expected != actual {
			t.Errorf("title of program error Expected: %v Actual: %v", expected, actual)
		}
	}
	{
		expected := 2
		actual := len(program.Series)

		if expected != actual {
			t.Errorf("series index error Expected: %v Actual: %v", expected, actual)
		}
	}
	{
		expected := "英訳で知る“百人一首”の世界"
		actual := program.Series[0].Title

		if expected != actual {
			t.Errorf("title of series error Expected: %v Actual: %v", expected, actual)
		}
	}
	{
		expected := 2
		actual := len(program.Series[0].Episodes)

		if expected != actual {
			t.Errorf("episode index error Expected: %v Actual: %v", expected, actual)
		}
	}
	{
		expected := "第1回【百人一首の世界。外国人の視点から。】"
		actual := program.Series[0].Episodes[0].Title

		if expected != actual {
			t.Errorf("title of episode error Expected: %v Actual: %v", expected, actual)
		}
	}
}
