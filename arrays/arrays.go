package arrays

import "fmt"

var table [10]int = [10]int{10, 0, 59}
var matriz [20][30]int

func MuestroArreglos() {
	table[7] = 33
	table[2] = 54

	table2 := [10]string{"Pablo", "Juan"}
	fmt.Println(table)
	fmt.Println(table2)

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
	matriz[7][24] = 15
	fmt.Println(matriz)
}
