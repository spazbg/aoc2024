package part1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// findValidMultiplications finds all valid mul(X,Y) and return their products
func findValidMultiplications(line string) []int {
	// Regular expression to match valid mul(X,Y)
	// - "mul" string
	// - opening parenthesis
	// - 1-3 digits first number
	// - comma
	// - 1-3 digits second number
	// - closing parenthesis
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(line, -1)

	// result of multiplication
	products := make([]int, 0, len(matches))
	for _, match := range matches {
		// string to int
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		// add to results
		products = append(products, x*y)
	}

	return products
}

func Run() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Find all valid multiplications and sum their products
		products := findValidMultiplications(line)
		for _, product := range products {
			totalSum += product
		}
	}

	fmt.Printf("Sum of all multiplication results: %d\n", totalSum)
}
