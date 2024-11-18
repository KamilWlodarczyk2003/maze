package main

import "fmt"

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

	dd.gridInit(3, 3, 15)
	dd.startInit(0, 0)
	dd.createMaze()
	fmt.Println(dd.grid)
	dd.drawMaze(dd.grid)

	ast := astar{
		grid:       dd.grid,
		start_pos:  []int{0, 0},
		finish_pos: [2]int{1, 1},
		open_list: stack{
			list: [][]int{},
		},
		closed_list: make(map[[2]int]bool),   // Initialize closed list as empty
		parent:      make(map[[2]int][2]int), // Initialize parent map as empty
		g_cost:      make(map[[2]int]int),    // Initialize g_cost map
	}
	ast.a_star_solving(dd.grid)
	fmt.Println(findNeighbors(1, 1, 3, 3))
	fmt.Println(findNeighbors(2, 1, 3, 3))
	fmt.Println(findNeighbors(1, 2, 3, 3))

	//ast.a_star_solving(dd.grid)

}
