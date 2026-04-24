package main

import "fmt"

func main() {
	q, r, err := Divide(10, 3)

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("quotient: ", q)
	fmt.Println("reminder: ", r)
}
