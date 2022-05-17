package goradiru

import (
	"testing"
	"time"
)

func Test_fmtTitle(t *testing.T) {
	type args struct {
		episode *Episode
	}
	jst, _ := time.LoadLocation("Asia/Tokyo")
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "単純連結",
			args: args{episode: &Episode{
				Title: "10時台",
				Program: &Program{
					Title: "子ども科学電話相談",
				},
				Series: &Series{
					Title: "2019年8月3日(土)「昆虫」「鳥」「水中の生物」",
				},
			}},
			want: "子ども科学電話相談 2019年8月3日(土)「昆虫」「鳥」「水中の生物」 10時台",
		},
		{
			name: "シリーズ名が空 / 番組名がエピソード名と一致",
			args: args{episode: &Episode{
				Title: "ウィークエンドサンシャイン",
				Start: time.Date(2022, time.March, 4, 5, 6, 0, 0, jst),
				Program: &Program{
					Title: "ウィークエンドサンシャイン",
				},
				Series: &Series{
					Title: "",
				},
			}},
			want: "ウィークエンドサンシャイン 20220304-0506",
		},
		{
			name: "シリーズ名が空 / 番組名がエピソード名に含まれる",
			args: args{episode: &Episode{
				Title: "カルチャーラジオ　科学と人間「過去と未来を知る進化生物学」（４）",
				Program: &Program{
					Title: "カルチャーラジオ　科学と人間",
				},
				Series: &Series{
					Title: "",
				},
			}},
			want: "カルチャーラジオ 科学と人間「過去と未来を知る進化生物学」（４）",
		},
		{
			name: "シリーズ名が空 / 番組名がエピソード名に含まれない",
			args: args{episode: &Episode{
				Title: "エピソード名",
				Program: &Program{
					Title: "番組名",
				},
				Series: &Series{
					Title: "",
				},
			}},
			want: "番組名 エピソード名",
		},
		{
			name: "不要な文字を除去",
			args: args{episode: &Episode{
				Title: "エピソード名　",
				Program: &Program{
					Title: "“番組名“",
				},
				Series: &Series{
					Title: "【シリーズ名】",
				},
			}},
			want: "番組名 シリーズ名 エピソード名",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fmtTitle(tt.args.episode); got != tt.want {
				t.Errorf("fmtTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
