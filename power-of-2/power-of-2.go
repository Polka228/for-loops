package main

import (
	"fmt"
)

func main() {
	var n,a,b int
	fmt.Println("The program determines if a number is a power of 2.Enter numder: ")
	fmt.Scan(&n)
	if n <= 0 || n%2 != 0 {
		fmt.Println(n, "is not a power of 2")
	} 
	for x := n; x != 1, x /= 2,++a {
		b = x % 2
	}
	if b == 0 {
		fmt.Println(n, "= 2^", a)
	} else {
		fmt.Println(n, "is not a power of 2")
	}
}