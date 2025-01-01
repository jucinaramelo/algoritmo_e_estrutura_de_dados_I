package main

import (
	"fmt"
)

func InsertionSortIP(v []int) {

	for insercao := 1; insercao < len(v); insercao++ {
		for i := insercao; i >= 1 && v[i] < v[i-1]; i-- {
			v[i], v[i-1] = v[i-1], v[i]
		}
	}
}

func main() {
	v := []int{23, 45, 98, 23, 21, 5}

	fmt.Println("Array desordenado: ", v)

	InsertionSortIP(v)

	fmt.Println("Array ordenado: ", v)
}
