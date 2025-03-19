package astar

import "fmt"

type Vertex interface {
	GetPotential(end *Vertex)
	GetAllEdges() []Edge
    Print()
}

type Edge interface {
    Print()
}

func AStar(start *Vertex, end *Vertex) {
	fmt.Printf("Please implement A*")
}
