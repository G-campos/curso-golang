package main

import "fmt"

func main() {
	x, y := 1, 2

	// apensa postfix
	x++ //x += 1 ou x = x + 1
	fmt.Println(x)

	y-- //y -= 1 ou y = y - 1
	fmt.Println(y)

	/* NOTA
	*  Não pode ser feita o incremento/decremento
	*  em conjunto com outra operação lógica,
	*  como uma comparação de igualdade.
	*  Exemplo: fmt.Println(x++ == y--)
	 */
}
