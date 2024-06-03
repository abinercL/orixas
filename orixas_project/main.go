package main

import (
	"fmt"
)

var dia, mes, ano int

// 1°parte
func main() {

	fmt.Print("Coloque sua Data de nascimento: ")
	fmt.Scan(&dia, &mes, &ano) // lê cada uma das variaveis juntas
	//chama a função somaDigitos para cada uma das variaveis
	diaSoma := SomaDigitos(dia)
	mesSoma := SomaDigitos(mes)
	anoSoma := SomaDigitos(ano)

	total := diaSoma + mesSoma + anoSoma // A soma das somas dos dígitos de dia, mes e ano é calculada e armazenada na variável total.
	if total > 16 {
		total = SomaDigitos(total)

	}
	fmt.Println("valor total é: ", total)

}
func SomaDigitos(numero int) int {
	soma := 0
	for numero > 0 {
		soma += numero % 10
		numero /= 10
	}
	return soma
}

//2°parte
// Pegar os dados iniciais do input e numerar os caracteres em posiçoes de 1 a 8
//func Impares_pares() {

//}

//apos isso somar somente os numeros impar e depois os numeros pares de acordo com cada posição
