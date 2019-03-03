package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// ルートコマンドの定義
var rootCmd = &cobra.Command{
	Use: "app",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(`
goradiru.go ― らじる らじるを取得

 Usage: goradiru <command> [arguments...] [options...] 

 コマンドの簡単な説明:
   dl     指定したエピソードをダウンロードします
   pg     最新のプログラムを取得します

  各コマンドの詳細は goradiru <command> -h/--help を参照してください。`)
	},
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
