package main

import (
	"bufio"
	"fmt"
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
	r, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	// 遅延クローズを忘れずに
	defer r.Close()

	i := 0
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()

		// ここは前のプログラムの使い回し
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		i++
		line = fmt.Sprintf("%03d:%s", i, line)
		fmt.Println(line)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
}
