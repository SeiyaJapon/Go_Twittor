package maps

import "fmt"

func ShowMaps() {
	countries := make(map[string]string)

	fmt.Println(countries)

	countries["España"] = "Madrid"
	countries["Japón"] = "Tokyo"

	fmt.Println(countries)
	fmt.Println(countries["Japón"])

	championship := map[string]int{
		"Real Madrid":     39,
		"Barcelona":       0,
		"Manchester City": 30,
		"Bayern Munich":   29,
	}

	fmt.Println(championship)

	for team, points := range championship {
		fmt.Printf("Equipo %s, tiene %d puntos \n", team, points)
	}

	delete(championship, "Barcelona")

	fmt.Println(championship)

	pointTeams, exists := championship["Juventus"]

	fmt.Printf("La puntuación es de %d, y el equipo existe = %t \n", pointTeams, exists)

	pointTeams, exists = championship["Bayern Munich"]

	fmt.Printf("La puntuación es de %d, y el equipo existe = %t \n", pointTeams, exists)
}
