package main

import (
	"fmt"
	"os"
	"strconv"
)

func args_play() {
	fmt.Println(os.Args)
	if os.Args[1] == "sakthi" {
		fmt.Println("super hereoooo")
	}
	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("%.2f\n", result)
	_, _ = err, result
	// multiple if statemnets
	if i, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
		fmt.Printf("this is the float num %.2f\n", i)
	}
}
