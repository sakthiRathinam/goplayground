package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	cube_value := cube(4)
	fmt.Println(cube_value, "first")
	fact, numb := factorial(2, 4, 5, 6)
	fmt.Println(fact, numb, "second")
	sum_of_three := repeated_nums("4")
	fmt.Println(sum_of_three, "third")
	naked_function := naked_return(1, 2, 4, 5)
	fmt.Println(naked_function, "fourth")
	animals := []string{"lion", "tiger", "bear"}
	result := search_item(animals, "bear")
	fmt.Println(result, "fourth")
	ss()
	first_class_define()
}

func cube(value float64) float64 {
	return math.Pow(value, 3)
}

func factorial(n ...int) (int, int) {
	fact := 1
	num := 0
	for i := 0; i < len(n); i++ {
		fact *= n[i]
		num += n[i]

	}
	return fact, num

}

func repeated_nums(n string) int {
	value, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal("this is not a string integer")
	}
	second_value, _ := strconv.Atoi(n + n)
	third_value, _ := strconv.Atoi(n + n + n)
	return value + second_value + third_value
}

func naked_return(n ...int) (s int) {
	for i := 0; i < len(n); i++ {
		s += n[i]
	}
	return
}

func search_item(animals []string, animal string) bool {
	for i := 0; i < len(animals); i++ {
		if strings.EqualFold(animal, animals[i]) {
			return true
		}
	}
	return false

}
func print(m string) {
	fmt.Println(m)
}
func ss() {
	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
}

// first class function

func f1(a int) {
	fmt.Println(a)
}

func first_class_define() {
	x := f1
	fmt.Printf("%T\n", x)
	x(4)
}
