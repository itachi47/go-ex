package main

import "fmt"

func main() {
	status := Paid

	fmt.Println(status)
	fmt.Println(status.String())

	var x PaymentStatus = 99
	fmt.Println(x)
}
