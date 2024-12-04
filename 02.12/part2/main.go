package part2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isSafe checks if a sequence follows the safety rules
func isSafe(numbers []int) bool {
	if len(numbers) < 2 {
		return true
	}

	isIncreasing := numbers[1] > numbers[0]

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]

		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}

		if isIncreasing && diff < 0 {
			return false
		}
		if !isIncreasing && diff > 0 {
			return false
		}
	}

	return true
}

// isSafeWithDampener checks if a sequence is safe or can be made safe by removing one number
func isSafeWithDampener(numbers []int) bool {
	// check if its already safe
	if isSafe(numbers) {
		return true
	}

	// Try removing each number one at a time
	for i := range numbers {
		// Create a new slice without the i el
		dampened := make([]int, 0, len(numbers)-1)
		dampened = append(dampened, numbers[:i]...)
		dampened = append(dampened, numbers[i+1:]...)

		if isSafe(dampened) {
			return true
		}
	}

	return false
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

		// Check if this report is safe with dampener
		if isSafeWithDampener(numbers) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports with Problem Dampener: %d\n", safeCount)
}
