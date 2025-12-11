package main

import (
	"log"
	"slices"
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
	// Strategy (for each line)
	// Start with 1 press, check if any 1 press obtains the diagram
	// If not, check with 2 presses (build all possible combinations of 2 presses)
	// Continue up to 2*len(buttons) presses (any button should be pressed at most twice)

	diagrams, buttonsList, _ := parseInput(input)

	sum := 0
	for i := 0; i < len(diagrams); i++ {
		diagram := diagrams[i]
		buttons := buttonsList[i]
		minimalPresses := getMinimalPressesLine(diagram, buttons)
		// fmt.Println(minimalPresses)

		sum += minimalPresses
	}

	return sum
}

func getMinimalPressesLine(diagram []bool, buttons [][]int) int {
	presses := 1
	// Create slice with increasing numbers 0,1,2,...,len(buttons)
	buttonsIndexes := make([]int, len(buttons))
	for i := range buttonsIndexes {
		buttonsIndexes[i] = i
	}

	for p := presses; p < 2*len(buttons); p++ {
		// Select a set of buttons to press and check if state matches diagram
		// Compute all combinations of p elements with repetition
		//   (It is possible to optimize discarding combinations where a button is repeated more than twice, but it is not needed)
		combos := combinationsWithRepetition(buttonsIndexes, p)
		for _, c := range combos {
			state := make([]bool, len(diagram))
			for _, i := range c {
				state = applyButton(state, buttons[i])
				if slices.Equal(state, diagram) {
					return p
				}
			}
		}
	}
	// If no combination is found the limit on p is too low
	log.Fatal("Fail")
	return 0
}

func applyButton(state []bool, button []int) []bool {
	for _, i := range button {
		state[i] = !state[i]
	}
	return state
}

// Thanks, Google <3
func combinationsWithRepetition(elements []int, k int) [][]int {
	var results [][]int
	var currentCombination []int

	// Recursive helper function
	var generate func(startIndex int, currentCombination []int)
	generate = func(startIndex int, currentCombination []int) {
		// Base case: if the current combination is complete (length k), add it to the results
		if len(currentCombination) == k {
			// Create a copy to avoid mutation issues
			comboCopy := make([]int, k)
			copy(comboCopy, currentCombination)
			results = append(results, comboCopy)
			return
		}

		// Iterate through elements starting from startIndex to allow repetition
		for i := startIndex; i < len(elements); i++ {
			// Add current element to the combination
			currentCombination = append(currentCombination, elements[i])
			// Recurse, importantly using 'i' (not 'i+1') to allow the same element to be picked again
			generate(i, currentCombination)
			// Backtrack: remove the last element to explore other possibilities
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	generate(0, currentCombination)
	return results
}

func parseInput(input string) ([][]bool, [][][]int, [][]int) {
	lines := strings.Split(input, "\n")
	diagrams := make([][]bool, 0, len(lines))
	buttonsAll := make([][][]int, 0, len(lines))
	joltageRequirementsAll := make([][]int, 0, len(lines))

	for _, line := range lines {
		// Diagrams
		split := strings.Split(line, "] ")
		diagramStr := split[0][1:]
		diagram := make([]bool, 0, len(diagramStr))
		for _, c := range diagramStr {
			switch c {
			case '#':
				diagram = append(diagram, true)
			case '.':
				diagram = append(diagram, false)
			default:
				log.Fatal(string(c))
			}
		}
		diagrams = append(diagrams, diagram)

		// Button wirings
		buttonsAndJoltageStr := split[1]
		buttonsStr := strings.Split(buttonsAndJoltageStr, " {")[0]
		buttonsListStr := strings.Split(buttonsStr, " ")
		buttons := make([][]int, 0, len(buttonsListStr))
		for _, bStr := range buttonsListStr {
			bWiringsStr := strings.Split(bStr[1:len(bStr)-1], ",")
			b := make([]int, 0, len(bWiringsStr))
			for _, bWiringStr := range bWiringsStr {
				bWiringInt, err := strconv.Atoi(bWiringStr)
				if err != nil {
					log.Fatal(err)
				}
				b = append(b, bWiringInt)
			}
			buttons = append(buttons, b)
		}
		buttonsAll = append(buttonsAll, buttons)

		// Joltage requirements
		joltageStr := strings.Split(buttonsAndJoltageStr, " {")[1]
		joltageStr = joltageStr[:len(joltageStr)-1]
		joltageListStr := strings.Split(joltageStr, ",")
		joltageRequirements := make([]int, 0, len(joltageListStr))
		for _, jStr := range joltageListStr {
			j, err := strconv.Atoi(jStr)
			if err != nil {
				log.Fatal(err)
			}
			joltageRequirements = append(joltageRequirements, j)
		}

		joltageRequirementsAll = append(joltageRequirementsAll, joltageRequirements)
	}

	return diagrams, buttonsAll, joltageRequirementsAll
}

func runPart2(input string) any {
	return "not implemented"
}
