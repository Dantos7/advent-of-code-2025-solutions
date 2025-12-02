package main

import (
	"strconv"
	"strings"

	"log"

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
	if part2 {
		return runPart2(input)
	}
	return runPart1(input)
}

func runPart1(input string) any {
	sum := 0
	intervals := strings.SplitSeq(input, ",")
	for interval := range intervals {
		sum += getRepeatedNumbersSumPart1(interval)
	}
	return sum
}

func getRepeatedNumbersSumPart1(interval string) int {
	// Naive version - enumerate every number in interval and check if it is repeated
	start_end := strings.Split(interval, "-")
	start, err := strconv.Atoi(start_end[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(start_end[1])
	if err != nil {
		log.Fatal(err)
	}
	sum := 0
	for i := start; i <= end; i++ {
		if isRepeatedPart1(strconv.Itoa(i)) {
			sum += i
		}
	}
	return sum
}

func isRepeatedPart1(input string) bool {
	// Check is done by splitting the string in half and checking if the 2 pieces are equal
	if len(input)%2 == 0 {
		return input[:len(input)/2] == input[len(input)/2:]
	}
	return false
}

func runPart2(input string) any {
	sum := 0
	intervals := strings.SplitSeq(input, ",")
	for interval := range intervals {
		sum_interval := getRepeatedNumbersSumPart2(interval)
		sum += sum_interval
	}
	return sum
}

func getRepeatedNumbersSumPart2(interval string) int {
	// Naive version - enumerate every number in interval and check if it is repeated
	start_end := strings.Split(interval, "-")
	start, err := strconv.Atoi(start_end[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(start_end[1])
	if err != nil {
		log.Fatal(err)
	}
	sum := 0
	for i := start; i <= end; i++ {
		if isRepeatedPart2(strconv.Itoa(i)) {
			sum += i
		}
	}
	return sum
}

func isRepeatedPart2(input string) bool {
	// Check is done by increasing progressively the span size. Then for each span length,
	// - check if the input is a multiple in length of the span
	// - construct a string repeating the first n characters, m times to obtain a string of the same length of the input
	// - check if the 2 strings match
	for span_size := 1; span_size <= len(input)/2; span_size++ {
		if len(input)%span_size == 0 {
			repeat_count := len(input) / span_size
			repeated := strings.Repeat(input[:span_size], repeat_count)
			if repeated == input {
				return true
			}
		}
	}
	return false
}
