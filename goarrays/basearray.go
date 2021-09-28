package main

import "fmt"

func main() {
	operations()
	exericse_1()
	exericse_2()
	numbers := [3]int{1, 2, 3}
	numbers[2] = 7
	fmt.Printf("%v\n", numbers)
	for indx, n := range numbers {
		fmt.Printf("index of %d is %d\n", indx, n)
	}
}

func operations() {
	grades := [3]int{2: 1, 1: 3, 0: 4}
	fmt.Printf("%v\n", grades)

}
