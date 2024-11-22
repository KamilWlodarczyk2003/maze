package main

import (
	"fmt"
	"math"
	"slices"
)

type astar struct {
	grid         [][]int
	start_pos    [2]int            //pozycja startowa
	finish_pos   [2]int            //wyjscie z labiryntu
	open_list    stack             //lista pól do przetworzenia
	closed_list  map[[2]int]bool   //odwiedzone pola
	parent       map[[2]int][2]int //rodzic danego pola
	g_cost       map[[2]int]int    //koszt dojścia na pole
	path         [][2]int          //droga do celu
	visited_list [][2]int
	//f_cost      map[[2]int]int
}

func (a *astar) manhEstimate(current []int) int { //ewaluacja kosztu dotarcia do celu
	return int((math.Abs(float64(current[1]-a.finish_pos[1])) + math.Abs(float64(current[0]-a.finish_pos[0]))))
}

func (a *astar) a_star_solving(grid [][]int) {

	//inicjalizacja elementów struktury
	a.grid = grid

	a.start_pos = [2]int{0, 0}

	a.open_list = stack{
		list: [][]int{},
	}
	a.closed_list = make(map[[2]int]bool)
	a.parent = make(map[[2]int][2]int)
	a.g_cost = make(map[[2]int]int)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if !(x == a.start_pos[1] && y == a.start_pos[0]) {
				a.g_cost[[2]int{y, x}] = math.MaxInt
			}
		}
	}
	//koniec inicjalizacji

	a.open_list.pushWithValue(a.start_pos[0], a.start_pos[1], 0) //dodanie wartości startowej do listy
	a.g_cost[[2]int(a.start_pos)] = 0

	for !a.open_list.emptyCheck() {

		current := a.open_list.stackPop() //zczytanie wartości z listy
		a.visited_list = append(a.visited_list, [2]int(current))
		current_pos := [2]int{current[0], current[1]}

		if current_pos == a.finish_pos { //sprawdza czy dotarł do końca
			break
		}

		a.closed_list[current_pos] = true //dodaje pole do odwiedzonych
		neighbors := findNeighbors(current[0], current[1], len(grid[0]), len(grid))

		for _, vn := range neighbors {
			if !a.closed_list[[2]int{vn[1], vn[0]}] { //czy odwiedzony
				g_val_prop := a.g_cost[[2]int{current[0], current[1]}] + 1
				if isPassable(a.grid, vn[1], vn[0], current[0], current[1]) { //czy nie ma ściany
					if g_val_prop < a.g_cost[[2]int{vn[1], vn[0]}] || a.g_cost[[2]int{vn[1], vn[0]}] == 0 { //jesli powstanie lepsza cenowo droga to punktu to niech zaktualizuje
						a.g_cost[[2]int{vn[1], vn[0]}] = g_val_prop                //obliczenie kosztu g
						f_cost := g_val_prop + a.manhEstimate([]int{vn[1], vn[0]}) //obliczenie kosztu f
						a.open_list.pushWithValue(vn[1], vn[0], f_cost)            //dodanie do listy
						a.parent[[2]int{vn[1], vn[0]}] = current_pos               //dodanie rodzica
					}

				}
			}
		}

	}

	//tworzenie ścieżki zczytując rodziców
	a.path = [][2]int{}
	current_pole := a.finish_pos
	a.path = append(a.path, current_pole)
	for {
		current_pole = a.parent[current_pole]

		a.path = append(a.path, current_pole)
		if current_pole == a.start_pos {
			break
		}

	}
	slices.Reverse(a.path) //odwraca listę aby zaczynać od punktu 0,0
	//fmt.Println(a.g_cost)
	fmt.Println(a.path)
	fmt.Println(a.visited_list)
}

func (a *astar) displaySolution(grid [][]int) { //funkcja od chatgpt wizualizująca rozwiązanie labiryntu w konsoli
	rows := len(grid)
	cols := len(grid[0])

	// Tworzenie map odwiedzonych pól i ścieżki
	visited := make(map[[2]int]bool)
	for pos := range a.closed_list {
		visited[[2]int{pos[1], pos[0]}] = true
	}

	pathMap := make(map[[2]int]bool)
	for _, pos := range a.path {
		pathMap[[2]int{pos[1], pos[0]}] = true
	}

	// Rysowanie labiryntu z rozwiązaniem
	for y := 0; y < rows; y++ {
		// Rysowanie górnych ścian
		for x := 0; x < cols; x++ {
			fmt.Print("+")
			if grid[y][x]&1 != 0 { // Ściana na północ
				fmt.Print("---")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("+")

		// Rysowanie bocznych ścian i zawartości komórek
		for x := 0; x < cols; x++ {
			if grid[y][x]&8 != 0 { // Ściana na zachód
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}

			// Sprawdzanie zawartości pola
			pos := [2]int{y, x}
			if pos == a.start_pos {
				fmt.Print(" S ") // Punkt startowy
			} else if pos == a.finish_pos {
				fmt.Print(" E ") // Punkt końcowy
			} else if pathMap[pos] {
				fmt.Print(" * ") // Ścieżka rozwiązania
			} else if visited[pos] {
				fmt.Print(" x ") // Odwiedzone pole
			} else {
				fmt.Print("   ") // Puste pole
			}
		}
		// Zamknięcie prawej ściany
		fmt.Println("|")
	}

	// Rysowanie dolnych ścian ostatniego wiersza
	for x := 0; x < cols; x++ {
		fmt.Print("+")
		if grid[rows-1][x]&4 != 0 { // Ściana na południe
			fmt.Print("---")
		} else {
			fmt.Print("   ")
		}
	}
	fmt.Println("+")
}
