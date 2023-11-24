package deferpanic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Fatalf("Ocurrió un error que generó Panic \n %v", rec)
		}
	}()
	if a := 1; a == 1 {
		panic("Se encontro el valor 1")
	}
}
