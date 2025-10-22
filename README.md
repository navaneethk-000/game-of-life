# GAME OF LIFE (Go implementation)

The Game of Life, also known as Conway's Game of Life, is a cellular automaton created by British mathematician John Horton Conway in 1970.

This project is a Go (Golang) implementation of the Game of Life , a zero-player simulation where the evolution of cells happens automatically based on simple rules.

Each cell in the grid is either alive or dead, and its next state depends on its neighboring cells according to these rules:

* Any live cell with fewer than two live neighbors dies (underpopulation).

* Any live cell with two or three live neighbors survives.

* Any live cell with more than three live neighbors dies (overpopulation).

* Any dead cell with exactly three live neighbors becomes alive (reproduction).

This Go implementation simulates the grid's evolution step by step in the terminal, showing how simple rules can create complex, dynamic patterns.