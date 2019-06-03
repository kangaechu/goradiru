package main

import (
	"fmt"
	"github.com/kangaechu/goradiru"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

func main() {

	// スクリプトのディレクトリを取得
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigName("conf")                         // 設定ファイルのファイル名
	viper.AddConfigPath(filepath.Join(dir, "./config")) // 設定ファイルのディレクトリ名
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error on reading config: %s \n", err))
	}

	// 設定ファイルの読み込み
	config := goradiru.GetConfig()
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("error on parsing config %s \n", err))
	}
	execute()
}
