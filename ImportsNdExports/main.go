package main

import (
	"fmt"

	"./concatenate"
)

func main() {
	a, b := concatenate.Swap("hello", "world")
	fmt.Println(a, b)
	c, d := concatenate.Concatenate("good", "girl")
	fmt.Println(c, d)
	fmt.Println(concatenate.Calladd(6, 7))
}
