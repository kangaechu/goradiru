package goradiru

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	ProgDir                   string `mapstructure:"ProgDir"`
	FileType                  string `mapstructure:"FileType"`
	DownloadedHistoryConfFile string `mapstructure:"DownloadedHistoryConfFile"`
	Programs                  []struct {
		Name string `mapstructure:"Name"`
		URL  string `mapstructure:"URL"`
	} `mapstructure:"Programs"`
}

var sharedInstance = &Config{ /* 初期化 */ }

func GetConfig() *Config {
	return sharedInstance
}

func LoadConfig() *Config {
	// スクリプトのディレクトリを取得
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigName("conf")                       // 設定ファイルのファイル名
	viper.AddConfigPath(filepath.Join(dir, "config")) // 設定ファイルのディレクトリ名
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error on reading config: %s", err))
	}
	config := GetConfig()

	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("error on parsing config %s", err))
	}
	config.ProgDir = filepath.Join(dir, config.ProgDir)
	config.DownloadedHistoryConfFile = filepath.Join(dir, config.DownloadedHistoryConfFile)
	return config

}
