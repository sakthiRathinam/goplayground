package main

import (
	"fmt"
	"os"
)

func main() {
	language := os.Args[1]
	switch language {
	case "python":
		fmt.Println("python is my favourite")
	case "go", "golang":
		fmt.Println("i also like go because is similar to python")
	default:
		fmt.Println("any other programming languages is also good start")
	}

}
