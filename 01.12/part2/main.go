package part2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculateSimilarityScore calculates
// for each number in the left list multiply it by its occurrence count in the right list
func calculateSimilarityScore(left, right []int) int {
	// Create a map to count occurrences in right list
	rightCounts := make(map[int]int)
	for _, num := range right {
		rightCounts[num]++
	}

	// Calculate similarity score
	totalScore := 0
	for _, num := range left {
		// Multiply number by its occurrence count in right list
		totalScore += num * rightCounts[num]
	}

	return totalScore
}

// Run executes part 2 of the solution
func Run() {
	var left, right []int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		leftNum, _ := strconv.Atoi(fields[0])
		rightNum, _ := strconv.Atoi(fields[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	result := calculateSimilarityScore(left, right)
	fmt.Printf("Total similarity score: %d\n", result)
}
