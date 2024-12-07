package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Position struct {
	row, col int
}

type State struct {
	position  Position
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

func travel(start Position, newObstacle Position, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

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
			return
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
			resultChan <- 1
			return
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
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' {
				start = Position{i, j}
				break
			}
		}
	}

	// Count loops for different obstacles
	resultChan := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '^' && grid[i][j] != '#' {
				wg.Add(1)
				go travel(start, Position{i, j}, resultChan, &wg)
			}
		}
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	nrLoops := 0
	for result := range resultChan {
		nrLoops += result
	}
	
	fmt.Println("Execution time: ", time.Since(startTime))
	fmt.Println(nrLoops)
}
