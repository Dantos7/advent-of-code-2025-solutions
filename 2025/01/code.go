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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return runPart2(input)
	}
	// solve part 1 here
	return runPart1(input)
}

func runPart1(input string) any {
	counter := 0
	var current uint64 = 50
	var direction rune = 'R'
	var span uint64
	var err error
	log.Println("Start:", current)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		direction = rune(line[0])
		span, err = strconv.ParseUint(line[1:], 10, 64)
		if err != nil {
			log.Fatal("Failed to parse span: ", err)
		}
		current = rotate(current, direction, span)
		log.Println(string(direction), span, "=>", current)
		if current == 0 {
			counter++
		}
	}
	return counter
}

func rotate(current uint64, direction rune, span uint64) uint64 {
	// Use modulus 100 to wrap around the circle (use it also for the span to ignore multiple full rotations)
	if direction == 'R' {
		return (current + span%100) % 100
	}
	return (current - span%100 + 100) % 100
}

func runPart2(input string) any {
	return "not implemented"
}
