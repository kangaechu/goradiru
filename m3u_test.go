package goradiru

import (
	"testing"
)

func Test_fmtTitle(t *testing.T) {
	type args struct {
		episode *Episode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "単純連結",
			args: args{episode: &Episode{
				Title: "エピソード名",
				Program: &Program{
					Title: "番組名",
				},
				Series: &Series{
					Title: "シリーズ名",
				},
			}},
			want: "番組名_シリーズ名_エピソード名",
		},
		{
			name: "番組名とシリーズ名が同じ",
			args: args{episode: &Episode{
				Title: "エピソード名",
				Program: &Program{
					Title: "番組名",
				},
				Series: &Series{
					Title: "番組名",
				},
			}},
			want: "番組名_エピソード名",
		},
		{
			name: "不要な文字を除去",
			args: args{episode: &Episode{
				Title: "エピソード名 ",
				Program: &Program{
					Title: "“番組名“",
				},
				Series: &Series{
					Title: "【シリーズ名】",
				},
			}},
			want: "番組名_ シリーズ名 _エピソード名",
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
