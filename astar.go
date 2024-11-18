package main

import (
	"fmt"
	"math"
)

type astar struct {
	grid        [][]int
	start_pos   []int             //pozycja startowa
	finish_pos  [2]int            //wyjscie z labiryntu
	open_list   stack             //lista pól do przetworzenia
	closed_list map[[2]int]bool   //odwiedzone pola
	parent      map[[2]int][2]int //rodzic danego pola
	g_cost      map[[2]int]int    //koszt dojścia na pole
	//f_cost      map[[2]int]int
}

func (a *astar) manhEstimate(current []int) int { //ewaluacja kosztu dotarcia do celu
	return int(math.Abs(float64(current[1] - a.finish_pos[1] + current[0] - a.finish_pos[0])))
}

func (a *astar) a_star_solving(grid [][]int) {

	//inicjalizacja elementów struktury
	a.grid = grid

	a.start_pos = []int{0, 0}

	a.open_list = stack{
		list: [][]int{},
		//visited: make(map[[2]int]bool),
	}
	a.closed_list = make(map[[2]int]bool)
	a.parent = make(map[[2]int][2]int)
	a.g_cost = make(map[[2]int]int)
	//koniec inicjalizacji

	//f := a.manhEstimate(a.start_pos)
	a.open_list.pushWithValue(a.start_pos[0], a.start_pos[1], 0)
	a.g_cost[[2]int(a.start_pos)] = 0

	for !a.open_list.emptyCheck() {
		current := a.open_list.stackPop()
		current_pos := [2]int{current[0], current[1]}

		if current_pos == a.finish_pos {
			break
		}

		a.closed_list[current_pos] = true
		neighbors := findNeighbors(current[0], current[1], len(grid[0]), len(grid))
		//fmt.Println(a.grid)
		//fixedNghb := [][]int{}
		for _, vn := range neighbors { //sprawdzenie czy elementy są w visited
			if !a.closed_list[[2]int{vn[1], vn[0]}] {
				if isPassable(a.grid, vn[1], vn[0], current[0], current[1]) {
					a.g_cost[[2]int{vn[1], vn[0]}] = a.g_cost[[2]int{current_pos[0], current_pos[1]}] + 1
					f_cost := a.g_cost[[2]int{vn[1], vn[0]}] + a.manhEstimate([]int{vn[1], vn[0]})
					a.open_list.pushWithValue(vn[1], vn[0], f_cost)
				}
			}
		}

	}

	fmt.Println(a.g_cost)
}
