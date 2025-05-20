package algorithms

import (
"math/rand"
generator "mazesolver/Board"
)

// AlgorithmType represents different maze generation algorithms
type AlgorithmType int

const (
RecursiveBacktracking AlgorithmType = iota
Prims
Kruskals
RecursiveDivision
AldousBroder
Wilson
HuntAndKill
)

// Algorithm interface defines the contract for maze generation algorithms
type Algorithm interface {
// Generate creates a maze in the provided grid. The grid should be initialized
// with all walls (generator.Wall) before calling this method. The algorithm
// should respect the existing grid dimensions and structure (odd-indexed cells).
Generate(grid [][]generator.GridSpace, rng *rand.Rand)
}
