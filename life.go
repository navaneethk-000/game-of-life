package main

import (
	"fmt"
)

type Grid struct {
	size uint
	data [][]bool
}

func NewArray(size uint) Grid {
	grid := Grid{
		size: size,
		data: make([][]bool, size),
	}

	return grid

}

func Display(grid Grid) string {

	return ""

}

func CountAliveNegighbours(grid, x, y) uint {
	//corner

	//edge
	//middle

	return count
}

func Rungeneration() {

}

func main() {

	newGrid := NewArray(size)
	Display(newGrid)
	fmt.Println(Display(newGrid))
}
