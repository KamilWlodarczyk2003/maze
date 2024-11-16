package main

type stack struct {
	list    [][]int
	visited [][]int
}

func (s *stack) stackPush(x int, y int) {
	s.list = append(s.list, []int{x, y})
}

func (s *stack) stackPop() []int { //zwraca pierwszy element i usuwa go
	var value []int = s.list[0] //zczytywanie pierwszej w kolejności wartości
	s.list = s.list[1:]         //usuwanie pierwszej wartości z listy
	return value                //zwrot wartości

}

func (s *stack) emptyCheck() bool { //sprawdza czy stos jest pusty
	if len(s.list) == 0 {
		return true
	} else {
		return false
	}
}
