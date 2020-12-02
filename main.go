package main

import (
	"fmt"
)

func sum(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	fmt.Println("Master Branch!")
	fmt.Println(sum(10, 20))
	fmt.Println(sub(20, 10))
}
