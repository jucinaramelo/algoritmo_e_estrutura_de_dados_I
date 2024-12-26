package main

import (
	"fmt"
)

func BubbleSortIP(v []int) {
	for varredura := 1; varredura < len(v); varredura++ {
		trocou := false

		for i := 0; i < len(v)-varredura; i++ {

			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				trocou = true
			}
		}
		if !trocou {
			return
		}
	}
}

func main() {
	v := []int{73, 23, 45, 98, 23, 21, 5}

	fmt.Println("Array desordenado: ", v)

	BubbleSortIP(v)

	fmt.Println("Array ordenado: ", v)
}
