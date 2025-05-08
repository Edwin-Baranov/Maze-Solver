package solver

import "mazesolver/Board"

// depthFirstSolver implements a depth-first search algorithm for maze solving
func (s *depthFirstSolver) SolveMaze(maze generator.Maze) Solution {
	visited := make(map[int]bool)
	path := make([]int, 0)
	// Get starting position
	gm, ok := maze.(generator.GridMaze)
	if !ok {
		return Solution{
			Found:   false,
			Message: "Maze does not implement GridMaze interface",
		}
	}
	startX, startY := gm.GetStartPos()
	cols := len(gm.GetGrid()[0])
	current := startX*cols + startY // Convert to flattened index

	if s.dfs(maze, current, visited, &path) {
		return Solution{
			Path:    path,
			Found:   true,
			Message: "Solution found using DFS",
		}
	}

	return Solution{
		Found:   false,
		Message: "No solution found using DFS",
	}
}

// dfs performs recursive depth-first search
func (s *depthFirstSolver) dfs(maze generator.Maze, current int, visited map[int]bool, path *[]int) bool {
	// Mark current position as visited
	visited[current] = true
	*path = append(*path, current)

	// Check if we're at the end (distance to end is 0)
	if maze.GetDistanceToEnd(current) == 0 {
		return true
	}

	// Try each neighbor
	for _, neighbor := range maze.GetNeighbors(current) {
		if !visited[neighbor] {
			if s.dfs(maze, neighbor, visited, path) {
				return true
			}
		}
	}

	// If no path found, backtrack by removing the current position
	*path = (*path)[:len(*path)-1]
	return false
}
