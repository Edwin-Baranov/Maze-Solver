package gridMaze

import (
	"github.com/stretchr/testify/assert"
	"mazesolver/Board"
	"testing"
)

func TestDefaultMazeSize(t *testing.T) {
	var generator generator.MazeGenerator = &gridMazeGenerator{}
	maze := generator.GenerateRandomMaze(0)

	gridMaze, _ := maze.(*gridMaze)

	assert.Equal(t, 64, len(gridMaze.grid))
	assert.Equal(t, 64, len(gridMaze.grid[0]))
}
