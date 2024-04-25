package main

import (
	"fmt"
)

func main() {
	var dia, mes, ano int
	fmt.Print("Coloque sua Data de nascimento: ")
	fmt.Scan(&dia, &mes, &ano)
	diaSoma := somaDigitos(dia)
	mesSoma := somaDigitos(mes)
	anoSoma := somaDigitos(ano)

	total := diaSoma + mesSoma + anoSoma
	if total > 16 {
		total = somaDigitos(total)
	}
	fmt.Println("valor total é: ", total)

}

func somaDigitos(numero int) int {
	soma := 0
	for numero > 0 {
		soma += numero % 10
		numero /= 10
	}
	return soma
}
