package main

import Maze "maze/src/Maze"

func main() {
	rows, cols := 25, 25
	maze := Maze.GenerateMaze(rows, cols)
	Maze.PrintMaze(maze)
}
