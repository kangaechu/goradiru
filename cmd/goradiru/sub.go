package main

import (
	"github.com/kangaechu/goradiru"
	"github.com/spf13/cobra"
)

// サブコマンドの定義
var dlCmd = &cobra.Command{
	Use: "dl",
	Run: func(c *cobra.Command, args []string) {
		goradiru.Download()
	},
}

// サブコマンドの定義
var pgCmd = &cobra.Command{
	Use: "pg",
	Run: func(c *cobra.Command, args []string) {
		goradiru.ListPrograms()
	},
}

func init() {
	// サブコマンドをルートコマンドに登録
	rootCmd.AddCommand(dlCmd)
	rootCmd.AddCommand(pgCmd)
}
