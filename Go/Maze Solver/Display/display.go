package display

import (
	"fmt"
	"mazesolver/Board"
	"mazesolver/Solver"
)

// DisplayMethod represents different ways to display the maze
type DisplayMethod int

const (
	Console DisplayMethod = iota // Console-based ASCII display
	GUI                          // Graphical display (to be implemented)
)

// MazeDisplayer defines the interface for displaying mazes and their solutions
type MazeDisplayer interface {
	// DisplayMaze shows the raw maze without a solution
	DisplayMaze(maze generator.Maze) error

	// DisplaySolution shows the maze with the solution path
	DisplaySolution(maze generator.Maze, solution solver.Solution) error
}

// NewDisplay creates a MazeDisplayer implementation for the specified method
func NewDisplay(method DisplayMethod) MazeDisplayer {
	switch method {
	case Console:
		return &consoleDisplay{}
	case GUI:
		return &guiDisplay{}
	default:
		return &consoleDisplay{} // Default to console display
	}
}

// Placeholder struct definitions - these will be implemented in separate files
type consoleDisplay struct{}
type guiDisplay struct{}

// ANSI color codes and characters for dark background terminal
const (
	red    = "\033[1;31m" // bright red
	white  = "\033[1;37m" // bright white
	grey   = "\033[1;90m" // bright black (grey)
	purple = "\033[1;35m" // bright magenta
	green  = "\033[1;32m" // bright green
	// bgBlack = "\033[40m"   // black background
	reset   = "\033[0m" // reset all attributes
	wall    = "██"
	path    = "  "
	solPath = "██"
)

func (d *consoleDisplay) DisplayMaze(maze generator.Maze) error {
	// Try to convert to GridMaze interface
	gridMaze, ok := maze.(generator.GridMaze)
	if !ok {
		return fmt.Errorf("maze does not implement GridMaze interface")
	}

	grid := gridMaze.GetGrid()
	startX, startY := gridMaze.GetStartPos()
	startX, startY = startX+1, startY+1
	endX, endY := gridMaze.GetEndPos()
	endX, endY = endX+1, endY+1
boardSizeX, boardSizeY := len(grid)+2, len(grid[0])+2
for j := 0; j < boardSizeY; j++ {
for i := 0; i < boardSizeX; i++ {
			// Check if it's a border wall
			if i == 0 || i == boardSizeX-1 || j == 0 || j == boardSizeY-1 {
				fmt.Print(grey + wall + reset)
			} else {
				if i == startX && j == startY {
					fmt.Print(purple + wall + reset)
				} else if i == endX && j == endY {
					fmt.Print(green + wall + reset)
				} else if grid[i-1][j-1] == generator.Wall {
					fmt.Print(path + reset)
				} else { // Empty
					fmt.Print(white + wall + reset)
				}
			}
		}
		fmt.Println()
	}
	return nil
}

func (d *consoleDisplay) DisplaySolution(maze generator.Maze, solution solver.Solution) error {
	gridMaze, ok := maze.(generator.GridMaze)
	if !ok {
		return fmt.Errorf("maze does not implement GridMaze interface")
	}

	if !solution.Found {
		fmt.Println(solution.Message)
		return nil
	}

	// Create a map of solution path positions for quick lookup
	pathMap := make(map[int]bool)
	for _, pos := range solution.Path {
		pathMap[pos] = true
	}

	grid := gridMaze.GetGrid()
	startX, startY := gridMaze.GetStartPos()
	startX, startY = startX+1, startY+1
	endX, endY := gridMaze.GetEndPos()
	endX, endY = endX+1, endY+1
boardSizeX, boardSizeY := len(grid)+2, len(grid[0])+2
for j := 0; j < boardSizeY; j++ {
for i := 0; i < boardSizeX; i++ {
			// Check if it's a border wall
			if i == 0 || i == boardSizeX-1 || j == 0 || j == boardSizeY-1 {
				fmt.Print(grey + wall + reset)
			} else {
				if i == startX && j == startY {
					fmt.Print(purple + wall + reset)
				} else if i == endX && j == endY {
					fmt.Print(green + wall + reset)
				} else if grid[i-1][j-1] == generator.Wall {
					fmt.Print(path + reset)
				} else if pathMap[(i-1)*len(grid[0])+(j-1)] {
					fmt.Print(red + solPath + reset)
				} else {
					fmt.Print(white + wall + reset)
				}
			}
		}
		fmt.Println()
	}
	return nil
}

func (d *guiDisplay) DisplayMaze(maze generator.Maze) error {
	return nil // TODO: Implement GUI display
}

func (d *guiDisplay) DisplaySolution(maze generator.Maze, solution solver.Solution) error {
	return nil // TODO: Implement GUI solution display
}
