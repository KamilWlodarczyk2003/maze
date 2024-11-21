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

func (d *dfs) wallGen(x1, y1, x2, y2 int) { //generowanie ściań labiryntu [&= ^x zeruje bit o wartości x]
	if x2 == x1+1 { //wschód
		d.grid[y1][x1] &= ^2
		d.grid[y2][x2] &= ^8
	} else if x2 == x1-1 { //zachód
		d.grid[y1][x1] &= ^8
		d.grid[y2][x2] &= ^2
	} else if y2 == y1+1 { //południe
		d.grid[y1][x1] &= ^4
		d.grid[y2][x2] &= ^1
	} else if y2 == y1-1 { //północ
		d.grid[y1][x1] &= ^1
		d.grid[y2][x2] &= ^4
	}

}

func (d *dfs) drawMaze(grid [][]int) { //funkcja od chatagpt wizualizująca labirynt w konsoli
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

	for v := range matrix { //tworzenie macierzy
		matrix[v] = make([]int, x)
	}

	for yv := range matrix { //wypełnienie macierzy
		for i := 0; i < x; i++ {
			matrix[yv][i] = val
		}
	}

	d.grid = matrix
}

func (d *dfs) startInit(x int, y int) {
	d.start_pos = []int{y, x}
}

func findNeighbors(x, y, maxX, maxY int) [][]int { //znajduje sąsiadów podając koordynaty pola oraz maksymalny zakres maciezry

	var neighbors [][]int

	if y > 0 { // Góra
		neighbors = append(neighbors, []int{y - 1, x})
	}

	if y < maxY-1 { // Dół
		neighbors = append(neighbors, []int{y + 1, x})
	}

	if x > 0 { // Lewo
		neighbors = append(neighbors, []int{y, x - 1})
	}

	if x < maxX-1 { // Prawo
		neighbors = append(neighbors, []int{y, x + 1})
	}

	return neighbors
}

func isPassable(grid [][]int, x1, y1, x2, y2 int) bool {

	if x2-x1 == 1 { //ruch na wschód
		return (grid[y1][x1]&2 == 0) && (grid[y2][x2]&8 == 0)
	} else if x2-x1 == -1 { //ruch na zachód
		return (grid[y1][x1]&8 == 0) && (grid[y2][x2]&2 == 0)
	} else if y2-y1 == 1 { //ruch na południe
		return (grid[y1][x1]&4 == 0) && (grid[y2][x2]&1 == 0)
	} else if y2-y1 == -1 { //ruch na północ
		return (grid[y1][x1]&1 == 0) && (grid[y2][x2]&4 == 0)
	}

	return false
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

		//fmt.Println(current)
		neighbors := findNeighbors(current[1], current[0], len(d.grid[0]), len(d.grid))

		fixedNghb := [][]int{}
		for _, vn := range neighbors {
			if !stos.visited[[2]int{vn[0], vn[1]}] { //sprawdzenie czy elementy są w visited
				fixedNghb = append(fixedNghb, vn)
			}
		}
		if len(fixedNghb) > 0 {
			stos.stackPush(current[0], current[1])              //dodanie na stos
			shuffle(fixedNghb)                                  //tasowanie
			next := fixedNghb[0]                                //pierwsza wartość z przetasowanej listy
			d.wallGen(current[1], current[0], next[1], next[0]) //wyburzenie ściany między next a aktualnym polem
			stos.stackPush(next[0], next[1])                    //dodanie sąsiada do stosu
			stos.visited[[2]int{next[0], next[1]}] = true       //dodanie sąsiada do odwiedzonych
		}

		//last_val = current

	}

}
