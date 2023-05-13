package main

import (
	"fmt"

	"github.com/SeiyaJapon/golang/go_twittor/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
