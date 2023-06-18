package main

import (
	"github.com/SeiyaJapon/golang/go_twittor/goroutines"
)

func main() {
	// fmt.Println(ejercicios.MultTable())

	// files.AppendTable()

	//files.ReadFile()

	// files.ReadFile2()

	// functions.Calculates()

	// functions.CallClosure()

	// arrayslices.Capacity()

	// maps.ShowMaps()

	// users.SingUpUser()

	// Pepe := new(models.Man)

	// ejer_interfaces.HumansBreathing(Pepe)

	// Lola := new(models.Woman)

	// ejer_interfaces.HumansBreathing(Lola)

	// defer_panic.PanicExample()

	chan1 := make(chan bool)

	go goroutines.MyNameSlow("Francisco Pérez", chan1)

	<-chan1 // awake
}
