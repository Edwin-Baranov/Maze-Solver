package algorithms

import (
"math/rand"
generator "mazesolver/Board"
)

type recursiveDivisionAlgorithm struct{}

// NewRecursiveDivisionAlgorithm creates a new instance of recursive division algorithm
func NewRecursiveDivisionAlgorithm() Algorithm {
return &recursiveDivisionAlgorithm{}
}

func (rd *recursiveDivisionAlgorithm) Generate(grid [][]generator.GridSpace, rng *rand.Rand) {
rows, cols := len(grid), len(grid[0])

// First fill everything with empty passages
for i := 0; i < rows; i++ {
for j := 0; j < cols; j++ {
grid[i][j] = generator.Empty
}
}

// Add border walls
for i := 0; i < rows; i++ {
grid[i][0] = generator.Wall
grid[i][cols-1] = generator.Wall
}
for j := 0; j < cols; j++ {
grid[0][j] = generator.Wall
grid[rows-1][j] = generator.Wall
}

// Start recursive division
rd.divide(grid, 1, 1, rows-2, cols-2, rng)

// Ensure start and end are clear
grid[0][0] = generator.Empty
grid[1][0] = generator.Empty
grid[0][1] = generator.Empty
}

func (rd *recursiveDivisionAlgorithm) divide(grid [][]generator.GridSpace, x1, y1, x2, y2 int, rng *rand.Rand) {
width := x2 - x1 + 1
height := y2 - y1 + 1

// Stop if the region is too small
if width < 3 || height < 3 {
return
}

// Choose orientation (horizontal or vertical wall)
horizontal := width < height
if width == height {
horizontal = rng.Intn(2) == 0
}

if horizontal {
// Choose where to place the wall (must be between cells)
wallY := y1 + 1 + 2*rng.Intn((height-1)/2)
if wallY >= y2 {
return
}

// Choose passage location (must be at a cell center)
passageX := x1 + 2*rng.Intn((width+1)/2)

// Build the wall
for x := x1; x <= x2; x++ {
if x != passageX {
grid[x][wallY] = generator.Wall
}
}

// Recursively divide regions
rd.divide(grid, x1, y1, x2, wallY-1, rng)
rd.divide(grid, x1, wallY+1, x2, y2, rng)

} else {
// Choose where to place the wall (must be between cells)
wallX := x1 + 1 + 2*rng.Intn((width-1)/2)
if wallX >= x2 {
return
}

// Choose passage location (must be at a cell center)
passageY := y1 + 2*rng.Intn((height+1)/2)

// Build the wall
for y := y1; y <= y2; y++ {
if y != passageY {
grid[wallX][y] = generator.Wall
}
}

// Recursively divide regions
rd.divide(grid, x1, y1, wallX-1, y2, rng)
rd.divide(grid, wallX+1, y1, x2, y2, rng)
}
}
