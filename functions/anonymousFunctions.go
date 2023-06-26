package functions

import "fmt"

func Calculates() {
	var num3 int = 32

	calc := func(num1 int, num2 int) int {
		return num1 + num2
	}

	fmt.Println(calc(10, 25))

	calc = func(num1 int, num2 int) int {
		return num1 * num2 * num3
	}

	fmt.Println(calc(10, 25))
}
