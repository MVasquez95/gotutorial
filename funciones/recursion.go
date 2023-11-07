package funciones

import "fmt"

func Exponencia(val int) {
	if val > 100000000 {
		return
	}
	fmt.Println(val)
	Exponencia(val * 2)
}
