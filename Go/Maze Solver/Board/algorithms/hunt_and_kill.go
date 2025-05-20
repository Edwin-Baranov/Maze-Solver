package algorithms

import (
"math/rand"
generator "mazesolver/Board"
)

type huntAndKillAlgorithm struct{}

// NewHuntAndKillAlgorithm creates a new instance of Hunt-and-Kill algorithm
func NewHuntAndKillAlgorithm() Algorithm {
return &huntAndKillAlgorithm{}
}

func (hk *huntAndKillAlgorithm) Generate(grid [][]generator.GridSpace, rng *rand.Rand) {
rows, cols := len(grid), len(grid[0])

// Track visited cells (only even-indexed positions)
visited := make([][]bool, rows)
for i := range visited {
visited[i] = make([]bool, cols)
}

// Start from (0,0)
currentX, currentY := 0, 0
visited[currentX][currentY] = true
grid[currentX][currentY] = generator.Empty

// Calculate total cells to visit (only even-indexed positions)
totalCells := ((rows + 1) / 2) * ((cols + 1) / 2)
visitedCount := 1

// Directions for movement (two steps at a time)
directions := [][2]int{{-2, 0}, {2, 0}, {0, -2}, {0, 2}} // Up, Down, Left, Right

// Continue until all cells have been visited
for visitedCount < totalCells {
// Try to walk randomly from current position
walking := true
for walking {
// Get unvisited neighbors
var unvisitedNeighbors [][2]int
for _, dir := range directions {
newX := currentX + dir[0]
newY := currentY + dir[1]
if newX >= 0 && newX < rows && newY >= 0 && newY < cols && !visited[newX][newY] {
unvisitedNeighbors = append(unvisitedNeighbors, [2]int{newX, newY})
}
}

if len(unvisitedNeighbors) > 0 {
// Choose random unvisited neighbor
next := unvisitedNeighbors[rng.Intn(len(unvisitedNeighbors))]
// Create passage
grid[next[0]][next[1]] = generator.Empty
grid[(currentX+next[0])/2][(currentY+next[1])/2] = generator.Empty
// Mark as visited
visited[next[0]][next[1]] = true
visitedCount++
// Move to the chosen cell
currentX, currentY = next[0], next[1]
} else {
walking = false
}
}

// Hunt mode - scan for an unvisited cell adjacent to a visited cell
foundNext := false
huntLoop:
for i := 0; i < rows; i += 2 { // Only check even indices
for j := 0; j < cols; j += 2 { // Only check even indices
if !visited[i][j] {
// Check each direction for a visited cell
for _, dir := range directions {
checkX := i + dir[0]
checkY := j + dir[1]
if checkX >= 0 && checkX < rows && checkY >= 0 && checkY < cols && visited[checkX][checkY] {
// Found a valid cell to continue from
currentX, currentY = i, j
visited[currentX][currentY] = true
visitedCount++
grid[currentX][currentY] = generator.Empty
grid[(currentX+checkX)/2][(currentY+checkY)/2] = generator.Empty
foundNext = true
break huntLoop
}
}
}
}
}

// If we couldn't find any cells in hunt mode, we must be done
if !foundNext {
break
}
}
}
