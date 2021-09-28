package main

import "fmt"

func main() {
	exersice1()
	exersice3()
	interfaces()
}

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)
}

func exersice1() {
	var salary money = 3455.344
	salary.print()
	s := salary.printStr()
	fmt.Println(s)
}

type book struct {
	title string
	price float64
}

func (s book) vat() float64 {
	return s.price * 0.09
}
func (s book) discount(num int) float64 {
	return s.price * float64(num) / 100
}

func (b *book) changePrice(price float64) {
	(*b).price = price
}
func exersice3() {
	new_book := book{title: "weaf", price: 100.}
	discount := new_book.discount(5)
	fmt.Println(discount)
	vat := new_book.vat()
	fmt.Println(vat)
	fmt.Println(new_book.price)
	new_book.changePrice(39.4)
	fmt.Println(new_book.price)
}
