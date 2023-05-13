package main

import (
	"fmt"
	"runtime"

	"github.com/SeiyaJapon/golang/go_twittor/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	os := runtime.GOOS

	if os == "Linux." || os == "darwin" {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os {
	case "Windows":
		fmt.Println("Esto es Windows")
	default:
		fmt.Printf("Esto es %s", os)
	}
}
