package ejerinterfaces

import (
	"fmt"

	"github.com/gotutorial/interfaces"
)

func HumanoRespirando(h interfaces.Humano) {
	h.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", h.Sexo())
}
