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
	banks := strings.SplitSeq(input, "\n")
	sumJoltage := 0
	for bank := range banks {
		outputJoltage := getBankOutputJoltage(bank)
		sumJoltage += outputJoltage
	}
	return sumJoltage
}

func getBankOutputJoltage(bank string) int {
	// Retrieves the maximum 2 digits number obtained in-order from the string

	// Use a few maps to optimize string scans
	// - Counter of each digit
	// - Positions of each digit

	// Then proceed by looking at combinations from highest to lowest
	// - If both digits are present in the bank, check positions to ensure order (not needed if the 2 digits are equal and the counter is >=2)
	// - If found, return the number
	// - Ignore combinations where one of the digits is not present in the bank

	counter := make(map[rune]int)
	positions := make(map[rune][]int)
	for i, char := range bank {
		counter[char]++
		positions[char] = append(positions[char], i)
	}

	for _, char0 := range "987654321" {
		for _, char1 := range "987654321" {
			if char0 == char1 && counter[char0] >= 2 {
				// If the 2 digits are equal, there is no need to check for positions (order is not relevant)
				// Hence, we can return
				num, err := strconv.Atoi(string(char0) + string(char1))
				if err != nil {
					log.Fatal(err)
				}
				return num
			} else if char0 != char1 && counter[char0] >= 1 && counter[char1] >= 1 {
				// If the counter is >= 1 for both the digits, proceed in checking the positions
				// For each position of the first digit, check if there is a position of the second digit that is higher (to ensure order)
				// If found, return the number
				for _, pos0 := range positions[char0] {
					for _, pos1 := range positions[char1] {
						if pos0 < pos1 {
							num, err := strconv.Atoi(string(char0) + string(char1))
							if err != nil {
								log.Fatal(err)
							}
							return num
						}
					}
				}
			} else {
				// If digits are different and one of the counter is zero, do nothing (combination certainly does not exist)
				continue
			}
		}
	}
	// Return dummy value if no number found (should not happen with valid input)
	return 11
}

func runPart2(input string) any {
	return "not implemented"
}
