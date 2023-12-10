package main

import (
	"github.com/gotutorial/middleware"
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
	}

	//Ejercicio 1
	numero, result := ejercicios.EvaluaNumero("100")
	fmt.Println(numero)
	fmt.Println(result)

	//Aceptar datos por teclado
	teclado.IngresoNumero()

	//Pruebas de FOR
	iteraciones.Iterar()

	//Ejercicio 2
	fmt.Println(ejercicios.TablaMul())

	files.GrabaTabla()
	files.JuntaTabla()
	files.LeeArchivo()

	//Funciones Anonimas
	funciones.Calculos()
	funciones.LlamarClosure()

	funciones.Exponencia(2)

	arrays.MuestroArreglos()
	arrays.MuestroSlices()
	arrays.Capacidad()

	maps.MonstrarMapas()

	users.AltaUsuario()

	Pablo := new(models.Hombre)
	e.HumanoRespirando(Pablo)

	Tasha := new(models.Mujer)
	e.HumanoRespirando(Tasha)

	deferpanic.EjemploPanic()

	canl := make(chan bool)
	go gorutines.MyNameSlow("Miguel", canl)
	defer func() {
		<-canl
	}()
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)

	webserver.MiWebServer()*/

	middleware.MiMiddleware()
}
