package main

import (
	"fmt"
	"github.com/user/addyAddAdd"
)

func main() {
	const World = "世界"
	var total int
	
	fmt.Printf("Hello world!\n")
	fmt.Println("Hello", World)
	total = addyAddAdd.AddyAddAdd(1, 2)
	fmt.Println("Sum1 = ", total)
	fmt.Println("Sum2 = ", addyAddAdd.AddyAddAdd(total, 4))
}

// New comment
