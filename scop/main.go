package main

import (
	"myapp/mathlib"
	"myapp/printlib"
)

func main() {
	sum := mathlib.Add(3, 5)
	printlib.PrintSum("Sum of 3 and 5", sum)

	product := mathlib.Multiply(4, 7)
	printlib.PrintSum("Product of 4 and 7", product)

	diff := mathlib.Subtract(10, 40)
	printlib.PrintSum("Difference of 10 and 40", diff)
}