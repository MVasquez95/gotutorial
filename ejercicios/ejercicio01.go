package ejercicios

import (
	"strconv"
)

func EvaluaNumero(numero string) (int, string) {
	num, err := strconv.Atoi(numero)
	var text string = ""
	if err != nil {
		text = "Error"
	} else {
		if num < 100 {
			text = "Es menor a 100"
		} else {
			text = "Es mayor a 100"
		}
	}

	return num, text
}
