package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("need 1 arguments.")
		os.Exit(1)
	}

	// ファイル読み込み
	infile := args[1]
	b, err := ioutil.ReadFile(infile)
	if err != nil {
		panic(err)
	}
	// byte型なんでstringに変換
	s := string(b)

	// 改行文字で区切って配列に
	lines := strings.Split(s, "\n")
	i := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		i++
		line = fmt.Sprintf("%03d:%s", i, line)
		fmt.Println(line)
	}
}
