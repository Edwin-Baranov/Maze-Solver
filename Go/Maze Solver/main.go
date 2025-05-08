package main

import (
	"flag"
	"fmt"
	"math/rand"
	"mazesolver/Board"
	"mazesolver/Board/GridMaze"
	"mazesolver/Display"
	"mazesolver/Solver"
	"time"
)

func main() {
	// Define command-line flags
	width := flag.Int("width", 20, "Width of the maze")
	height := flag.Int("height", 20, "Height of the maze")
	seed := flag.Int64("seed", 0, "Random seed for maze generation (0 for random)")
	flag.Parse()

	// If seed is 0, generate a random seed
	if *seed == 0 {
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		*seed = rng.Int63()
	}

	// Create a maze generator
	var mazeGen generator.MazeGenerator = gridMaze.NewGenerator()
	maze := mazeGen.GenerateRandomMaze(*width, *height, int(*seed))

	// Create a console display
	display := display.NewDisplay(display.Console)

	fmt.Printf("Generated Maze (size: %dx%d, seed: %d):\n", *width, *height, *seed)
	if err := display.DisplayMaze(maze); err != nil {
		fmt.Printf("Error displaying maze: %v\n", err)
		return
	}

	// Create a DFS solver
	solver := solver.NewSolver(solver.DFS)

	// Solve the maze
	solution := solver.SolveMaze(maze)

	fmt.Println("\nSolution:")
	if err := display.DisplaySolution(maze, solution); err != nil {
		fmt.Printf("Error displaying solution: %v\n", err)
		return
	}

	// Display summary information
	fmt.Printf("\nSummary:\n")
	fmt.Printf("- Algorithm: DFS (Depth-First Search)\n")
	fmt.Printf("- Board Size: %dx%d\n", *width, *height)
	fmt.Printf("- Seed: %d\n", *seed)
	if solution.Found {
		fmt.Printf("- Path Length: %d steps\n", len(solution.Path))
	}
}
