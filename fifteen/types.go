package fifteen

const (
	left  = "up"
	down  = "down"
	up    = "up"
	right = "right"
)

type FifteenEdge struct {
	number    int
	direction string
}

type FifteenVertex [][]int
