package main

import (
	"fmt"
)

var dia, mes, ano int // Declaração de variáveis globais

// Função principal que recebe a data de nascimento, calcula a soma reduzida e chama a função SomaDigitos.
func main() {

	fmt.Print("Coloque sua Data de nascimento: ") // Solicita a data de nascimento ao usuário
	fmt.Scan(&dia, &mes, &ano)                    // Lê os valores digitados pelo usuário

	//calcula a soma das variaveis globais
	diaSoma := SomaDigitos(dia)
	mesSoma := SomaDigitos(mes)
	anoSoma := SomaDigitos(ano)

	total := diaSoma + mesSoma + anoSoma // Soma as somas individuais
	if total > 16 {                      // Se a soma total for maior que 16, calcula a soma novamente
		total = SomaDigitos(total)
	}
	fmt.Println("valor total é: ", total) // Imprime o valor final da soma reduzida
}

// Função que calcula a soma dos dígitos de um número inteiro.
func SomaDigitos(numero int) int {
	soma := 0        // Variável para armazenar a soma dos dígitos
	for numero > 0 { // Loop para percorrer os dígitos do número
		soma += numero % 10 // Adiciona o dígito atual à soma
		numero /= 10        // Divide o número por 10 para obter o próximo dígito
	}
	return soma // Retorna a soma dos dígitos
}

// 2°parte
// Pegar os dados iniciais do input e numerar os caracteres em posiçoes de 1 a 8
func Impares_pares(dataString string) {

}

//apos isso somar somente os numeros impar e depois os numeros pares de acordo com cada posição
