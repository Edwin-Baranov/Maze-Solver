package generator

type Maze interface {
	GetNeighbors(current int) []int // Returns valid neighboring positions

	GetDistanceToEnd(current int) float64 // Returns estimated distance to end position
}

type MazeGenerator interface {
	GenerateRandomMaze(params ...int) Maze
}
