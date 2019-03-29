package goradiru

// 設定ファイル内のProgramsにあるEpisodeをダウンロード
func Download() {
	config := GetConfig()

	dps := LoadDownloadedPrograms(config.DownloadedHistoryConfFile)
	for _, program := range config.Programs {
		program, err := CreateProgram(program.Url)
		if err != nil {
			panic(err)
		}
		err = program.Download(dps)
		if err != nil {
			panic(err)
		}
	}
	dps.Save()
}

func ListPrograms() {
	err := getAvailablePrograms()
	if err != nil {
		panic(err)
	}

}
