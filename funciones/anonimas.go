package funciones

import "fmt"

func Calculos() {

	num3 := 21
	num4 := 777

	cal := func(num1, num2 int) int {
		return num1 + num2 + num3 + num4
	}
	fmt.Println(cal(19, 20))

	cal = func(num1, num2 int) int {
		return num1 * num2 * num3
	}
	fmt.Println(cal(19, 20))
}
