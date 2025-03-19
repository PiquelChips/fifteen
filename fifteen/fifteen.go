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

func (vertex *FifteenVertex) GetAllEdges() []FifteenEdge {
    zeroX, zeroY := vertex.getZero()
    edges := []FifteenEdge{}

    if zeroX < 3 {
        edge := FifteenEdge{number: (*vertex)[zeroY][zeroX+1], direction: left}
        edges = append(edges, edge)
    }
    
    // down
    if zeroY > 0 {
        edge := FifteenEdge{number: (*vertex)[zeroY-1][zeroX], direction: down}
        edges = append(edges, edge)
    }

    // up
    if zeroY < 3 {
        edge := FifteenEdge{number: (*vertex)[zeroY+1][zeroX], direction: up}
        edges = append(edges, edge)
    }

    if zeroX > 0 {
        edge := FifteenEdge{number: (*vertex)[zeroY][zeroX-1], direction: right}
        edges = append(edges, edge)
    }

	return edges
}

// Return x and y coords of the zero
func (vertex *FifteenVertex) getZero() (int, int) {
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

    fmt.Printf("%d,%s %s\n", edge.number, space, edge.direction)
}
