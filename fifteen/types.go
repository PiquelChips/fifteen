package fifteen

const (
	left  = "Left"
	down  = "Down"
	up    = "Up"
	right = "Right"
)

type FifteenEdge struct {
	number    int
	direction string
}

type FifteenVertex [][]int
