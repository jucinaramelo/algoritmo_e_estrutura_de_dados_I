package main 

import (
	"fmt"
)

func SelectionSortIP(v []int){
	
	for varredura := 0; varredura < len(v)-1; varredura++ {
		iMenor := varredura

		for i := varredura +1; i < len(v); i++ {
			if v[i] < v[iMenor]{
				iMenor = i
			}
		}

		v[iMenor], v[varredura] = v[varredura], v[iMenor]
	}
}

func main (){
	v := []int{73, 23, 89, 33, 1, 2, 33}
	fmt.Println("Array antes da ordenação: ", v)

	SelectionSortIP(v)

	fmt.Println("Array após a ordenação: ", v)
}