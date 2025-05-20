package algorithms

import (
"math/rand"
generator "mazesolver/Board"
)

type kruskalsAlgorithm struct{}

// NewKruskalsAlgorithm creates a new instance of Kruskal's algorithm
func NewKruskalsAlgorithm() Algorithm {
return &kruskalsAlgorithm{}
}

// DisjointSet represents a disjoint set data structure
type DisjointSet struct {
parent []int
rank   []int
}

// NewDisjointSet creates a new disjoint set
func NewDisjointSet(size int) *DisjointSet {
ds := &DisjointSet{
parent: make([]int, size),
rank:   make([]int, size),
}
for i := range ds.parent {
ds.parent[i] = i
}
return ds
}

// Find returns the root of the set containing x
func (ds *DisjointSet) Find(x int) int {
if ds.parent[x] != x {
ds.parent[x] = ds.Find(ds.parent[x]) // Path compression
}
return ds.parent[x]
}

// Union joins two sets together
func (ds *DisjointSet) Union(x, y int) {
rootX := ds.Find(x)
rootY := ds.Find(y)
if rootX != rootY {
if ds.rank[rootX] < ds.rank[rootY] {
ds.parent[rootX] = rootY
} else if ds.rank[rootX] > ds.rank[rootY] {
ds.parent[rootY] = rootX
} else {
ds.parent[rootY] = rootX
ds.rank[rootX]++
}
}
}

func (k *kruskalsAlgorithm) Generate(grid [][]generator.GridSpace, rng *rand.Rand) {
rows, cols := len(grid), len(grid[0])


// Create a list of all possible walls between cells
type wall struct {
x1, y1, x2, y2 int
}
walls := make([]wall, 0)

// Start from 0,0 for display coordinates
grid[0][0] = generator.Empty

// Add all possible walls to the list (only from odd-indexed cells)
for i := 1; i < rows-1; i += 2 {
for j := 1; j < cols-1; j += 2 {
// Mark cells as passages (except 0,0 which is already marked)
if i != 0 || j != 0 {
grid[i][j] = generator.Empty
}

// Add vertical wall if within bounds
if i+2 < rows-1 {
walls = append(walls, wall{i, j, i+2, j})
}
// Add horizontal wall if within bounds
if j+2 < cols-1 {
walls = append(walls, wall{i, j, i, j+2})
}
}
}

// Create disjoint set for tracking connected regions
ds := NewDisjointSet(rows * cols)

// Shuffle walls randomly
for i := len(walls) - 1; i > 0; i-- {
j := rng.Intn(i + 1)
walls[i], walls[j] = walls[j], walls[i]
}

// Process walls in random order
for _, w := range walls {
cell1 := w.y1 + w.x1*cols
cell2 := w.y2 + w.x2*cols

// If cells are not connected, connect them
if ds.Find(cell1) != ds.Find(cell2) {
ds.Union(cell1, cell2)
// Create passage
grid[(w.x1+w.x2)/2][(w.y1+w.y2)/2] = generator.Empty
}
}
}
