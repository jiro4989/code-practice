package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// ./main 2 3
// -> 5
func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("need 1 arguments.")
		os.Exit(1)
	}

	data := []byte(args[1])

	now := time.Now()
	s := now.Format("20060102_030405") // この書式嫌い...
	if err := ioutil.WriteFile(s+".txt", data, os.ModePerm); err != nil {
		panic(err)
	}
}
