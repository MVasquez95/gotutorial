package main

import (
	"fmt"

	"github.com/gotutorial/ejercicios"
)

func main() {
	////Convertir un entero a texto, se necesita importar package variables
	/*estado, texto := variables.ConvierteaTexto(1920)
	numero, result := ejercicios.EvaluaNumero("100")
	fmt.Println(numero)
	fmt.Println(result)
	numero, result = ejercicios.EvaluaNumero("18")
	fmt.Println(numero)
	fmt.Println(result)

	//Condicional
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("No estamos usando windows, usamos: ", os)
	} else {
		fmt.Println("Estamos usando Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	} */

	numero, result := ejercicios.EvaluaNumero("100")
	fmt.Println(numero)
	fmt.Println(result)
	numero, result = ejercicios.EvaluaNumero("18")
	fmt.Println(numero)
	fmt.Println(result)
}
