package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	trigger()
	exercise2()
	exercise3()
	error_detect()
}

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello %v\n", n)
	wg.Done()
}

func trigger() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello("sakthi", &wg)
	wg.Wait()
}

func sum_of_two(a float64, b float64, wg *sync.WaitGroup) float64 {
	s := a + b
	fmt.Printf("%.2f\n", s)
	wg.Done()
	return s
}

func exercise2() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go sum_of_two(float64(2*i), float64(3*i), &wg)

	}
	wg.Wait()

}

func exercise3() {
	var wg sync.WaitGroup
	wg.Add(50)
	x := 100
	for i := 0; i < 50; i++ {
		go func(n int, wg *sync.WaitGroup) {
			x := math.Sqrt(float64(n))
			fmt.Printf("%.3f\n", x)
			wg.Done()
		}(x+i, &wg)
	}

	wg.Wait()
}

func deposit(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b += n
	m.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b -= n
	m.Unlock()
	wg.Done()
}

func error_detect() {
	var wg sync.WaitGroup

	var m sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &m)
		go withdraw(&balance, i, &wg, &m)
	}
	wg.Wait()
	fmt.Println("Final balance value:", balance)
}
