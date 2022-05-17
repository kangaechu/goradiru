package main

import (
	"fmt"
	"os"

	"github.com/kangaechu/goradiru"
)

var version string
var revision string

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf(os.Args[0]+": %s-%s\n", version, revision)
		os.Exit(0)
	}

	// 設定ファイルの読み込み
	goradiru.LoadConfig()
	execute()
}
