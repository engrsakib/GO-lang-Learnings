package printlib

import (
	"fmt"
)

func PrintMessage(message string) {
	fmt.Println(message)
}

func PrintSum(str string, result int) {
	fmt.Printf("%s: %d\n", str, result)
}