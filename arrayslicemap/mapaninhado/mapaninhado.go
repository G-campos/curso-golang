package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 113252.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	// delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)

		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
