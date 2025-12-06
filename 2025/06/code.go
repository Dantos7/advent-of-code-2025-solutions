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
	flippedMatrix := parseInput1(input)
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

func parseInput1(input string) [][]string {
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
	flippedMatrix := parseInput2(input)
	total_value := 0

	for i := 0; i < len(flippedMatrix); i++ {
		line := flippedMatrix[i]
		operator := line[len(line)-1]
		operands := line[:len(line)-1]
		partial_result := 0
		switch strings.TrimSpace(operator) {
		case "+":
			partial_result = sumStringifiedInts2nd(operands) // noqa: typos
		case "*":
			partial_result = multiplyStringifiedInts2nd(operands) // noqa: typos
		default:
			log.Fatal("Unknown operator:", operator)
		}
		// fmt.Println(partial_result)
		total_value += partial_result
	}

	return total_value
}

func parseInput2(input string) [][]string {
	// Compared to parseInput1, it doesn't ignore whitespaces
	// Use operators positions to identify start and end of each column
	matrix := make([][]string, len(strings.Split(input, "\n")))
	paddedLines := padLinesRight(strings.Split(input, "\n"))
	operators_line := paddedLines[len(matrix)-1]
	for i, line := range paddedLines {
		matrix[i] = make([]string, 0)
		start_j := 0
		end_j := findEndOfColumn(operators_line, start_j+1)
		for start_j < end_j {
			matrix[i] = append(matrix[i], line[start_j:end_j+1])
			start_j = end_j + 2
			end_j = findEndOfColumn(operators_line, start_j+1)
		}
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

func padLinesRight(lines []string) []string {
	maxLineLength := 0
	for _, line := range lines {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}
	paddedLines := make([]string, 0)
	// Pad lines to the right to the maximum length
	for _, line := range lines {
		paddedLine := line
		if len(line) < maxLineLength {
			paddedLine += strings.Repeat(" ", maxLineLength-len(line))
		}
		paddedLines = append(paddedLines, paddedLine)
	}
	return paddedLines
}

func findEndOfColumn(line string, start int) int {
	for i := start; i < len(line); i++ {
		if line[i] == '+' || line[i] == '*' {
			return i - 2
		}
	}
	return len(line) - 1
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

func sumStringifiedInts2nd(operands []string) int { // noqa: typos
	sum := 0
	maximumLength := 0
	for _, o := range operands {
		if len(o) > maximumLength {
			maximumLength = len(o)
		}
	}

	// fmt.Print("Operands (+): ", operands, " ")

	for i := range maximumLength {
		new_operand := ""
		for _, o := range operands {
			// Skip operands with not enough digits for the current iteration
			if len(o) <= i || string(o[i]) == " " {
				continue
			}
			new_operand += string(o[i])
		}
		v, err := strconv.Atoi(new_operand)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Print(v, " ")
		sum += v
	}
	// fmt.Print(": ", sum)
	// fmt.Println()
	return sum
}

func multiplyStringifiedInts2nd(operands []string) int { // noqa: typos
	product := 1
	maximumLength := 0
	for _, o := range operands {
		if len(o) > maximumLength {
			maximumLength = len(o)
		}
	}

	// fmt.Print("Operands (*): ", operands, " ")

	for i := range maximumLength {
		new_operand := ""
		for _, o := range operands {
			// Skip operands with not enough digits for the current iteration
			if len(o) <= i || string(o[i]) == " " {
				continue
			}
			new_operand += string(o[i])
		}
		v, err := strconv.Atoi(new_operand)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Print(v, " ")
		product *= v
	}
	// fmt.Print(": ", product)
	// fmt.Println()
	return product
}
