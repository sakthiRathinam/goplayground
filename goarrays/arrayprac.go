package main

import (
	"fmt"
)

func exericse_1() {
	var cities [2]string
	fmt.Printf("%v\n", cities)
	grades := [3]float64{1: 2.3, 2: 3.4}
	fmt.Printf("%v\n", grades)
	task_done := [...]bool{5: true}
	fmt.Printf("%v\n", task_done)
	for i := 0; i < len(cities); i++ {
		fmt.Printf("%v\n", cities[i])
	}
	for indx, element := range grades {
		fmt.Printf("index of %d is %f\n", indx, element)
	}

}

func exericse_2() {
	nums := []int{30, -1, -6, 90, -6}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 && nums[i] >= 0 {
			count++
		}
	}
	fmt.Println(count, "total filters in that array")
}
