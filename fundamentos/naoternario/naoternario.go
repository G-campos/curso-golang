package main

import "fmt"

func obterResultado(nota float64) string {
	//O go não tem operadores ternários
	// return nota	>= 6 ? "Aprovado" : "Reprovado"

	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
