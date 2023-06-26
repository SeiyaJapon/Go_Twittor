package middleware

import "fmt"

func add(a, b int) int {
	return a + b
}

func substrac(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func MyMiddleware() {
	fmt.Println("Init")

	result := operateMidd(add)(2, 3)
	fmt.Println(result)

	result = operateMidd(substrac)(10, 6)
	fmt.Println(result)

	result = operateMidd(mult)(2, 4)
	fmt.Println(result)
}

func operateMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")

		return f(a, b)
	}
}
