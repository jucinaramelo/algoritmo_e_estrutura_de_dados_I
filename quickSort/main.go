package main

import (
    "fmt"
    "math/rand"
    "time"
)

func QuickSort(v []int, ini int, fim int) {
	if ini < fim {
		indexPivot := partition(v, ini, fim)
		QuickSort(v, ini, indexPivot-1)
		QuickSort(v, indexPivot+1, fim)
	}
}

func partition(v []int, ini int, fim int) int {
	// Semente aleatória baseada no tempo atual
	rand.Seed(time.Now().UnixNano())

	// Escolher um índice aleatório entre ini e fim
	pivotIndex := rand.Intn(fim-ini+1) + ini

	// Trocar o pivô aleatório com o último elemento
	v[pivotIndex], v[fim] = v[fim], v[pivotIndex]

	// Agora o pivô está na posição fim
	indexPivot := fim
	pIndex := ini
	for i := ini; i < indexPivot; i++ {
		if v[i] <= v[indexPivot] {
			v[i], v[pIndex] = v[pIndex], v[i]
			pIndex++
		}
	}
	v[pIndex], v[indexPivot] = v[indexPivot], v[pIndex]
	return pIndex
}

func main() {
    // Array de exemplo para ordenar
    v := []int{10, 7, 8, 9, 1, 5, 2, 6}
    fmt.Println("Array original:", v)

    // Chama o QuickSort para ordenar o array
    QuickSort(v, 0, len(v)-1)

    // Imprime o array ordenado
    fmt.Println("Array ordenado:", v)
}
