package main

type stack struct {
	list    [][]int
	visited [][]int
}

func (s *stack) stackPush(x int, y int) { //dodaje element do stosu
	s.list = append(s.list, []int{x, y})
}

func (s *stack) stackPop() []int { //zwraca pierwszy element i usuwa go
	var value []int = s.list[0]
	s.list = s.list[1:]
	s.visited = append(s.visited, value)
	return value

}

func (s *stack) emptyCheck() bool { //sprawdza czy stos jest pusty
	if len(s.list) == 0 {
		return true
	} else {
		return false
	}
}