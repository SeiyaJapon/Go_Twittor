package arrayslices

import "fmt"

var tableS []int = []int{2, 5, 4}
var array1 [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func ShowSlice() {
	fmt.Println(tableS)

	piece := array1[3:]   // slice creado desde un vector, desde la posicón 3
	piece2 := array1[:5]  // slice creado desde un vector, desde la posición 0 hasta la 5
	piece3 := array1[6:7] // slice creado desde un vector, desde la posición 6 hasta la 7

	fmt.Println(piece)
	fmt.Println(piece2)
	fmt.Println(piece3)
}

func Capacity() {
	elements := make([]int, 5, 20)

	fmt.Printf("Largo %d, Capacidad %d", len(elements), cap(elements))

	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
