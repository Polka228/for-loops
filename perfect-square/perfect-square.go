package main

import (
	"fmt"
	"math"
)

func main() {
	var n,a int
	fmt.Println("The program determines if an integer is a perfect square. Enter the number: ")
	fmt.Scan(&n)
	a = math.Sqrt(n)
	if a %= 0 {
		fmt.Println(a,"*",a,"=",n)
		fmt.Println(n,"is a perfect square") 
	} else{
		fmt.Println(n,"isn't a perfect square")
	}
}