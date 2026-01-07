package main // ১. প্যাকেজ নাম অবশ্যই main হতে হবে
import "fmt"

// এটি একটি কাস্টম ফাংশন
func sayHello(a int, b int) int {
    return a + b
}

func ArraySum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}	
	return sum
}

func getMultiAndDiv(x int, y int) (int, int) {
	return x * y, x / y
}

func syasWelcome(name string) {
	fmt.Printf("Welcome, %s!\n", name)
}