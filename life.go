package main

import (
	"fmt"
)

type Grid struct {
	size uint
	data [][]bool
}

func NewGrid(size uint) Grid {
	grid := Grid{
		size: size,
		data: make([][]bool, size),
	}

	for i := range grid.data {
		grid.data[i] = make([]bool, size)
	}
	return grid
}

func Display(grid Grid) string {

	return ""

}

func CountAliveNegighbours(grid Grid, x uint, y uint) uint {
	return 0
}

func Rungeneration() {

}

func main() {

	newGrid := NewGrid(uint(3))
	fmt.Println(newGrid)
	Display(newGrid)
	fmt.Println(Display(newGrid))
}
