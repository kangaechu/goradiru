package goradiru

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

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

func LoadConfig() *config {
	// スクリプトのディレクトリを取得
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigName("conf")                       // 設定ファイルのファイル名
	viper.AddConfigPath(filepath.Join(dir, "config")) // 設定ファイルのディレクトリ名
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error on reading config: %s \n", err))
	}
	config := GetConfig()

	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("error on parsing config %s \n", err))
	}
	config.ProgDir = filepath.Join(dir, config.ProgDir)
	config.DownloadedHistoryConfFile = filepath.Join(dir, config.DownloadedHistoryConfFile)
	return config

}
