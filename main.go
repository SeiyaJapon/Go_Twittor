package main

import (
	"github.com/SeiyaJapon/golang/go_twittor/ejer_interfaces"
	"github.com/SeiyaJapon/golang/go_twittor/models"
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

	Pepe := new(models.Man)

	ejer_interfaces.HumansBreathing(Pepe)

	Lola := new(models.Woman)

	ejer_interfaces.HumansBreathing(Lola)
}
