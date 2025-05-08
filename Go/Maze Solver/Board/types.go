package generator

// GridSpace represents a cell in the maze grid
type GridSpace uint8

const (
	Wall GridSpace = iota
	Empty
)

// GridMaze defines interface for grid-based mazes
type GridMaze interface {
	Maze
	GetGrid() [][]GridSpace
	GetStartPos() (int, int)
	GetEndPos() (int, int)
}
