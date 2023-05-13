package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var err error

func MultTable() {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Dame el numero 1: ")

		if scanner.Scan() {
			num1, err = strconv.Atoi(scanner.Text())
		}

		if err == nil {
			break
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", num1, i, i*num1)
	}
}
