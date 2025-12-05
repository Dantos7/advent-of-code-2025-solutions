package main

import (
	// "strconv"
	// "strings"

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
	freshRangesStr := strings.Split(strings.Split(input, "\n\n")[0], "\n")
	availableIngredients := strings.Split(strings.Split(input, "\n\n")[1], "\n")

	freshRangesInt := make([][2]int, len(freshRangesStr))
	for i, rangeStr := range freshRangesStr {
		var start, end int
		var err error
		start_end := strings.Split(rangeStr, "-")
		start, err = strconv.Atoi(start_end[0])
		if err != nil {
			log.Fatal(err)
		}
		freshRangesInt[i][0] = start
		end, err = strconv.Atoi(start_end[1])
		if err != nil {
			log.Fatal(err)
		}
		freshRangesInt[i][1] = end
	}

	counterFreshAvailable := 0
	for _, ingredientID := range availableIngredients {
		ingredientIDInt, err := strconv.Atoi(ingredientID)
		if err != nil {
			log.Fatal(err)
		}
		for _, rangeIngredients := range freshRangesInt {
			if ingredientIDInt >= rangeIngredients[0] && ingredientIDInt <= rangeIngredients[1] {
				counterFreshAvailable++
				break
			}
		}
	}

	return counterFreshAvailable
}

func runPart2(input string) any {
	return "not implemented"
}
