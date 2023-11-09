package arrays

import "fmt"

var tabla []int = []int{2, 5, 4}
var array [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlices() {
	fmt.Println(tabla)

	slice1 := array[3:]  //Slice creado desde un vector, desde la posición 3
	slice2 := array[:5]  //Slice creado desde un vector desde 0 a 5
	slice3 := array[6:7] //Slice creado desde un vector desde 6 a 7

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

//Muestra la capacidad de nuestro array
func Capacidad() {
	elements := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d \n", len(elements), cap(elements))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
}
