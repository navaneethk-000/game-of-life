package main

import (
	"reflect"
	"testing"
)

func TestNewGrid(t *testing.T) {
	size := uint(3)
	actual := NewGrid(size)

	expected := Grid{
		size: 3,
		data: [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error")
	}

	if actual.size != expected.size {
		t.Errorf("Expected size : %d but got %d", expected.size, actual.size)
	}

}

func TestDisplayGrid(t *testing.T) {
	size := 3
	newGrid := NewGrid(uint(size))
	newGrid.data[0][0] = true
	actual := DisplayGrid(newGrid)
	expected :=
		"* x x \n" +
			"x x x \n" +
			"x x x \n"

	if actual != expected {
		t.Errorf("Expected \n%v but got \n%v", expected, actual)
	}
}
func TestCountAliveNeighbourTopLeftCorner(t *testing.T) {
	grid := NewGrid(5, 0, 1)
	actual := CountAliveNeighbours(grid, 0, 0)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = NewGrid(5, 0, 1, 1, 1)
	actual = CountAliveNeighbours(grid, 0, 0)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = NewGrid(5, 0, 1, 1, 1, 1, 0)
	actual = CountAliveNeighbours(grid, 0, 0)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighbourTopRightCorner(t *testing.T) {
	grid := NewGrid(5, 0, 3)
	actual := CountAliveNeighbours(grid, 0, 4)

	if actual != 1 {
		t.Errorf("Expected neighbour count : 1, but got %d", actual)
	}
	grid = NewGrid(5, 0, 3, 1, 3)
	actual = CountAliveNeighbours(grid, 0, 4)
	if actual != 2 {
		t.Errorf("Expected neighbour count : 2, but got %d", actual)
	}

	grid = NewGrid(5, 0, 3, 1, 3, 1, 4)
	actual = CountAliveNeighbours(grid, 0, 4)
	if actual != 3 {
		t.Errorf("Expected neighbour count : 3, but got %d", actual)
	}
}

func TestCountAliveNeighbourBottomLeftCorner(t *testing.T) {
	grid := NewGrid(5, 3, 0)
	actual := CountAliveNeighbours(grid, 4, 0)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = NewGrid(5, 3, 0, 3, 1)
	actual = CountAliveNeighbours(grid, 4, 0)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = NewGrid(5, 3, 0, 3, 1, 4, 1)
	actual = CountAliveNeighbours(grid, 4, 0)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighbourBottomRightCorner(t *testing.T) {
	grid := NewGrid(5, 4, 3)
	actual := CountAliveNeighbours(grid, 4, 4)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = NewGrid(5, 4, 3, 3, 3)
	actual = CountAliveNeighbours(grid, 4, 4)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = NewGrid(5, 4, 3, 3, 3, 3, 4)
	actual = CountAliveNeighbours(grid, 4, 4)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighboursTopEdge(t *testing.T) {
	grid := NewGrid(5, 0, 1)
	actual := CountAliveNeighbours(grid, 0, 2)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = NewGrid(5, 0, 1, 1, 1)
	actual = CountAliveNeighbours(grid, 0, 2)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = NewGrid(5, 0, 1, 1, 1, 1, 2)
	actual = CountAliveNeighbours(grid, 0, 2)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = NewGrid(5, 0, 1, 1, 1, 1, 2, 1, 3)
	actual = CountAliveNeighbours(grid, 0, 2)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = NewGrid(5, 0, 1, 1, 1, 1, 2, 1, 3, 0, 3)
	actual = CountAliveNeighbours(grid, 0, 2)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}
}

func TestCountAliveNeighboursBottomEdge(t *testing.T) {
	grid := NewGrid(5, 4, 1)
	actual := CountAliveNeighbours(grid, 4, 2)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = NewGrid(5, 4, 1, 3, 1)
	actual = CountAliveNeighbours(grid, 4, 2)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = NewGrid(5, 4, 1, 3, 1, 3, 2)
	actual = CountAliveNeighbours(grid, 4, 2)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = NewGrid(5, 4, 1, 3, 1, 3, 2, 3, 3)
	actual = CountAliveNeighbours(grid, 4, 2)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = NewGrid(5, 4, 1, 3, 1, 3, 2, 3, 3, 4, 3)
	actual = CountAliveNeighbours(grid, 4, 2)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}
}

func TestCountAliveNeighboursLeftEdge(t *testing.T) {
	grid := NewGrid(5, 1, 0)
	actual := CountAliveNeighbours(grid, 2, 0)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = NewGrid(5, 1, 0, 1, 1)
	actual = CountAliveNeighbours(grid, 2, 0)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = NewGrid(5, 1, 0, 1, 1, 2, 1)
	actual = CountAliveNeighbours(grid, 2, 0)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = NewGrid(5, 1, 0, 1, 1, 2, 1, 3, 1)
	actual = CountAliveNeighbours(grid, 2, 0)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = NewGrid(5, 1, 0, 1, 1, 2, 1, 3, 1, 3, 0)
	actual = CountAliveNeighbours(grid, 2, 0)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}
}

func TestCountAliveNeighboursRightEdge(t *testing.T) {
	grid := NewGrid(5, 1, 4)
	actual := CountAliveNeighbours(grid, 2, 4)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = NewGrid(5, 1, 4, 1, 3)
	actual = CountAliveNeighbours(grid, 2, 4)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = NewGrid(5, 1, 4, 1, 3, 2, 3)
	actual = CountAliveNeighbours(grid, 2, 4)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = NewGrid(5, 1, 4, 1, 3, 2, 3, 3, 3)
	actual = CountAliveNeighbours(grid, 2, 4)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = NewGrid(5, 1, 4, 1, 3, 2, 3, 3, 3, 3, 4)
	actual = CountAliveNeighbours(grid, 2, 4)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}
}

