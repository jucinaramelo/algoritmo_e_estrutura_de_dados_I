package main

import (
	"fmt"
)

// Merge combina dois subarrays ordenados (e e d) em um único array ordenado (v) (conquista).
func merge(v[]int, e []int, d[]int){ // Complexidade O(n)
	indexV := 0
	indexE := 0
	indexD := 0

	// Laço principal: compara elementos de "e" e "d" e os insere em "v" em ordem crescente
	for indexE < len(e) && indexD < len(d){
		if e[indexE] <= d[indexD]{
			v[indexV] = e[indexE] // Adiciona o menor elemento de "e" a "v"
			indexE++ // Avança no subarray esquerdo
		} else{
			v[indexV] = d[indexD] // Adiciona o menor elemento de "d" a "v"
			indexD++
		}
		indexV++ // Avança no array principal "v"
	}

	// Adiciona quaisquer elementos restantes do subarray esquerdo "e" a "v"
	for indexE < len(e){
		v[indexV] = e[indexE]
		indexE++
		indexV++
	}


	// Adiciona quaisquer elementos restantes do subarray direito "d" a "v"
	for indexD < len(d){
		v[indexV] = d[indexD]
		indexD++
		indexV++
	}
}

// MergeSort ordena o array "v" usando o algoritmo Merge Sort (divisão).
func MergeSort(v[]int){ // Complexidade T(n)
	if len(v) > 1 {
		meio := len(v)/2

		// Cria o subarray esquerdo "e" com metade dos elementos
		tamE := meio
		e:= make([]int, tamE)
		for indexE := 0; indexE < tamE; indexE++{
			e[indexE] = v[indexE]
		}

		// Cria o subarray direito "d" com o restante dos elementos
		tamD := len(v) - tamE
		d := make([]int, tamD)
		for indexD := tamE; indexD < len(v); indexD++ {
			d[indexD-tamE] = v[indexD]
		}

		// Ordena recursivamente os subarrays esquerdo e direito
		MergeSort(e) // Complexidade T(n/2) para o lado esquerdo
		MergeSort(d) // Complexidade T(n/2) para o lado direito

		// Mescla os subarrays ordenados no array principal
		merge(v, e, d) // Complexidade O(n)

	}
}

func main(){
	v := []int{44, 43, 22, 5, 20, 26}

	fmt.Println("Vetor desordenado: ", v)

	MergeSort(v)

	fmt.Println("Vetor ordenado: ", v)
}
