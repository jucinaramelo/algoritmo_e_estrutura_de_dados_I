// Implementação do Counting Sort em Go com base nos passos fornecidos
package main

import (
	"fmt"
)

// Função countingSort implementa o algoritmo Counting Sort
type sortType = int

func countingSort(v []sortType) []sortType {
	// PASSO 1: Encontrar o maior e o menor valor no array
	min, max := v[0], v[0] // Inicializa min e max como o primeiro elemento do array
	for _, num := range v {
		// Percorre todos os elementos do array 'v'
		// '_' ignora o índice, e 'num' representa o valor atual
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}

	// PASSO 2: Criar o vetor de contagem (count)
	size := max - min + 1 // Tamanho baseado no intervalo entre min e max
	count := make([]int, size)

	// PASSO 3: Contar as ocorrências de cada elemento em v
	for _, num := range v {
		// Para cada elemento 'num' em 'v', incrementa sua contagem
		// O índice no vetor de contagem é ajustado subtraindo 'min'
		count[num-min]++
	}

	// PASSO 4: Realizar a soma cumulativa no vetor de contagem
	for i := 1; i < len(count); i++ {
		// Soma cumulativa: acumula o valor atual com o anterior
		count[i] += count[i-1]
	}

	// PASSO 5: Criar vetor para armazenar os elementos ordenados
	sorted := make([]sortType, len(v))

	// PASSO 6: Mapear os elementos do vetor original para o vetor ordenado
	//          Percorrendo o array original de trás para frente (para garantir estabilidade)
	for i := len(v) - 1; i >= 0; i-- {
		// Calcula o índice no vetor ordenado usando o vetor de contagem
		index := count[v[i]-min] - 1 // Encontrar a posição correta
		sorted[index] = v[i]         // Coloca o elemento na posição correta
		count[v[i]-min]--            // Decrementa o contador para o próximo elemento igual
	}

	return sorted // Retorna o vetor ordenado
}

func main() {
	// Exemplo de uso do algoritmo Counting Sort
	v := []sortType{5, 2, 1, 7, 3, 18, 4, 9, 10, 3}
	fmt.Println("Array original:", v)

	// Chama a função countingSort
	sortedArr := countingSort(v)
	fmt.Println("Array ordenado:", sortedArr)
}
