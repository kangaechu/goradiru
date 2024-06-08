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
			name: "不要な文字を除去",
			args: args{episode: &Episode{
				ProgramTitle: "【エピソード名】　",
			}},
			want: "エピソード名",
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
