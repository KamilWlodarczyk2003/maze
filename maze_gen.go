package main

import (
	"fmt"
	"math/rand"
)

type dfs struct {
	grid      [][]int
	start_pos []int
}

/*	Pomysł od chatagpt

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

func (d *dfs) wallGen(x1, y1, x2, y2 int) {
	// Ruch w poziomie (wschód/zachód)
	if x2 == x1+1 { // Ruch na wschód
		d.grid[y1][x1] &= ^2 // Usuń ścianę wschodnią z bieżącej komórki
		d.grid[y2][x2] &= ^8 // Usuń ścianę zachodnią z sąsiedniej komórki
	} else if x2 == x1-1 { // Ruch na zachód
		d.grid[y1][x1] &= ^8 // Usuń ścianę zachodnią z bieżącej komórki
		d.grid[y2][x2] &= ^2 // Usuń ścianę wschodnią z sąsiedniej komórki
	}

	// Ruch w pionie (północ/południe)
	if y2 == y1+1 { // Ruch na południe
		d.grid[y1][x1] &= ^4 // Usuń ścianę południową z bieżącej komórki
		d.grid[y2][x2] &= ^1 // Usuń ścianę północną z sąsiedniej komórki
	} else if y2 == y1-1 { // Ruch na północ
		d.grid[y1][x1] &= ^1 // Usuń ścianę północną z bieżącej komórki
		d.grid[y2][x2] &= ^4 // Usuń ścianę południową z sąsiedniej komórki
	}
}

func (d *dfs) drawMaze(grid [][]int) {
	rows := len(grid)
	cols := len(grid[0])

	// Iteruj przez wiersze
	for y := 0; y < rows; y++ {
		// Rysowanie górnych ścian komórek (poziome ściany)
		for x := 0; x < cols; x++ {
			fmt.Print("+")
			if grid[y][x]&1 != 0 { // Ściana na północ
				fmt.Print("---")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("+")

		// Rysowanie bocznych ścian i wnętrza komórek
		for x := 0; x < cols; x++ {
			if grid[y][x]&8 != 0 { // Ściana na zachód
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print("   ") // Wnętrze komórki
		}
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
	d.start_pos = []int{y, x}
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

func isPassable(grid [][]int, x1, y1, x2, y2 int) bool {
	// Debugging - pokaż wartości komórek w formacie binarnym
	//fmt.Printf("Current cell [%d, %d]: %b, Next cell [%d, %d]: %b\n", x1, y1, grid[y1][x1], x2, y2, grid[y2][x2])

	// Oblicz kierunek ruchu
	dx := x2 - x1
	dy := y2 - y1

	// Sprawdź przechodność na podstawie kierunku
	switch {
	case dx == 1: // Ruch na wschód
		return (grid[y1][x1]&2 == 0) && (grid[y2][x2]&8 == 0)
	case dx == -1: // Ruch na zachód
		return (grid[y1][x1]&8 == 0) && (grid[y2][x2]&2 == 0)
	case dy == 1: // Ruch na południe
		return (grid[y1][x1]&4 == 0) && (grid[y2][x2]&1 == 0)
	case dy == -1: // Ruch na północ
		return (grid[y1][x1]&1 == 0) && (grid[y2][x2]&4 == 0)
	default:
		// Jeśli nie ma ruchu między sąsiadującymi polami, zwróć false
		return false
	}
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

func (d *dfs) createMaze() { //tworzenie labiryntu
	stos := stack{
		list:    [][]int{},
		visited: make(map[[2]int]bool),
	}
	stos.stackPush(d.start_pos[0], d.start_pos[1])
	//current := []int{0, 0}
	//last_val := []int{-1, -1}

	for !stos.emptyCheck() {
		//fmt.Println("lista:", stos.list)
		//fmt.Println("visited:", stos.visited)

		current := stos.stackPop()
		//if last_val[0] != -1 && last_val[1] != -1 {
		//d.wallGen(current[1], current[0], last_val[1], last_val[0])
		//}

		//fmt.Println(current)
		neighbors := findNeighbors(current[1], current[0], len(d.grid[0]), len(d.grid))

		fixedNghb := [][]int{}
		for _, vn := range neighbors { //sprawdzenie czy elementy są w visited
			if !stos.visited[[2]int{vn[0], vn[1]}] {
				fixedNghb = append(fixedNghb, vn)
			}
		}
		if len(fixedNghb) > 0 {
			stos.stackPush(current[0], current[1])
			shuffle(fixedNghb)
			next := fixedNghb[0]
			d.wallGen(current[1], current[0], next[1], next[0])
			stos.stackPush(next[0], next[1])
			stos.visited[[2]int{next[0], next[1]}] = true
		}

		//last_val = current

	}

}
