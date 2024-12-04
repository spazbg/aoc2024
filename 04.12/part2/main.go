package part2

import (
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	dRow, dCol int
}

var directions = []Direction{
	{-1, -1}, // up-left
	{-1, 1},  // up-right
	{1, -1},  // down-left
	{1, 1},   // down-right
}

func Run() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	input := strings.TrimSpace(string(data))
	fmt.Printf("Result: %v\n", solvePart2(input))
}

func solvePart2(input string) int {
	// split input into rows
	var grid []string
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		if line != "" {
			grid = append(grid, line)
		}
	}

	total := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'A' {
				total += findXPatterns(grid, row, col)
			}
		}
	}
	return total
}

func findXPatterns(grid []string, centerRow, centerCol int) int {
	if grid[centerRow][centerCol] != 'A' {
		return 0
	}

	// get sequences in each diagonal
	sequences := make(map[Direction]string)
	for _, dir := range directions {
		var seq strings.Builder
		// Get M-A-S or S-A-M in direction
		for i := -1; i <= 1; i++ {
			row := centerRow + dir.dRow*i
			col := centerCol + dir.dCol*i
			if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
				seq.Reset()
				break
			}
			seq.WriteByte(grid[row][col])
		}
		if seq.Len() == 3 {
			sequences[dir] = seq.String()
		}
	}

	// For each diagonal we need 1 valid MAS/SAM
	validDiagonals := 0
	for _, pair := range [][2]Direction{
		{{-1, -1}, {1, 1}}, // diagonal 1
		{{-1, 1}, {1, -1}}, // diagonal 2
	} {
		dir1, dir2 := pair[0], pair[1]
		seq1, ok1 := sequences[dir1]
		seq2, ok2 := sequences[dir2]
		if ok1 && ok2 && seq1[1] == 'A' && seq2[1] == 'A' {
			if isValidSequence(seq1) && isValidSequence(seq2) {
				validDiagonals++
			}
		}
	}

	// checking for 2 valid diagonals to form X
	if validDiagonals == 2 {
		return 1
	}
	return 0
}

func isValidSequence(seq string) bool {
	return seq == "MAS" || seq == "SAM"
}
