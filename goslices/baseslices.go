package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var nilSlice []string
	excersie1()
	exercise2()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise8()
	fmt.Println(nilSlice == nil)
	emptySlice := make([]string, 5)
	fmt.Printf("%#v\n", emptySlice)
}

func excersie1() {
	str_slice := []string{"heklp", "new", "golang"}
	for indx, elem := range str_slice {
		fmt.Printf("%v %v\n", indx, elem)

	}

}

func exercise2() {
	mySlice := []float64{1.2, 5.6}

	mySlice = append(mySlice, float64(6))

	a := 10
	mySlice[0] = float64(a)

	mySlice = append(mySlice, 10.10)

	mySlice = append(mySlice, float64(a))

	fmt.Println(mySlice)
}

func exercise3() {
	new_slice := []float64{1.4, 1.6, 1.7}
	new_slice = append(new_slice, 10.1)
	new_slice = append(new_slice, 4.1, 5.5, 6.6)
	fmt.Printf("%#v\n", new_slice)
	n := []float64{4.6, 8.8}
	new_slice = append(new_slice, n...)
	fmt.Printf("%#v\n", new_slice)
}

func exercise4() {
	if len(os.Args) < 3 {
		log.Fatalf("this program should nt be run without the args")
	}
	args_slice := os.Args[1:]
	num, slice := 0., 1.
	for i := 0; i < len(args_slice); i++ {
		value, err := strconv.ParseFloat(args_slice[i], 64)
		if err != nil {
			continue
		}
		num += value
		slice *= value

	}
	fmt.Printf("%v imhereee\n", num)
	fmt.Printf("%v\n", slice)
}

func exercise5() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum := 0
	for indx, elem := range nums[1 : len(nums)-2] {
		_ = indx
		fmt.Printf("%v\n", elem)
		sum += elem
	}
	fmt.Printf("%v\n", sum)
}

func exercise6() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	new_friends := make([]string, len(friends))
	copy(new_friends, friends)
	new_friends[0] = "Sakthi"
	fmt.Printf("%#v %#v\n", friends, new_friends)
}
func exercise7() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	new_friends := []string{}
	new_friends = append(new_friends, friends[0:len(friends)]...)
	new_friends[0] = "sakthi"
	fmt.Printf("%#v %#v\n", friends, new_friends)

}
func exercise8() {
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	new_years := append(years[0:3], years[len(years)-3:]...)
	fmt.Printf("%#v\n", new_years)
}
