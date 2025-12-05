package main

import (
	"fmt"
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
	freshRanges := getFreshRanges(input)
	availableIngredients := strings.Split(strings.Split(input, "\n\n")[1], "\n")

	counterFreshAvailable := 0
	for _, ingredientID := range availableIngredients {
		ingredientIDInt, err := strconv.Atoi(ingredientID)
		if err != nil {
			log.Fatal(err)
		}
		for _, rangeIngredients := range freshRanges {
			if ingredientIDInt >= rangeIngredients[0] && ingredientIDInt <= rangeIngredients[1] {
				counterFreshAvailable++
				break
			}
		}
	}

	return counterFreshAvailable
}

func runPart2(input string) any {
	counter := 0
	freshRanges := getFreshRanges(input)
	var err error
	for {
		newFreshRanges := make([][2]int, 0)
		for i, r := range freshRanges {
			ingRange1 := r
			var ingRange2 [2]int
			merged := false
			for j := i + 1; j < len(freshRanges); j++ {
				ingRange1, ingRange2, err = getUnionOfRanges(ingRange1, freshRanges[j])
				if err != nil {
					log.Fatal(err)
				} else if ingRange2 == [2]int{0, 0} {
					merged = true
					freshRanges[j] = ingRange1
					break
				}
			}
			if !merged {
				newFreshRanges = append(newFreshRanges, ingRange1)
			}
		}
		if sameValues(freshRanges, newFreshRanges) {
			break
		}
		freshRanges = newFreshRanges
	}

	// Sorting is for visualization purposes only
	// sort.Slice(freshRanges, func(i, j int) bool {
	// 	return freshRanges[i][0] < freshRanges[j][0]
	// })
	for _, tr := range freshRanges {
		// fmt.Println("Final Intervals: ", tr, tr[1]-tr[0]+1)
		counter += tr[1] - tr[0] + 1
	}

	return counter
}

func getFreshRanges(input string) [][2]int {
	freshRangesStr := strings.Split(strings.Split(input, "\n\n")[0], "\n")
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
	return freshRangesInt
}

func getUnionOfRanges(ingRange1 [2]int, ingRange2 [2]int) ([2]int, [2]int, error) {
	// Returns the union of range 1 and range 2
	start1, end1 := ingRange1[0], ingRange1[1]
	start2, end2 := ingRange2[0], ingRange2[1]

	switch {
	// Case 0: same ranges - empty difference
	case ingRange1 == ingRange2:
		return ingRange1, [2]int{0, 0}, nil

	// Case 1: start1 < end1 < start2 < end2
	case end1 < start2:
		// Assuming that the intervals are consistent, we can avoid checking start1 < end1 and start2 < end2
		return ingRange1, ingRange2, nil

	// Case 2: start2 < end2 < start1 < end1
	case end2 < start1:
		// Assuming that the intervals are consistent, we can avoid checking start1 < end1 and start2 < end2
		return ingRange1, ingRange2, nil

	// Case 3: start1 <= start2 <= end1 < end2
	case start1 <= start2 && start2 <= end1 && end1 < end2:
		return [2]int{start1, end2}, [2]int{0, 0}, nil

	// Case 4: start2 < start1 < end2 <= end1
	case start2 <= start1 && start1 <= end2 && end2 <= end1:
		return [2]int{start2, end1}, [2]int{0, 0}, nil

	// Case 5: start1 <= start2 < end2 <= end1
	case start1 <= start2 && end2 <= end1:
		return ingRange1, [2]int{0, 0}, nil

	// Case 6: start2 <= start1 <= end1 <= end2
	case start2 <= start1 && end1 <= end2:
		// start1 < end1 can be omitted assuming consistency of intervals
		return ingRange2, [2]int{0, 0}, nil

	// Uncaught case: return error
	default:
		return [2]int{0, 0}, [2]int{0, 0}, fmt.Errorf("Uncaught case: range1=%v, range2=%v", ingRange1, ingRange2)
	}
}

func sameValues(a, b [][2]int) bool {
	if len(a) != len(b) {
		return false
	}

	freq := make(map[[2]int]int)
	for _, v := range a {
		freq[v]++
	}
	for _, v := range b {
		freq[v]--
		if freq[v] < 0 {
			return false
		}
	}

	for _, count := range freq {
		if count != 0 {
			return false
		}
	}
	return true
}
