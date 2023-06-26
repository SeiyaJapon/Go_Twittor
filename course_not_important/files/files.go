package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/SeiyaJapon/golang/go_twittor/ejercicios"
)

var fileName string = "./files/tmp/table.txt"

func StoreTable() {
	var text string = ejercicios.MultTable()

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())

		return
	}

	fmt.Fprintln(file, text)

	file.Close()
}

func AppendTable() {
	var text string = ejercicios.MultTable()

	if !ApppendToFile(fileName, text) {
		fmt.Println("Error al concatenar contenido")
	}
}

func ApppendToFile(filename string, text string) bool {
	buffer, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())

		return false
	}

	_, err = buffer.WriteString(text)

	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())

		return false
	}

	buffer.Close()

	return true
}

func ReadFile() {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	fmt.Println(string(file))
}

func ReadFile2() {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		record := scanner.Text()

		fmt.Println("> " + record)
	}
}
