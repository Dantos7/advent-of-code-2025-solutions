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

	// Assume position of S is fixed (0, m/2)
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
	// Perform visit according to rules
	// Avoid visiting the same branch twice by using the visited matrix

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
	// Strategy:
	// 	- Visit like in part one
	//  - From each visited location in the bottom, count the paths from a previous splitter
	//  - Proceed recursively (use caching to avoid double computation)
	// Complexity: O(2^(n+1)) time
	matrix := parseMatrix(input) // Size: n*m

	visited := make([][]bool, len(matrix))
	for i := range matrix {
		visited[i] = make([]bool, len(matrix[i]))
	}

	// use cache to avoid re-computing stuff
	cache := make([][]int, len(matrix))
	for i := range matrix {
		cache[i] = make([]int, len(matrix[i]))
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	// Assume position of S is fixed (0, m/2)
	start := [2]int{0, len(matrix[0]) / 2}
	visitAndCount(matrix, start, visited)

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

	totalPaths := countPaths(matrix, visited, cache)

	return totalPaths
}

func countPaths(matrix [][]rune, visited [][]bool, cache [][]int) int {
	// Non-recursive function that iterates over the last line and calls the recursive function to compute paths
	lastLineVisited := visited[len(visited)-1]
	totalPaths := 0
	for i, v := range lastLineVisited {
		if v {
			totalPaths += countEndPaths(matrix, visited, [2]int{len(visited) - 1, i}, cache)
		}
	}
	return totalPaths
}

func countEndPaths(matrix [][]rune, visited [][]bool, end [2]int, cache [][]int) int {
	// Recursive function counting from a single end the paths to it
	if cache[end[0]][end[1]] != -1 {
		return cache[end[0]][end[1]]
	}
	i := end[0] - 1
	j := end[1]
	paths := 0
	for i >= 0 {
		if visited[i][j] {
			if j < len(matrix[0])-1 && matrix[i][j+1] == '^' && visited[i-1][j+1] {
				paths += 1 * countEndPaths(matrix, visited, [2]int{i - 1, j + 1}, cache)
			}
			if j > 0 && matrix[i][j-1] == '^' && visited[i-1][j-1] {
				paths += 1 * countEndPaths(matrix, visited, [2]int{i - 1, j - 1}, cache)
			}
		} else if matrix[i][j] == 'S' {
			paths += 1
			break
		} else {
			break
		}
		i--
	}

	cache[end[0]][end[1]] = paths
	return paths
}
