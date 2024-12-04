package part2

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func findValidMultiplicationsWithState(content string) int {
	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	type command struct {
		typ  string // mul do don't
		pos  int
		x, y int
	}

	var commands []command

	// Find all multiplications
	mulMatches := mulRe.FindAllStringSubmatchIndex(content, -1)
	for _, match := range mulMatches {
		x, _ := strconv.Atoi(content[match[2]:match[3]])
		y, _ := strconv.Atoi(content[match[4]:match[5]])
		pos := match[0]
		commands = append(commands, command{
			typ: "mul",
			pos: pos,
			x:   x,
			y:   y,
		})
	}

	// Find all do()
	doMatches := doRe.FindAllStringIndex(content, -1)
	for _, match := range doMatches {
		pos := match[0]
		commands = append(commands, command{
			typ: "do",
			pos: pos,
		})
	}

	// Find all dont()
	dontMatches := dontRe.FindAllStringIndex(content, -1)
	for _, match := range dontMatches {
		pos := match[0]
		commands = append(commands, command{
			typ: "dont",
			pos: pos,
		})
	}

	// Sort by position
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].pos < commands[j].pos
	})

	sum := 0
	enabled := true // enabled by default

	for _, cmd := range commands {
		switch cmd.typ {
		case "do":
			enabled = true
		case "dont":
			enabled = false
		case "mul":
			if enabled {
				sum += cmd.x * cmd.y
			}
		}
	}

	return sum
}

func Run() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sum := findValidMultiplicationsWithState(string(content))
	fmt.Printf("Final sum of all multiplication results: %d\n", sum)
}
