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
