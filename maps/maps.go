package maps

import (
	"fmt"
)

func MonstrarMapas() {
	paises := make(map[string]string)
	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	//Asignmar varios valores al mapa
	championship := map[string]int{
		"Real Madrid":   14,
		"Milan":         7,
		"Bayern Múnich": 6,
		"Liverpool":     6,
		"Barcelona":     5,
	}
	fmt.Println(championship)

	//Imprimir los valores del mapa con formato usando for
	for equipo, copas := range championship {
		fmt.Printf("El %s, tiene %d champions ganadas \n", equipo, copas)
	}

	//Borrar una de las llaves del mapa
	delete(championship, "Barcelona")
	fmt.Println(championship)

	copas, existe := championship["Liverpool"]
	fmt.Printf("Las champions ganadas son %d, y el equipo existe = %t \n", copas, existe)

}
