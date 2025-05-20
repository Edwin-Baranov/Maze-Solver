package algorithms

import (
"math/rand"
generator "mazesolver/Board"
)

type aldousBroderAlgorithm struct{}

// NewAldousBroderAlgorithm creates a new instance of Aldous-Broder algorithm
func NewAldousBroderAlgorithm() Algorithm {
return &aldousBroderAlgorithm{}
}

func (ab *aldousBroderAlgorithm) Generate(grid [][]generator.GridSpace, rng *rand.Rand) {
rows, cols := len(grid), len(grid[0])

// Track visited cells (only for even-indexed positions)
visited := make([][]bool, rows)
for i := range visited {
visited[i] = make([]bool, cols)
}

// Start at (0,0)
currentX, currentY := 0, 0
visited[currentX][currentY] = true
grid[currentX][currentY] = generator.Empty

// Calculate total cells to visit (only even-indexed positions)
totalCells := ((rows + 1) / 2) * ((cols + 1) / 2)
visitedCount := 1

// Directions for movement (two steps at a time)
directions := [][2]int{{-2, 0}, {2, 0}, {0, -2}, {0, 2}} // Up, Down, Left, Right

// Continue until all even-indexed cells are visited
for visitedCount < totalCells {
// Choose random direction
dir := directions[rng.Intn(len(directions))]
newX := currentX + dir[0]
newY := currentY + dir[1]

// Check if the new position is valid (must stay on even indices)
if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
// If we haven't visited this cell yet
if !visited[newX][newY] {
// Mark cell as visited
visited[newX][newY] = true
visitedCount++

// Create passage by carving through the wall between cells
grid[newX][newY] = generator.Empty
grid[currentX+dir[0]/2][currentY+dir[1]/2] = generator.Empty
}

// Move to the new cell regardless of whether it was visited
currentX = newX
currentY = newY
}
}
}
