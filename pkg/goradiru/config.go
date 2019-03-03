package goradiru

type config struct {
	ProgDir                   string `mapstructure:"ProgDir"`
	FileType                  string `mapstructure:"FileType"`
	DownloadedHistoryConfFile string `mapstructure:"DownloadedHistoryConfFile"`
	Programs                  []struct {
		Name string `mapstructure:"Name"`
		Url  string `mapstructure:"Url"`
	} `mapstructure:"Programs"`
}

var sharedInstance = &config{ /* 初期化 */ }

func GetConfig() *config {
	return sharedInstance
}
