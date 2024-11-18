package main

import "sort"

type stack struct {
	list    [][]int
	visited map[[2]int]bool
}

func (s *stack) stackPush(x int, y int) { //dodaje element do stosu
	s.list = append(s.list, []int{x, y})
}

func (s *stack) stackPop() []int { //zwraca pierwszy element i usuwa go
	var value []int = s.list[len(s.list)-1]
	s.list = s.list[:(len(s.list) - 1)]
	//s.visited = append(s.visited, value)
	return value

}

func (s *stack) emptyCheck() bool { //sprawdza czy stos jest pusty
	if len(s.list) == 0 {
		return true
	} else {
		return false
	}
}

func (s *stack) pushWithValue(x int, y int, f int) {
	s.list = append(s.list, []int{x, y, f})

	sort.Slice(s.list, func(i, j int) bool {
		return s.list[i][2] > s.list[j][2] // Sortowanie według wartości f
	})
}
