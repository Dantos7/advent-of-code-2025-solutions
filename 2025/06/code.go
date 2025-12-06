package main

import (
	"log"
	"strconv"
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
	flippedMatrix := parseInput(input)
	total_value := 0

	for i := 0; i < len(flippedMatrix); i++ {
		line := flippedMatrix[i]
		operator := line[len(line)-1]
		operands := line[:len(line)-1]
		// fmt.Println(operator, operands)
		partial_result := 0
		switch operator {
		case "+":
			partial_result = sumStringifiedInts(operands)
		case "*":
			partial_result = multiplyStringifiedInts(operands)
		default:
			log.Fatal("Unknown operator:", operator)
		}
		// fmt.Println(partial_result)
		total_value += partial_result
	}

	return total_value
}

func parseInput(input string) [][]string {
	matrix := make([][]string, len(strings.Split(input, "\n")))
	for i, line := range strings.Split(input, "\n") {
		matrix[i] = strings.Fields(line) // strings.Fields ignores contiguous whitespaces
	}

	// Flip matrix for better efficiency (switch x and y coordinates)
	flippedMatrix := make([][]string, len(matrix[0]))
	for i := range flippedMatrix {
		flippedMatrix[i] = make([]string, len(matrix))
	}

	for i := range flippedMatrix {
		for j := 0; j < len(flippedMatrix[0]); j++ {
			flippedMatrix[i][j] = matrix[j][i]
		}
	}

	return flippedMatrix
}

func runPart2(input string) any {
	return "not implemented"
}

func sumStringifiedInts(operands []string) int {
	sum := 0
	for _, o := range operands {
		v, err := strconv.Atoi(o)
		if err != nil {
			log.Fatal(err)
		}
		sum += v
	}
	return sum
}

func multiplyStringifiedInts(operands []string) int {
	result := 1
	for _, o := range operands {
		v, err := strconv.Atoi(o)
		if err != nil {
			log.Fatal(err)
		}
		result *= v
	}
	return result
}
