package algorithms

import (
"math/rand"
generator "mazesolver/Board"
)

type primsAlgorithm struct{}

// NewPrimsAlgorithm creates a new instance of Prim's algorithm
func NewPrimsAlgorithm() Algorithm {
return &primsAlgorithm{}
}

type wallCell struct {
x, y       int
fromX, fromY int
}

func (p *primsAlgorithm) Generate(grid [][]generator.GridSpace, rng *rand.Rand) {
rows, cols := len(grid), len(grid[0])


// Start at (0,0) for consistency with the display coordinates
startX, startY := 0, 0
grid[startX][startY] = generator.Empty

// Initialize walls list with walls adjacent to start
walls := make([]wallCell, 0)
if startX+2 < rows {
walls = append(walls, wallCell{x: startX+2, y: startY, fromX: startX, fromY: startY})
}
if startY+2 < cols {
walls = append(walls, wallCell{x: startX, y: startY+2, fromX: startX, fromY: startY})
}

// Continue until no walls remain
for len(walls) > 0 {
// Pick a random wall
wallIdx := rng.Intn(len(walls))
currentWall := walls[wallIdx]

// Remove the chosen wall from list (swap with last element and truncate)
walls[wallIdx] = walls[len(walls)-1]
walls = walls[:len(walls)-1]

// If the cell on the opposite side of the wall is still a wall
if grid[currentWall.x][currentWall.y] == generator.Wall {
// Create a passage
grid[currentWall.x][currentWall.y] = generator.Empty
grid[(currentWall.x+currentWall.fromX)/2][(currentWall.y+currentWall.fromY)/2] = generator.Empty

// Add adjacent walls
if currentWall.x+2 < rows && grid[currentWall.x+2][currentWall.y] == generator.Wall {
walls = append(walls, wallCell{x: currentWall.x+2, y: currentWall.y, fromX: currentWall.x, fromY: currentWall.y})
}
if currentWall.x-2 >= 0 && grid[currentWall.x-2][currentWall.y] == generator.Wall {
walls = append(walls, wallCell{x: currentWall.x-2, y: currentWall.y, fromX: currentWall.x, fromY: currentWall.y})
}
if currentWall.y+2 < cols && grid[currentWall.x][currentWall.y+2] == generator.Wall {
walls = append(walls, wallCell{x: currentWall.x, y: currentWall.y+2, fromX: currentWall.x, fromY: currentWall.y})
}
if currentWall.y-2 >= 0 && grid[currentWall.x][currentWall.y-2] == generator.Wall {
walls = append(walls, wallCell{x: currentWall.x, y: currentWall.y-2, fromX: currentWall.x, fromY: currentWall.y})
}
}
}
}
