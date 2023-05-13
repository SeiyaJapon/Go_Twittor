package main

import (
	"fmt"

	"github.com/SeiyaJapon/golang/go_twittor/ejercicios"
)

func main() {
	number, text := ejercicios.ConvertToString("99")
	fmt.Println(number)
	fmt.Println(text)

	number2, text2 := ejercicios.ConvertToString("100")
	fmt.Println(number2)
	fmt.Println(text2)

	number3, text3 := ejercicios.ConvertToString("101")
	fmt.Println(number3)
	fmt.Println(text3)

	number4, text4 := ejercicios.ConvertToString("asdas")
	fmt.Println(number4)
	fmt.Println(text4)
}
