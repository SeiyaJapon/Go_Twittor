package defer_panic

import (
	"fmt"
	"log"
)

func ShowDefer() {
	fmt.Println("Este es un primer mensaje")

	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Estes es el segundo mensaje")
}

func PanicExample() {
	defer func() {
		reco := recover()

		if reco != nil {
			log.Fatalf("ocurrió un error que generó Panic \n %v", reco)
		}
	}()

	a := 1

	if 1 == a {
		panic("Se encontró el valor 1")
	}
}
