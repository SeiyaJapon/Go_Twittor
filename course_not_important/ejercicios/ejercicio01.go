package ejercicios

import (
	"strconv"
)

func ConvertToString(myValue string) (int, string) {
	response, err := strconv.Atoi(myValue)

	if err != nil {
		return 0, "Hubo un error: " + err.Error()
	}

	if response > 100 {
		return response, "Es mayor a 100"
	} else if response < 100 {
		return response, "Es menor a 100"
	} else {
		return response, "Es igual a 100"
	}
}
