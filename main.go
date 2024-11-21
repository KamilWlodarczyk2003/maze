package main

import (
	"fmt"
)

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

	grid := [][]int{
		{9, 5, 5, 1, 3, 11},
		{8, 5, 3, 10, 8, 6},
		{8, 1, 4, 4, 0, 3},
		{8, 2, 11, 11, 8, 2},
		{8, 4, 2, 8, 4, 2},
		{12, 7, 12, 6, 13, 6},
	}

	dd.gridInit(10, 10, 15)
	dd.startInit(0, 0)
	dd.createMaze()
	//dd.grid = grid
	fmt.Println(dd.grid)
	dd.drawMaze(dd.grid)

	ast := astar{
		grid:       grid,
		start_pos:  [2]int{0, 0},
		finish_pos: [2]int{5, 5},
		open_list: stack{
			list: [][]int{},
		},
		closed_list: make(map[[2]int]bool),
		parent:      make(map[[2]int][2]int),
		g_cost:      make(map[[2]int]int),
	}
	fmt.Println(ast.manhEstimate([]int{2, 0}))
	ast.a_star_solving(dd.grid)
	ast.displaySolution(dd.grid)
	//fmt.Println(findNeighbors(1, 1, 3, 3))
	//fmt.Println(findNeighbors(2, 1, 3, 3))
	//fmt.Println(findNeighbors(1, 2, 3, 3))

	//ast.a_star_solving(dd.grid)

}
