package main

import (
	"github.com/PiquelChips/fifteen/fifteen"
)

func main() {
	vertex := fifteen.FifteenVertex{
		{0, 6, 7, 8},
		{9, 11, 1, 5},
		{2, 13, 5, 15},
		{4, 10, 12, 3},
	}

	vertex.Print()

	edges := vertex.GetAllEdges()

	for _, edge := range edges {
        vertex.GetNextVertex(edge).Print()
	}
}
