package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	infile := "access_log"
	r, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()

		// TODO

		ip := ""
		fmt.Println(ip)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
}
