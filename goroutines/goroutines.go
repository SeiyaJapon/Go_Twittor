package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MyNameSlow(name string, chanel1 chan bool) {
	characters := strings.Split(name, "")

	for _, character := range characters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(character)
	}

	chanel1 <- true
}
