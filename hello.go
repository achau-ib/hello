package main

import (
	"fmt"
	"addy-add-add"
)

func main() {
	const World = "世界"
	int sum;
	fmt.Printf("Hello world!\n")
	fmt.Println("Hello", World)
	sum = AddyAddAdd(1, 2)
	fmt.Println("Sum1 = ", sum)
	fmt.Println("Sum2 = ", AddyAddAdd(sum, 4))
}
