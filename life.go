package main

import (
	"fmt"
	"strings"
)

type Grid struct {
	size uint
	data [][]bool
}

func NewGrid(size uint, live ...int) Grid {
	grid := make([][]bool, size)

	for i := 0; i < int(size); i++ {
		grid[i] = make([]bool, size)
	}
	for i := 0; i < len(live); i = i + 2 {
		x := live[i]
		y := live[i+1]
		grid[x][y] = true
	}

	return Grid{
		size: size,
		data: grid,
	}

}

func DisplayGrid(grid Grid) string {
	var builder strings.Builder
	for _, row := range grid.data {
		for _, i := range row {
			if i {
				builder.WriteString("* ")
			} else {
				builder.WriteString("x ")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()

}

func b2i(b bool) uint {
	if b {
		return 1
	}
	return 0
}

func CountAliveNeighbours(grid Grid, x uint, y uint) uint {
	s := grid.size - 1
	if grid.size == 1 {
		return uint(0)
	}

	//top left corner
	if x == 0 && y == 0 {
		return b2i(grid.data[x][y+1]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x+1][y+1])
	}

	//top right corner
	if x == 0 && y == s {
		return b2i(grid.data[x][y-1]) + b2i(grid.data[x+1][y-1]) + b2i(grid.data[x+1][y])
	}

	//bottom left corner
	if x == s && y == 0 {
		return b2i(grid.data[x-1][y]) + b2i(grid.data[x-1][y+1]) + b2i(grid.data[x][y+1])
	}

	//bottom right corner
	if x == s && y == s {
		return b2i(grid.data[x-1][y]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x][y-1])
	}

	//top edge
	if x == 0 {
		return b2i(grid.data[x][y-1]) + b2i(grid.data[x][y+1]) + b2i(grid.data[x+1][y-1]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x+1][y+1])
	}

	//bottom edge
	if x == s {
		return b2i(grid.data[x][y-1]) + b2i(grid.data[x][y+1]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x-1][y]) + b2i(grid.data[x-1][y+1])
	}

	//left edge
	if y == 0 {
		return b2i(grid.data[x-1][y]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x-1][y+1]) + b2i(grid.data[x][y+1]) + b2i(grid.data[x+1][y+1])
	}

	//right edge
	if y == s {
		return b2i(grid.data[x-1][y]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x][y-1]) + b2i(grid.data[x+1][y-1])
	}

	//middle
	return b2i(grid.data[x-1][y-1]) + b2i(grid.data[x-1][y]) + b2i(grid.data[x-1][y+1]) + b2i(grid.data[x][y-1]) + b2i(grid.data[x][y+1]) + b2i(grid.data[x+1][y-1]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x+1][y+1])
}

func RunGeneration(grid Grid) Grid {
	ret := NewGrid(grid.size)

	for i := 0; i < int(grid.size); i++ {
		for j := 0; j < int(grid.size); j++ {
			aliveNeighbours := CountAliveNeighbours(grid, uint(i), uint(j))

			if grid.data[i][j] {
				if aliveNeighbours < uint(2) {
					ret.data[i][j] = false
				} else if aliveNeighbours == 2 || aliveNeighbours == 3 {
					ret.data[i][j] = true
				} else if aliveNeighbours > uint(3) {
					ret.data[i][j] = false
				}
			}
		}

	}
	return ret
}

func main() {

	newGrid := NewGrid(6)
	fmt.Println(DisplayGrid(newGrid))
	fmt.Println(CountAliveNeighbours(newGrid, 1, 1))
}
