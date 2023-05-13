package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var legend string
var err error

func InputNumbers() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Dame el numero 1: ")

	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato es incorrecto: " + err.Error())
		}
	}

	fmt.Println("Dame el numero 2: ")

	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato es incorrecto: " + err.Error())
		}
	}

	fmt.Println("Dame la leyenda: ")

	if scanner.Scan() {
		legend = scanner.Text()
	}

	fmt.Println(legend, num1*num2)
}
