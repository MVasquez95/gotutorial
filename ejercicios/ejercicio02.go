package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int
var err error
var text string

func TablaMul() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese un número del 1 al 10: ")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil || 1 > num || 10 < num {
				continue
			} else {
				break
			}
		}
	}
	for i := 0; i < 13; i++ {
		text += fmt.Sprintf("%d x %d = %d \n", num, i, num*i)
	}
	return text
}
