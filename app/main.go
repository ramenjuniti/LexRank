package main

import (
	"fmt"
	"io/ioutil"
	"lexrank"
	"os"
)

func main() {
	file, err := ioutil.ReadFile(`./sample.txt`)
	if err != nil {
		fmt.Println("ファイルが読み込めません")
		os.Exit(1)
	}
	text := string(file)
	summary := lexrank.New()
	summary.Summarize(text, "。")
	fmt.Println(summary.LexRankScores[0].Index)
}
