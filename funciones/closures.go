package funciones

import "fmt"

//Normalmente utilizado para el ofuscamiento de código
func tabla(val int) func() int {
	numero := val
	sec := 0
	return func() int {
		sec++
		return numero * sec
	}
}

func LlamarClosure() {
	tabladel := 2
	mitabla := tabla(tabladel)
	for i := 1; i < 13; i++ {
		fmt.Println(mitabla())
	}
}
