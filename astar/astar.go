package astar

import "fmt"

type Vertex interface {
	GetPotential(end *Vertex)
	GetAllEdges() []interface{}
}

func AStar(start *Vertex, end *Vertex) {
	fmt.Printf("Please implement A*")
}
