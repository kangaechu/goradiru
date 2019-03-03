package main

import (
	"fmt"
	"github.com/kangaechu/goradiru/pkg/goradiru"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("conf")     // 設定ファイルのファイル名
	viper.AddConfigPath("./config") // 設定ファイルのディレクトリ名
	err := viper.ReadInConfig()
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
