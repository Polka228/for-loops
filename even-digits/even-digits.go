package main

import (
	"fmt"
)

func main() {
	var as,a,b,c int
	fmt.Println("The program determines if all digits in the integer are even. Enter the number: ")
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
	if a %= 2 && b %= 2 && c %= 2{
		fmt.Println("All digits in",as,"are even")
	} else {
		fmt.Println("Not all digits in",as,"are even")
	}
}