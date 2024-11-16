package main

import (
	"fmt"
	"math/rand"
)

type dfs struct {
	grid      [][]int
	start_pos []int
}

/*	Pomysł chatagpt

1. Reprezentacja binarna ścian
Każda komórka może być reprezentowana jako liczba całkowita, gdzie każdy bit odpowiada obecności ściany w określonym kierunku:

Bit 0: Ściana na północ (N).
Bit 1: Ściana na wschód (E).
Bit 2: Ściana na południe (S).
Bit 3: Ściana na zachód (W).

Przykład:

1111 (15): Wszystkie ściany są obecne.
1011 (11): Brak ściany na południu.
0000 (0): Brak ścian (wolna komórka).
*/

func (d *dfs) wallGen(x1, y1, x2, y2 int) { // zasady powstawania ścian
	if x2 == (x1 + 1) {
		if d.grid[y1][x1]&4 != 0 { //sprawdzenie czy dany bit nie został już usunięty
			d.grid[y1][x1] -= 4
		}
		if d.grid[y2][x2]&1 != 0 {
			d.grid[y2][x2] -= 1
		}
	}
	if x2 == (x1 - 1) {
		if d.grid[y1][x1]&1 != 0 {
			d.grid[y1][x1] -= 1
		}
		if d.grid[y2][x2]&4 != 0 {
			d.grid[y2][x2] -= 4
		}
	}
	if y2 == (y1 + 1) {
		if d.grid[y1][x1]&2 != 0 {
			d.grid[y1][x1] -= 2
		}
		if d.grid[y2][x2]&8 != 0 {
			d.grid[y2][x2] -= 8
		}
	}
	if y2 == (y1 - 1) {
		if d.grid[y1][x1]&8 != 0 {
			d.grid[y1][x1] -= 8
		}
		if d.grid[y2][x2]&2 != 0 {
			d.grid[y2][x2] -= 2
		}
	}
}

func (d *dfs) drawMaze(grid [][]int) { //Funkcja od chatgpt do wizualizacji labiryntu w konsoli
	rows := len(grid)
	cols := len(grid[0])

	// Iteruj przez wiersze
	for y := 0; y < rows; y++ {
		// Rysowanie górnej części komórek (poziome ściany)
		for x := 0; x < cols; x++ {
			fmt.Print("+")
			if grid[y][x]&8 != 0 { // Ściana na północ
				fmt.Print("---")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("+")

		// Rysowanie bocznych ścian i wnętrza komórek
		for x := 0; x < cols; x++ {
			if grid[y][x]&1 != 0 { // Ściana na zachód
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print("   ") // Wnętrze komórki
		}
		fmt.Println("|")
	}

	// Rysowanie dolnej części ostatniego wiersza
	for x := 0; x < cols; x++ {
		fmt.Print("+")
		if grid[rows-1][x]&2 != 0 { // Ściana na południe
			fmt.Print("---")
		} else {
			fmt.Print("   ")
		}
	}
	fmt.Println("+")
}

func (d *dfs) gridInit(x int, y int, val int) { //tworzy macierz o rozmiarach x y wypełniając ją wartościa val
	matrix := make([][]int, y)

	for v := range matrix {
		matrix[v] = make([]int, x)
	}

	for yv := range matrix {
		for i := 0; i < x; i++ {
			matrix[yv][i] = val
		}
	}

	d.grid = matrix
}

func (d *dfs) startInit(x int, y int) {
	d.start_pos = make([]int, 2)
	d.start_pos[0] = y
	d.start_pos[1] = x
}

func findNeighbors(x, y, maxX, maxY int) [][]int { //Funkcja od ChatGPT, znajduje sąsiadów na podstawie aktualnych x,y
	// Lista sąsiadów
	var neighbors [][]int

	// Góra
	if y > 0 {
		neighbors = append(neighbors, []int{y - 1, x})
	}
	// Dół
	if y < maxY-1 {
		neighbors = append(neighbors, []int{y + 1, x})
	}
	// Lewo
	if x > 0 {
		neighbors = append(neighbors, []int{y, x - 1})
	}
	// Prawo
	if x < maxX-1 {
		neighbors = append(neighbors, []int{y, x + 1})
	}

	return neighbors
}

func shuffle(slice [][]int) { //funkcja od ChatGPT, przetasowuje slice
	if len(slice) <= 1 {
		return
	}

	// Przetasowanie slice
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func (d *dfs) createMaze() {
	stos := stack{
		list:    [][]int{},
		visited: [][]int{},
	}
	stos.stackPush(d.start_pos[0], d.start_pos[1])
	//current := []int{0, 0}
	last_val := []int{-1, -1}

	for !stos.emptyCheck() {
		//fmt.Println("lista:", stos.list)
		//fmt.Println("visited:", stos.visited)

		current := stos.stackPop()
		if last_val[0] != -1 && last_val[1] != -1 {
			d.wallGen(current[1], current[0], last_val[1], last_val[0])
		}

		//fmt.Println(current)
		neighbors := findNeighbors(current[1], current[0], len(d.grid[0]), len(d.grid))

		filteredNeighbors := [][]int{}
		for _, vn := range neighbors { //sprawdzenie czy elementy są w visited
			isVisited := false
			for _, vs := range stos.visited {
				if vn[0] == vs[0] && vn[1] == vs[1] {
					isVisited = true
					break
				}
			}
			for _, vs := range stos.list { //sprawdzenie czy elementy są już na stosie
				if vn[0] == vs[0] && vn[1] == vs[1] {
					isVisited = true
					break
				}
			}
			if !isVisited {
				filteredNeighbors = append(filteredNeighbors, vn)
			}
		}
		neighbors = filteredNeighbors

		shuffle(neighbors)

		for _, v := range neighbors {
			stos.stackPush(v[0], v[1])
		}

		last_val = current

	}

}
