package mathlib

import (
	"fmt"
	"math"
)

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func Power(base int, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func Factorial(n int) int {
	if n < 0 {
		return -1 // Indicate error for negative input
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}