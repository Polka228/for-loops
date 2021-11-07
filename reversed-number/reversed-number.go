package main

import (
	"fmt"
)

func main() {
	var as,a,b,c int
	fmt.Println("The program reverses the digits in the number. Enter the number: ")
	fmt.Scan(&as)
	for a = as ; a < 9 ; {
		a = a / 10
	}
	for b = as ; b < 99 ; {
		b = b / 100
	}
	for c = as ; c < 999 ; {
		c = c / 1000
	}
	fmt.Println("The reversed number is",c,b,a)
}