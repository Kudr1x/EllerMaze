package maze

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell struct {
	row, col int
}

func GenerateMaze(rows, cols int) [][]bool {
	maze := make([][]bool, rows)
	for i := range maze {
		maze[i] = make([]bool, cols)
		for j := range maze[i] {
			maze[i][j] = true
		}
	}

	rand.Seed(time.Now().UnixNano())

	startCell := Cell{0, 0}
	stack := []Cell{startCell}
	maze[startCell.row][startCell.col] = false

	for len(stack) > 0 {
		currentCell := stack[len(stack)-1]
		neighbors := getUnvisitedNeighbors(currentCell, maze)

		if len(neighbors) > 0 {
			nextCell := neighbors[rand.Intn(len(neighbors))]
			removeWall(currentCell, nextCell, maze)
			stack = append(stack, nextCell)
			maze[nextCell.row][nextCell.col] = false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return maze
}

func getUnvisitedNeighbors(cell Cell, maze [][]bool) []Cell {
	directions := []struct{ row, col int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var neighbors []Cell

	for _, dir := range directions {
		nextRow, nextCol := cell.row+dir.row*2, cell.col+dir.col*2

		if nextRow >= 0 && nextRow < len(maze) && nextCol >= 0 && nextCol < len(maze[0]) && maze[nextRow][nextCol] {
			neighbors = append(neighbors, Cell{nextRow, nextCol})
		}
	}

	return neighbors
}

func removeWall(currentCell, nextCell Cell, maze [][]bool) {
	wallRow, wallCol := (currentCell.row+nextCell.row)/2, (currentCell.col+nextCell.col)/2
	maze[wallRow][wallCol] = false
}

func PrintMaze(maze [][]bool) {
	for _, row := range maze {
		for _, cell := range row {
			if cell {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
