package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k float64
	fmt.Println("The program determines if a number n is a power of k. Enter n and k: ")
	fmt.Scan(&n, &k)
	if math.Sqrt(n) == k{
		fmt.Println(n,"is a power of",k)
	} else {
		fmt.Println(n,"isn't a power of",k)
	}
}	