package middleware

import "fmt"

func sum(a, b int) int {
	return a + b
}

func res(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	result := operacionesMidd(sum)(2, 3)
	println(result)
	result = operacionesMidd(res)(10, 6)
	println(result)
	result = operacionesMidd(mul)(4, 7)
	println(result)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")
		return f(a, b)
	}
}
