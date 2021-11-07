package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("The program determines if an integer is prime.Enter the number: ")
	fmt.Scan(&n)
	a := 2
	for ; a < n ; a++ {
		if n%a == 0 {
			fmt.Println(n, "isn't a prime number")
		}
	}
	fmt.Println(n,"is a prime number")
}	