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
	// Strategy: Visit the nodes according to the rules and count the splits
	// Complexity: O(2^(n+1)) time
	matrix := parseMatrix(input) // Size: n*m

	visited := make([][]bool, len(matrix))
	for i := range matrix {
		visited[i] = make([]bool, len(matrix[i]))
	}

	// Assume position of S is fixed (0, m/2+1)
	start := [2]int{0, len(matrix[0]) / 2}
	splits := visitAndCount(matrix, start, visited)

	// For visualization purposes - Print visited
	// for i, line := range visited {
	// 	for j, v := range line {
	// 		c := matrix[i][j]
	// 		if v && c == '.' {
	// 			fmt.Print("|")
	// 		} else {
	// 			fmt.Print(string(c))
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return splits
}

func parseMatrix(input string) [][]rune {
	// Parses the input and converts the string to a matrix of chars
	lines := strings.Split(input, "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = make([]rune, len(line))
		for j, c := range line {
			matrix[i][j] = c
		}
	}
	return matrix
}

func visitAndCount(matrix [][]rune, start [2]int, visited [][]bool) int {
	splits := 0
	i := start[0] + 1
	j := start[1]
	for i < len(matrix) {
		c := matrix[i][j]
		visited[i][j] = true
		switch c {
		case '.':
			// Case 1: empty spot -> continue downwards
			i++
		case '^':
			// Case 2: ^ -> split, continue on one side and recurse on the other
			// Avoid visiting twice the same branch
			if !visited[i][j+1] {
				visited[i][j+1] = true
				splits += visitAndCount(matrix, [2]int{i, j + 1}, visited)
			}
			if !visited[i][j-1] {
				splits += 1
				visited[i][j-1] = true
				j -= 1
			} else {
				return splits
			}
		default:
			log.Fatal("Unknown char encountered: ", string(c))
		}
	}
	return splits
}

func runPart2(input string) any {
	return "not implemented"
}
