package ejer_interfaces

import (
	"fmt"

	"github.com/SeiyaJapon/golang/go_twittor/interfaces"
)

func HumansBreathing(human interfaces.Human) {
	human.Breath()

	fmt.Printf("Soy un/a %s y estoy respirando \n", human.Sex())
}
