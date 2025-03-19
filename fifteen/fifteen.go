package fifteen

import "fmt"

var Solution FifteenVertex = [][]int{
	{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 10, 11, 12},
	{13, 14, 15, 0},
}

func (vertex *FifteenVertex) GetPotential(*FifteenVertex) {
}

func (vertex *FifteenVertex) GetNextVertex(edge FifteenVertex) FifteenVertex {
	return make(FifteenVertex, 0)
}

func (vertex *FifteenVertex) GetAllEdges() []FifteenVertex {
	return make([]FifteenVertex, 0)
}

func (vertex *FifteenVertex) GetZero() (int, int) {
	for y := range *vertex {
		for x, number := range (*vertex)[y] {
			if number == 0 {
				return x, y
			}
		}
	}

	return -1, -1
}

func (vertex *FifteenVertex) Print() {
	for y := range *vertex {
		for _, num := range (*vertex)[y] {
			if num < 10 {
				fmt.Printf("%d  ", num)
			} else {
				fmt.Printf("%d ", num)
			}
		}

		fmt.Printf("\n")
	}
}

func (edge *FifteenEdge) Print() {
    space := ""
    if edge.number < 10 {
        space = " "
    }

    fmt.Printf("%d%s, %s\n", edge.number, space, edge.direction)
}
