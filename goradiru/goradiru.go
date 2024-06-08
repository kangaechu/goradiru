package goradiru

// Download 設定ファイル内のProgramsにあるEpisodeをダウンロード
func Download() {
	config := GetConfig()

	dps := LoadDownloadedPrograms(config.DownloadedHistoryConfFile)
	for _, program := range config.Programs {
		program, err := CreateProgram(program.URL)
		if err != nil {
			panic(err)
		}
		err = program.Download(dps)
		if err != nil {
			panic(err)
		}
	}
	err := dps.Save()
	if err != nil {
		panic(err)
	}
}

func ListPrograms() {
	err := getAvailablePrograms()
	if err != nil {
		panic(err)
	}

}
