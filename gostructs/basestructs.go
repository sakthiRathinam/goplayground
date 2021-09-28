package main

import "fmt"

func main() {
	exercise1()
}

func exercise1() {
	type grade struct {
		grade   int
		courses string
	}
	type person struct {
		name   string
		age    int
		colors []string
		gr     grade
	}

	me := person{name: "sakthi", age: 32, colors: []string{"awefawf", "aewfawfawf"}, gr: grade{grade: 100, courses: "python"}}
	you := person{name: "fuck", age: 32, colors: []string{"awefawf", "aewfawfawf"}, gr: grade{grade: 80, courses: "golang"}}
	fmt.Println(you, me)
	var colors []string = you.colors
	fmt.Printf("Type: %T, Value: %v\n", colors, colors)
	Andrei := me
	Andrei.colors = append(Andrei.colors, "awefawfaewf")
	for i := 0; i < len(me.colors); i++ {
		fmt.Println(me.colors[i])
	}

}
