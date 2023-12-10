package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MyNameSlow(nombre string, canl chan bool) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
	canl <- true
}
