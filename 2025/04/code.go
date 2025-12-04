package main

import (
	"fmt"
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
	matrix := buildMatrix(input)
	counter := 0

	// For visual purposes only
	// // Copy the matrix to be able to mark removable positions
	// editableMatrix := make([][]rune, len(matrix))
	// for i := range matrix {
	// 	editableMatrix[i] = make([]rune, len(matrix[i]))
	// 	copy(editableMatrix[i], matrix[i])
	// }

	for i, row := range matrix {
		for j, c := range row {
			if c == '@' && isAccessible(matrix, i, j) {
				counter++
				// editableMatrix[i][j] = 'x' // For visualization purposes only
			}
		}
	}

	// For visualization purposes only
	// visualizeMatrix(matrix)
	// visualizeMatrix(editableMatrix)
	return counter
}

// https://dev.to/chigbeef_77/bool-int-but-stupid-in-go-3jb3
func Bool2int(b bool) int {
	// The compiler currently only optimizes this form.
	// See issue 6011.
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

func isAccessible(matrix [][]rune, i int, j int) bool {
	// A roll is accessible only if there are less than 4 rolls around it (adjacent columns)

	rows := len(matrix)
	cols := len(matrix[0])
	maximumRolls := 3
	roll_symbol := '@'

	return (Bool2int(i+1 < rows && matrix[i+1][j] == roll_symbol) +
		Bool2int(i+1 < rows && j+1 < cols && matrix[i+1][j+1] == roll_symbol) +
		Bool2int(i+1 < rows && j-1 >= 0 && matrix[i+1][j-1] == roll_symbol) +
		Bool2int(i-1 >= 0 && matrix[i-1][j] == roll_symbol) +
		Bool2int(i-1 >= 0 && j+1 < cols && matrix[i-1][j+1] == roll_symbol) +
		Bool2int(i-1 >= 0 && j-1 >= 0 && matrix[i-1][j-1] == roll_symbol) +
		Bool2int(j-1 >= 0 && matrix[i][j-1] == roll_symbol) +
		Bool2int(j+1 < cols && matrix[i][j+1] == roll_symbol)) <= maximumRolls
}

// For visualization purposes only
func visualizeMatrix(matrix [][]rune) {
	for _, row := range matrix {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
	fmt.Println()
}

func runPart2(input string) any {
	return "not implemented"
}

func buildMatrix(input string) [][]rune {
	lines := strings.Split(input, "\n")

	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix
}
