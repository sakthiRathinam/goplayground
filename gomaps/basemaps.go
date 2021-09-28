package main

import "fmt"

func main() {
	exercise1()
	exercise2()
}

func exercise1() {
	var first_map map[int]bool
	fmt.Printf("%#v\n", first_map)
	second_map := map[int]string{1: "awefawf", 3: "awefawffae"}
	second_map[10] = "abba"
	fmt.Println(second_map[10])
	fmt.Println(second_map[109], "Awefawffewa")
}

func exercise2() {
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 1)
	for key, val := range m {
		fmt.Printf("%#v %#v\n", key, val)
	}
}
