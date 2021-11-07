package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Println("The program determines if a number n is a power of k.Enter n and k: ")
	fmt.Scan(&n, &k)
	if n%2 == 0 {
		fmt.Println("")
	}
	math.Sqrt(n)
}
