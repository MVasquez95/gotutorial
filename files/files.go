package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gotutorial/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

// Graba la tabla
func GrabaTabla() {
	text := ejercicios.TablaMul()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo: " + err.Error())
		return
	}
	fmt.Fprintln(file, text)
	file.Close()
}

// Concatena la nueva tabla con lo existente
func JuntaTabla() {
	text := ejercicios.TablaMul()
	if !append(text) {
		fmt.Println("Error al concatenar contenido")
	}
}

func append(text string) bool {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append: " + err.Error())
		return false
	}

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error durante al escribir en el archivo: " + err.Error())
		return false
	}
	file.Close()
	return true
}

// Lectura de archivo
func LeeArchivo() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error durante al leer el archivo: " + err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	file.Close()
}
