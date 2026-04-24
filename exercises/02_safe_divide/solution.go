package main

import "errors"

func Divide(a, b int) (quotient int, remainder int, err error) {
	if b == 0 {
		err = errors.New("division by 0")
		return
	}

	quotient = a / b
	remainder = a % b
	return
}
