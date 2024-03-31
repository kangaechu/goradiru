package cmd

import (
	"github.com/kangaechu/goradiru/goradiru"

	"github.com/spf13/cobra"
)

// dlCmd represents the dl command
var dlCmd = &cobra.Command{
	Use:   "dl",
	Short: "指定したエピソードをダウンロード",
	Long:  `指定したエピソードをダウンロード`,
	Run: func(_ *cobra.Command, _ []string) {
		goradiru.LoadConfig()
		goradiru.Download()
	},
}

func init() {
	rootCmd.AddCommand(dlCmd)
}
