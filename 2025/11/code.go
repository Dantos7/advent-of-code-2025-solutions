package main

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return runPart2(input)
	}
	// solve part 1 here
	return runPart1(input)
}

func runPart1(input string) any {
	vertices := parseInput(input)
	start := "you"

	count := countPathsToOut(start, vertices)
	return count
}

type Edge struct {
	From string
	To   string
}

func parseInput(input string) map[string][]Edge {
	vertices := make(map[string][]Edge)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, " ")
		id := parts[0][:len(parts[0])-1]
		edges := make([]Edge, 0, len(parts)-1)
		for _, to := range parts[1:] {
			edge := Edge{From: id, To: to}
			edges = append(edges, edge)
		}
		vertices[id] = edges
	}
	return vertices
}

func countPathsToOut(start string, vertices map[string][]Edge) int {
	count := 0
	if start == "out" {
		return 1
	}
	edges, ok := vertices[start]
	if !ok {
		// Dead end
		return 0
	} else {
		for _, e := range edges {
			count += countPathsToOut(e.To, vertices)
		}
	}
	return count
}

func runPart2(input string) any {
	return "not implemented"
}