func TestCountAliveNeighboursMiddle(t *testing.T) {
	grid := NewGrid(5, 1, 2)
	actual := CountAliveNeighbours(grid, 2, 2)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = NewGrid(5, 1, 2, 1, 3)
	actual = CountAliveNeighbours(grid, 2, 2)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = NewGrid(5, 1, 2, 1, 3, 2, 3)
	actual = CountAliveNeighbours(grid, 2, 2)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = NewGrid(5, 1, 2, 1, 3, 2, 3, 3, 3)
	actual = CountAliveNeighbours(grid, 2, 2)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = NewGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2)
	actual = CountAliveNeighbours(grid, 2, 2)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}

	grid = NewGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2, 3, 1)
	actual = CountAliveNeighbours(grid, 2, 2)
	if actual != 6 {
		t.Errorf("Expected neighbour count to be 6, but got %d", actual)
	}

	grid = NewGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2, 3, 1, 2, 1)
	actual = CountAliveNeighbours(grid, 2, 2)
	if actual != 7 {
		t.Errorf("Expected neighbour count to be 7, but got %d", actual)
	}

	grid = NewGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2, 3, 1, 2, 1, 1, 1)
	actual = CountAliveNeighbours(grid, 2, 2)
	if actual != 8 {
		t.Errorf("Expected neighbour count to be 8, but got %d", actual)
	}
}

func TestRule1(t *testing.T) {
	grid := NewGrid(4, 1, 1, 0, 0)
	actualNewGrid := RunGeneration(grid)
	expectedNewGrid := NewGrid(4)
	if !reflect.DeepEqual(actualNewGrid, expectedNewGrid) {
		t.Errorf("Expected %v but got %v!\n", expectedNewGrid, actualNewGrid)
	}
}

func TestRule2(t *testing.T) {
	grid := NewGrid(5, 1, 2, 2, 1, 2, 2, 3, 2)
	actualNewGrid := RunGeneration(grid)
	expectedNewGrid := NewGrid(5, 1, 2, 2, 1, 2, 2, 3, 2)
	if !reflect.DeepEqual(actualNewGrid, expectedNewGrid) {
		t.Errorf("Expected %v but got %v!\n", expectedNewGrid, actualNewGrid)
	}
}
