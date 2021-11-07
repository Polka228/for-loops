package main

import (
	"fmt"
)

func main() {
	var as int
	fmt.Println("The program prints the first digit of the given integer.Enter the number: ")
	fmt.Scan(&as)
	for ; as < 9; {
		as = as / 10
	}
	fmt.Println("The first digit is",as)
}