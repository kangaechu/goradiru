package goradiru

import "fmt"

// 設定ファイル内のProgramsにあるEpisodeをダウンロード
func Download() {
	config := GetConfig()

	GetDownloadedPrograms()
	for _, program := range config.Programs {
		program, err := CreateProgram(program.Url)
		if err != nil {
			panic(err)
		}
		fmt.Println(program.Title)
		err = program.Download()
		if err != nil {
			panic(err)
		}
	}
}

func ListPrograms() {
	getPrograms()

}
