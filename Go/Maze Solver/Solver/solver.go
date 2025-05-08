package solver

import "mazesolver/Board"

type Solution struct {
	Path    []int  // Sequence of positions forming the solution path
	Found   bool   // Whether a solution was found
	Message string // Additional information (error messages, etc.)
}

// PathFinder defines the interface for maze solving algorithms
type PathFinder interface {
	// SolveMaze attempts to find a path from start to end in the given maze
	// Returns a Solution containing the path if found
	SolveMaze(maze generator.Maze) Solution
}

// Algorithm represents different maze solving algorithms
type Algorithm int

const (
	DFS   Algorithm = iota // Depth-first search
	BFS                    // Breadth-first search
	AStar                  // A* search
)

// NewSolver creates a PathFinder implementation for the specified algorithm
func NewSolver(algorithm Algorithm) PathFinder {
	switch algorithm {
	case DFS:
		return &depthFirstSolver{}
	case BFS:
		return &breadthFirstSolver{}
	case AStar:
		return &aStarSolver{}
	default:
		return &depthFirstSolver{} // Default to DFS
	}
}

// Placeholder struct definitions - these will be implemented in separate files
type depthFirstSolver struct{}
type breadthFirstSolver struct{}
type aStarSolver struct{}

func (s *breadthFirstSolver) SolveMaze(maze generator.Maze) Solution {
	return Solution{Found: false, Message: "BFS not implemented yet"}
}

func (s *aStarSolver) SolveMaze(maze generator.Maze) Solution {
	return Solution{Found: false, Message: "A* not implemented yet"}
}
