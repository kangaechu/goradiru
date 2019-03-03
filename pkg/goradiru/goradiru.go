package goradiru

// 設定ファイル内のProgramsにあるEpisodeをダウンロード
func Download() {
	config := GetConfig()

	GetDownloadedPrograms()
	for _, program := range config.Programs {
		program, err := CreateProgram(program.Url)
		if err != nil {
			panic(err)
		}
		err = program.Download()
		if err != nil {
			panic(err)
		}
	}
}

func ListPrograms() {
	err := getAvailablePrograms()
	if err != nil {
		panic(err)
	}

}
