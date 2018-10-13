package main

import (
	"bufio"
	"fmt"
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
	r, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	// 遅延クローズを忘れずに
	defer r.Close()

	// 書き込み対象
	w, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer w.Close()

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

		// ファイル書き込み
		fmt.Fprintln(w, line)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
}
