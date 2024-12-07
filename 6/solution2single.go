package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Position struct {
	row, col int
}

type State struct {
	position Position
	direction int
}

var directions = []Position{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

var grid [][]rune

func outOfBounds(position Position) bool {
	return position.row < 0 || position.row >= len(grid) || position.col < 0 || position.col >= len(grid[0])
}

func travel(start Position, newObstacle Position) bool {
	direction := 0
	positions := make(map[State]bool)
	positions[State{start, direction}] = true

	current := start

	for {
		nextPosition := Position{
			row: current.row + directions[direction].row,
			col: current.col + directions[direction].col,
		}

		// Check if the next position is out of bounds
		if outOfBounds(nextPosition) {
			return false
		}

		// Change directions if obstacle is found
		if grid[nextPosition.row][nextPosition.col] == '#' || (nextPosition == newObstacle) {
			direction = (direction + 1) % 4
			continue
		}

		// Move forward
		current = nextPosition
		state := State{current, direction}

		if positions[state] {
			return true
		}

		positions[state] = true
	}
}

func main() {
	startTime := time.Now()
	// Reading input from stdin
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			grid = append(grid, []rune(line))
		}
	}

	// Finding the starting point marked with '^'
	var start Position
	foundStart := false
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' {
				start = Position{i, j}
				foundStart = true
				break
			}
		}
		if foundStart {
			break
		}
	}

	// Count loops for different obstacles
	nrLoops := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '^' && grid[i][j] != '#' {
				newObstacle := Position{i, j}
				if travel(start, newObstacle) {
					nrLoops++
				}
			}
		}
	}

	fmt.Println(time.Since(startTime))
	fmt.Println(nrLoops)
}

