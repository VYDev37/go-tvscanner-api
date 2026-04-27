package main

import (
	"fmt"
	scan "tv-screener-test/scanner"
)

func main() {
	data, err := scan.FetchIDXData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", data)
}
