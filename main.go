package main

import (
	"fmt"

	"github.com/gotutorial/ejercicios"
)

func main() {
	//estado, texto := variables.ConvierteaTexto(1920)
	numero, result := ejercicios.EvaluaNumero("100")
	fmt.Println(numero)
	fmt.Println(result)
	numero, result = ejercicios.EvaluaNumero("18")
	fmt.Println(numero)
	fmt.Println(result)
}
