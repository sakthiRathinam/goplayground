package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	var name string = "sakthi"
	country := "india"
	fmt.Println("Your name : ", name)
	fmt.Println("Your country : ", country)
	fmt.Println("He Says \"Hello\"")
	fmt.Println("C:\\Users\\a.txt")
}

func exercise2() {
	non_char := 'ă'
	fmt.Printf("%T\n", non_char)
	ma, m := "ma", "m"
	new_string := ma + m + string(non_char)
	fmt.Println(new_string)
}

func exercise3() {
	s1 := "țară means country in Romanian"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v", s1[i])
	}
	for indx, elem := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", indx, elem)
	}
}
