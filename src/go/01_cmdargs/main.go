package main

import (
	"fmt"
	"os"
	"strconv"
)

// ./main 2 3
// -> 5
func main() {
	args := os.Args

	// 引数が2つ渡されなかったらエラー
	// ※Goでは第一引数に必ず実行するバイナリ自身が含まれる
	if len(args) < 3 {
		fmt.Println("need 2 arguments.")
		os.Exit(1)
	}

	fmt.Println("args[0]:" + args[0]) // ./main
	fmt.Println("args[1]:" + args[1]) // コマンド実行時に渡す引数はここから
	fmt.Println("args[2]:" + args[2])

	// 型変換
	x, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("args[1] is not integer.", args[1])
		panic(err)
	}

	y, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("args[2] is not integer.", args[2])
		panic(err)
	}

	sum := x + y
	s := fmt.Sprintf("sum:%d\n", sum)
	fmt.Println(s)
}
