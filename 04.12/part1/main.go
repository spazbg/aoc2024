package part1

import (
	"fmt"
	"os"
	"strings"
)

const targetWord = "XMAS"

// Direction for search direction
type Direction struct {
	dRow, dCol int
}

// 8 search directions
var directions = []Direction{
	{0, 1},   // right
	{0, -1},  // left
	{1, 0},   // down
	{-1, 0},  // up
	{1, 1},   // diagonal down right
	{-1, -1}, // diagonal up left
	{1, -1},  // diagonal down left
	{-1, 1},  // diagonal up right
}

func Run() {
	// Read input file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	input := strings.TrimSpace(string(data))
	result := solvePart1(input)
	fmt.Printf("Result: %v\n", result)
}

func solvePart1(input string) int {
	// Split into lines
	grid := strings.Split(input, "\n")

	// Count total occurrences
	total := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			total += countWordsFromPosition(grid, row, col)
		}
	}

	return total
}

func countWordsFromPosition(grid []string, startRow, startCol int) int {
	count := 0

	// Try each direction from position
	for _, dir := range directions {
		if checkWord(grid, startRow, startCol, dir) {
			count++
		}
	}

	return count
}

func checkWord(grid []string, startRow, startCol int, dir Direction) bool {
	if startRow < 0 || startRow >= len(grid) || startCol < 0 || startCol >= len(grid[0]) {
		return false
	}

	// Check fit the word in this direction
	endRow := startRow + dir.dRow*(len(targetWord)-1)
	endCol := startCol + dir.dCol*(len(targetWord)-1)

	if endRow < 0 || endRow >= len(grid) || endCol < 0 || endCol >= len(grid[0]) {
		return false
	}

	// Check each char of the word
	for i := 0; i < len(targetWord); i++ {
		row := startRow + dir.dRow*i
		col := startCol + dir.dCol*i
		if grid[row][col] != targetWord[i] {
			return false
		}
	}

	return true
}
