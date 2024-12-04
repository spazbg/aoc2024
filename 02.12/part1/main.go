package part1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isSafe checks if a report follows the safety rules:
// numbers are increasing or decreasing
// adjacent number differences must be 1 or 2 or 3
func isSafe(numbers []int) bool {
	if len(numbers) < 2 {
		return true
	}

	// Determine if sequence is increasing or decreasing based on first two numbers
	isIncreasing := numbers[1] > numbers[0]

	// Check each adjacent pair
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]

		// Check if difference is within allowed range (1-3)
		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}

		// Check if sequence maintains its direction (up/down)
		if isIncreasing && diff < 0 {
			return false
		}
		if !isIncreasing && diff > 0 {
			return false
		}
	}

	return true
}

func Run() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		// Split line into numbers
		fields := strings.Fields(scanner.Text())
		numbers := make([]int, 0, len(fields))

		// Convert strings to integers
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Printf("Error converting %s to number: %v\n", field, err)
				continue
			}
			numbers = append(numbers, num)
		}

		// Check if this report is safe
		if isSafe(numbers) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}
