package main

import (
	"github.com/kangaechu/goradiru"
)

func main() {
	// 設定ファイルの読み込み
	goradiru.LoadConfig()
	execute()
}
