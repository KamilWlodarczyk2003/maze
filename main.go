package main

/*
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
*/

func main() {

	//pole := makeGrid(10, 10, 15)
	//stos := stack{
	//list:    [][]int{},
	//visited: [][]int{},
	//}

	dd := dfs{
		grid:      [][]int{},
		start_pos: []int{},
	}

	dd.gridInit(10, 10, 15)
	dd.startInit(0, 0)
	dd.createMaze()
	//fmt.Println(dd.grid)
	dd.drawMaze(dd.grid)

}
