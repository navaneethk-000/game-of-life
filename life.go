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

// func CountNeighbours(grid Grid, x uint, y uint) uint {

// 	return 0
// }

func CountAliveNeighbours(grid Grid, x uint, y uint) uint {

	// size := int(grid.size)
	var count uint
	for row := int(x) - 1; row <= int(x)+1; row++ {

		for col := int(y) - 1; col <= int(y)+1; col++ {
			if row == int(x) && col == int(y) {
				continue
			}
			count++

		}
	}
	return count
}

func Rungeneration() {

}

func main() {

	newGrid := NewGrid(6)
	fmt.Println(DisplayGrid(newGrid))
	fmt.Println(CountAliveNeighbours(newGrid, 1, 1))
}
