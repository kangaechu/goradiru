package cmd

import (
	"github.com/kangaechu/goradiru/goradiru"

	"github.com/spf13/cobra"
)

// pgCmd represents the pg command
var pgCmd = &cobra.Command{
	Use:   "pg",
	Short: "最新のプログラムを取得",
	Long:  `最新のプログラムを取得`,
	Run: func(_ *cobra.Command, _ []string) {
		goradiru.LoadConfig()
		goradiru.ListPrograms()
	},
}

func init() {
	rootCmd.AddCommand(pgCmd)
}
