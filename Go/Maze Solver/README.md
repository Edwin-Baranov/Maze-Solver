# Golang Maze Solver

A maze generation and solving program implemented in Go that uses Depth-First Search (DFS) algorithm.

## Usage
```bash
go run main.go [flags]
```

Available flags:
- `-width`: Width of the maze (default: 20)
- `-height`: Height of the maze (default: 20)
- `-seed`: Random seed for maze generation (default: 0, generates random seed)

## Examples

```bash
# Generate default 20x20 maze
go run main.go

# Create a larger 30x40 maze
go run main.go -width 30 -height 40

# Generate specific maze using a seed
go run main.go -seed 12345
```

The program will:
1. Display the generated maze
2. Show the solution path using DFS algorithm
3. Print a summary with board size, seed used, and solution path length
