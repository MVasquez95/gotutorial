package ejercicios

import (
	"strconv"
)

func EvaluaNumero(numero string) (int, string) {
	num, err := strconv.Atoi(numero)
	if err != nil {
		return 0, "Hubo un error: " + err.Error()
	} else {
		if num < 100 {
			return num, "Es menor a 100"
		} else {
			return num, "Es menor a 100"
		}
	}
}
