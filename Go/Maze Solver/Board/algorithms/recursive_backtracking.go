package algorithms

import (
"math/rand"
generator "mazesolver/Board"
)

type recursiveBacktrackingAlgorithm struct{}

// NewRecursiveBacktrackingAlgorithm creates a new instance of recursive backtracking algorithm
func NewRecursiveBacktrackingAlgorithm() Algorithm {
return &recursiveBacktrackingAlgorithm{}
}

func (rb *recursiveBacktrackingAlgorithm) Generate(grid [][]generator.GridSpace, rng *rand.Rand) {

// Start at top left corner (0,0) to match original implementation
rb.carve(grid, 0, 0, rng)
}

// carve uses recursive backtracking to create the maze
func (rb *recursiveBacktrackingAlgorithm) carve(grid [][]generator.GridSpace, x, y int, rng *rand.Rand) {
// Set current cell as empty
grid[x][y] = generator.Empty

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
if newX >= 0 && newX < len(grid) &&
newY >= 0 && newY < len(grid[0]) &&
grid[newX][newY] == generator.Wall {
grid[x+dir[0]/2][y+dir[1]/2] = generator.Empty
rb.carve(grid, newX, newY, rng)
}
}
}
