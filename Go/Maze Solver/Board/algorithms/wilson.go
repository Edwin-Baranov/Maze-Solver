package algorithms

import (
"math/rand"
generator "mazesolver/Board"
)

type wilsonAlgorithm struct{}

// NewWilsonAlgorithm creates a new instance of Wilson's algorithm
func NewWilsonAlgorithm() Algorithm {
return &wilsonAlgorithm{}
}

type cell struct {
x, y int
}

func (w *wilsonAlgorithm) Generate(grid [][]generator.GridSpace, rng *rand.Rand) {
rows, cols := len(grid), len(grid[0])

// Track visited cells (only even-indexed positions)
visited := make([][]bool, rows)
for i := range visited {
visited[i] = make([]bool, cols)
}

// Start from (0,0)
visited[0][0] = true
grid[0][0] = generator.Empty

// Calculate total cells to visit (only even-indexed positions)
totalCells := ((rows + 1) / 2) * ((cols + 1) / 2)
visitedCount := 1

// Directions for movement (two steps at a time)
directions := [][2]int{{-2, 0}, {2, 0}, {0, -2}, {0, 2}} // Up, Down, Left, Right

// Continue until all even-indexed cells are connected
for visitedCount < totalCells {
// Pick a random unvisited cell as the start of a random walk
currentX, currentY := w.getRandomUnvisitedCell(visited, rng)
path := []cell{{currentX, currentY}}

// Perform loop-erased random walk until hitting a visited cell
for !visited[currentX][currentY] {
dir := directions[rng.Intn(len(directions))]
newX := currentX + dir[0]
newY := currentY + dir[1]

if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
// Check if this position creates a loop
var loopIndex int
eraseLoop := false
for i, p := range path {
if p.x == newX && p.y == newY {
eraseLoop = true
loopIndex = i
break
}
}

// If we found a loop, erase it by truncating the path
if eraseLoop {
path = path[:loopIndex+1]
} else {
path = append(path, cell{newX, newY})
}

currentX = newX
currentY = newY
}
}

// Add the path to the maze
for i := 0; i < len(path)-1; i++ {
curr := path[i]
next := path[i+1]
// Mark cells and the passage between them
grid[curr.x][curr.y] = generator.Empty
grid[(curr.x+next.x)/2][(curr.y+next.y)/2] = generator.Empty
grid[next.x][next.y] = generator.Empty
if !visited[curr.x][curr.y] {
visited[curr.x][curr.y] = true
visitedCount++
}
}
// Mark the last cell in the path if not visited
if !visited[path[len(path)-1].x][path[len(path)-1].y] {
visited[path[len(path)-1].x][path[len(path)-1].y] = true
visitedCount++
}
}
}

func (w *wilsonAlgorithm) getRandomUnvisitedCell(visited [][]bool, rng *rand.Rand) (int, int) {
rows, cols := len(visited), len(visited[0])
for {
// Only select even-indexed cells
x := rng.Intn((rows+1)/2) * 2
y := rng.Intn((cols+1)/2) * 2
if !visited[x][y] {
return x, y
}
}
}
