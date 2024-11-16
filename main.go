package main

import "fmt"

func makeGrid(x int, y int, val int) [][]int { //tworzy macierz o rozmiarach x y wypełniając ją wartościa val
	matrix := make([][]int, y)

	for v := range matrix {
		matrix[v] = make([]int, x)
	}

	for yv := range matrix {
		for i := 0; i < x; i++ {
			matrix[yv][i] = val
		}
	}

	return matrix
}

func main() {

	test1 := makeGrid(5, 2, 15)
	fmt.Println(test1)

	stos := &stack{
		list:    [][]int{},
		visited: [][]int{},
	}
	stos.stackPush(12, 12)
	stos.stackPush(15, 15)
	fmt.Println(stos.list)

	test := stos.stackPop()
	fmt.Println(test)
	fmt.Println(stos.list)

}
