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
	size := 6
	newGrid := NewGrid(uint(size))
	newGrid.data[0][0] = true
	actual := DisplayGrid(newGrid)
	expected :=
		"* x x \n" +
			"x x x \n" +
			"x x x \n"

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
