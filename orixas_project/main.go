package main

import (
	"fmt"
)

func main() {
	var dia, mes, ano int
	fmt.Print("Coloque sua Data de nascimento: ")
	fmt.Scan(&dia, &mes, &ano) // lê cada uma das variaveis juntas
	//chama a função somaDigitos para cada uma das variaveis
	diaSoma := somaDigitos(dia)
	mesSoma := somaDigitos(mes)
	anoSoma := somaDigitos(ano)

	total := diaSoma + mesSoma + anoSoma // A soma das somas dos dígitos de dia, mes e ano é calculada e armazenada na variável total.
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
