package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coord struct {
	x, y int
}

func inBounds(x, y, height, width int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func solve1(height, width int, anntennas map[rune][]coord) int {
	result := 0

	antinodes := make(map[coord]bool)
	grid := make([][]rune, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]rune, width)
		for x := 0; x < width; x++ {
			grid[y][x] = '.' // Fill each cell with a dot
		}
	}

	for key := range anntennas {
		for i, coord1 := range anntennas[key] {
			for j := i + 1; j < len(anntennas[key]); j++ {
				coord2 := anntennas[key][j]

				if inBounds(coord1.x, coord1.y, height, width) && !antinodes[coord1] {
					antinodes[coord1] = true
					result++
				}
				if inBounds(coord2.x, coord2.y, height, width) && !antinodes[coord2] {
					antinodes[coord2] = true
					result++
				}

				dx := coord1.x - coord2.x
				dy := coord1.y - coord2.y

				x, y := coord1.x + dx, coord1.y + dy
				for inBounds(x, y, height, width) {
					if _, ok := antinodes[coord{x, y}]; !ok {
						grid[y][x] = '#'
						antinodes[coord{x, y}] = true
						result++
					}
					x, y = x + dx, y + dy
				}

				x, y = coord2.x - dx, coord2.y - dy
				for inBounds(x, y, height, width) {
					if _, ok := antinodes[coord{x, y}]; !ok {
						grid[y][x] = '#'
						antinodes[coord{x, y}] = true
						result++
					}
					x, y = x - dx, y - dy
				}

			}
		}
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}

	return result
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	anntennas := make(map[rune][]coord)
	var height, width int

	y := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) != 0 {
			width = len(line)
			for x, char := range line {
				if char != '.' {
					coord := coord{(x), (y)}
					anntennas[char] = append(anntennas[char], coord)
				}
			}
			y++
		}
	}
	height = y

	result := solve1(height, width, anntennas)

	fmt.Println(result)
}
