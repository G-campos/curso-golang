package main

import "fmt"

func notaParaConceito(n float64) string {
	/*
	* DESAFIO:
	* Refatorar a estrutura de ifelseif
	* para estrutura de switch case
	* aonde o switch não tenha condicional
	*
	*  if n >= 9 && n <= 10 {
	* 	return "A"
	*  } else if n >= 8 && n < 9 {
	* 	return "B"
	*  } else if n >= 5 && n < 8 {
	* 	return "C"
	*  } else if n >= 3 && n < 5 {
	* 	return "D"
	*  } else {
	* 	return "B"
	*  }
	 */

	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
