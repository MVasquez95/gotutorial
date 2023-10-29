package main

import (
	"fmt"

	"github.com/gotutorial/variables"
)

func main() {
	estado, texto := variables.ConvierteaTexto(1920)
	fmt.Println(estado)
	fmt.Println(texto)
}
