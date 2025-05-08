package gridMaze

import (
	"fmt"
	"math/rand"
	generator "mazesolver/Board"
)

type gridMaze struct {
	grid           [][]generator.GridSpace
	startX, startY int
	endX, endY     int
}

func (gm *gridMaze) GetGrid() [][]generator.GridSpace {
	return gm.grid
}

func (gm *gridMaze) GetStartPos() (int, int) {
	return gm.startX, gm.startY
}

func (gm *gridMaze) GetEndPos() (int, int) {
	return gm.endX, gm.endY
}

func (gm *gridMaze) initGrid(params ...int) {
	sizeX := defaultSize //TODO: investicate if needed due to params always having default values
	sizeY := defaultSize
	switch len(params) {
	case 2:
		// Add 2 to account for border walls
		sizeX = params[0]
		sizeY = params[1]
	}

	gm.grid = make([][]generator.GridSpace, sizeX)
	for i := range gm.grid {
		gm.grid[i] = make([]generator.GridSpace, sizeY)
	}
}

// Position represents a position in the grid
type Position struct {
	X, Y int
}

// Returns the current position from a flattened index
func (gm *gridMaze) toPosition(index int) Position {
	cols := len(gm.grid[0])
	return Position{
		X: index / cols,
		Y: index % cols,
	}
}

// Returns the flattened index from a position
func (gm *gridMaze) toIndex(pos Position) int {
	return pos.X*len(gm.grid[0]) + pos.Y
}

// GetNeighbors returns valid neighboring positions
func (gm *gridMaze) GetNeighbors(current int) []int {
	pos := gm.toPosition(current)
	neighbors := make([]int, 0, 4)
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // Up, Right, Down, Left

	for _, dir := range directions {
		newX, newY := pos.X+dir[0], pos.Y+dir[1]

		// Check bounds
		if newX < 0 || newX >= len(gm.grid) || newY < 0 || newY >= len(gm.grid[0]) {
			continue
		}

		// Check if walkable
		if gm.grid[newX][newY] != generator.Wall {
			neighbors = append(neighbors, gm.toIndex(Position{X: newX, Y: newY}))
		}
	}

	return neighbors
}

// GetDistanceToEnd returns the Manhattan distance to the end position
func (gm *gridMaze) GetDistanceToEnd(current int) float64 {
	pos := gm.toPosition(current)
	return float64(abs(pos.X-gm.endX) + abs(pos.Y-gm.endY))
}

// abs returns the absolute value of x
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

const (
	defaultSize = 64
)

type gridMazeGenerator struct{}

// NewGenerator creates a new gridMazeGenerator instance
func NewGenerator() generator.MazeGenerator {
	return &gridMazeGenerator{}
}

// generateMaze creates a maze using randomized DFS
func (gm *gridMaze) generateMaze(seed int64) {
	// Create a local random generator
	rng := rand.New(rand.NewSource(seed))

	// Start at top left corner
	gm.startX, gm.startY = 0, 0

	// Initialize end point based on grid dimensions being even or odd
	if len(gm.grid)%2 == 0 {
		gm.endX = len(gm.grid) - 2
	} else {
		gm.endX = len(gm.grid) - 1
	}

	if len(gm.grid[0])%2 == 0 {
		gm.endY = len(gm.grid[0]) - 2
	} else {
		gm.endY = len(gm.grid[0]) - 1
	}

	// Carve paths
	gm.carvePath(gm.startX, gm.startY, rng)

	// Ensure end is reachable
	gm.grid[gm.endX][gm.endY] = generator.Empty
}

// carvePath uses recursive backtracking to create the maze
func (gm *gridMaze) carvePath(x, y int, rng *rand.Rand) {
	// Set current cell as empty
	gm.grid[x][y] = generator.Empty

	// Define movement directions (up, right, down, left)
	directions := [][2]int{{-2, 0}, {0, 2}, {2, 0}, {0, -2}}
	rng.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	// Try each direction
	for _, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]

		// Check if the new position is within bounds and still a wall
		if newX >= 0 && newX < len(gm.grid) &&
			newY >= 0 && newY < len(gm.grid[0]) &&
			gm.grid[newX][newY] == generator.Wall {
			gm.grid[x+dir[0]/2][y+dir[1]/2] = generator.Empty
			gm.carvePath(newX, newY, rng)
		}
	}
}

func (gmg *gridMazeGenerator) GenerateRandomMaze(params ...int) generator.Maze {
	var maze *gridMaze
	switch len(params) {
	case 1:
		seed := params[0]
		maze = &gridMaze{}
		maze.initGrid()
		maze.generateMaze(int64(seed))
		return maze

	case 3:
		sizeX := params[0]
		sizeY := params[1]
		seed := params[2]
		maze = &gridMaze{}
		maze.initGrid(sizeX, sizeY)
		maze.generateMaze(int64(seed))
		return maze

	default:
		fmt.Println("Invalid number of parameters")
		return nil
	}
}
