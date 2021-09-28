package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	continue_loop()
	check_both()
	print_all_years()
	swith_practice()
	first_three := 0
	for i := 0; i < 51; i++ {
		if first_three == 3 {
			break
		}
		if i%7 == 0 {
			fmt.Println(i)
			first_three++

		}
	}
}

func continue_loop() {
	for i := 0; i < 51; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

func check_both() {
	for i := 0; i < 500; i++ {
		if i%5 == 0 || i%7 == 0 {
			fmt.Printf("%v ", i)
		}
		fmt.Printf("")
	}
}

func print_all_years() {
	birth_year, current_year := 1997, 2021
	for i := birth_year; i <= current_year; {
		fmt.Println(i)
		i++
	}
}
func swith_practice() {
	age, err := strconv.ParseInt(os.Args[1], 10, 64)
	_ = err
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age <= 10:
		fmt.Println("underage")
	case age == 18:
		fmt.Println("you havew bevcime majaor")
	default:
		fmt.Println("you are makot")
	}

}
