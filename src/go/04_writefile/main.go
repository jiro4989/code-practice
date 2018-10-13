package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("need 2 arguments.")
		os.Exit(1)
	}

	infile := args[1]
	outfile := args[2]

	// ファイル読み込み
	b, err := ioutil.ReadFile(infile)
	if err != nil {
		panic(err)
	}
	// byte型なんでstringに変換
	s := string(b)

	// 改行文字で区切って配列に
	lines := strings.Split(s, "\n")
	newLines := make([]string, len(lines))
	i := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		i++
		line = fmt.Sprintf("%03d:%s", i, line)
		newLines[i-1] = line
	}

	data := []byte(strings.Join(newLines, "\n"))
	if err := ioutil.WriteFile(outfile, data, os.ModePerm); err != nil {
		panic(err)
	}
}
