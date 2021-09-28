package main

import "fmt"

func main() {
	args_play()
	loop_play()
	price, in_stock := 100, true
	if price > 80 {
		fmt.Printf("stock price is %d\n", price)
	}
	if price <= 100 && in_stock {
		fmt.Println("awfawefawef")
	}

}
