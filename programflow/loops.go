package main

import "fmt"

func loop_play() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		} else if i > 3 {
			fmt.Println(i)
		} else {
			fmt.Println(i, "im hereee")
		}
	}
}
