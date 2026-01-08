package main

import (
	"fmt"
	"myapp/mathlib"
)

func main() {
	sum := mathlib.Add(3, 5)
	fmt.Println("Sum:", sum)
}