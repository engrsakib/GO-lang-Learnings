package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
}

// 2. Function to Filter and Map
func getMaleNames(people []Person) []string {
	var names []string // Create an empty slice to hold the result

	// Loop through the people list
	for _, p := range people {
		// FILTER: Keep only if Gender is NOT "Female"
		if p.Gender != "Female" {
			// MAP: Extract only the 'Name' and append to our new list
			names = append(names, p.Name)
		}
	}

	return names
}

func add(x int, y int) int {
    return x + y
}

func main() {
	// a := 10
	// b := 20
	// sum := a + b
	// fmt.Println("Sum:", sum)
	// name := "Alice"
	// greeting := fmt.Sprintf("Hello, %s!, How are you?", name)
	// fmt.Println(greeting)

	// var roll int = 5
	// fmt.Println("Roll Number:", roll)

	// const pi = 3.14
	// fmt.Println("Value of Pi:", pi)

	// people := []Person{
	// 	{Name: "Sakib", Age: 25, Gender: "Male"},
	// 	{Name: "Bushra", Age: 24, Gender: "Female"},
	// 	{Name: "John", Age: 30, Gender: "Male"},
	// 	{Name: "Lisa", Age: 28, Gender: "Female"},
	// 	{Name: "Rahim", Age: 35, Gender: "Male"},
	// }

	// // 4. Call the function
	// result := getMaleNames(people)

	// // 5. Print the final result
	// fmt.Println(result)

	// age := 18
	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// }else if age < 0 {
	// 	fmt.Println("Invalid age.")
	// } else {
	// 	fmt.Println("You are a minor.")
	// }

	// day := "Friday"

	// switch day {
	// 	case "Saturday", "Friday":
    // 		fmt.Println("It's a Weekend! Enjoy.")
	// 	case "Sunday":
    // 		fmt.Println("Start of the week.")
	// 	default:
    // 		fmt.Println("Just a regular work day.")
	// }

	// // function two numbers ans return
	// sum := add(15, 25)
	// fmt.Println("Sum from function:", sum)

	result := sayHello(10, 20)
	fmt.Println("Result from sayHello function:", result)

	sum := func(x int, y int) int {
		return x + y
	}
	fmt.Println("Sum from add function:", sum(5, 10))

	greet := func(name string) {
        fmt.Printf("Hello %s\n", name)
    }

    greet("Sakib")

	arraySums := ArraySum([]int{100, 2, 3, 4, 5})
	fmt.Println("Sum of array elements:", arraySums)

	multi, div := getMultiAndDiv(20, 5)
	fmt.Println("Multiplication:", multi)
	fmt.Println("Division:", div)

	syasWelcome("Alice")
}