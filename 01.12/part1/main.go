package part1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// CalculateTotalDistance calculates the total distance between sorted pairs
func CalculateTotalDistance(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}
		totalDistance += diff
	}
	return totalDistance
}

// Part1
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

	result := CalculateTotalDistance(left, right)
	fmt.Printf("Total distance: %d\n", result)
}
