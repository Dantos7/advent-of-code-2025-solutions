package main

import (
	"log"
	"math"
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
	boxes := parseInput(input)
	var merges int
	if len(boxes) <= 20 {
		// Example -> 10 merges
		merges = 10
	} else {
		// User input -> 1000 merges
		merges = 1000
	}
	boxesToEuclidianDistances := make(map[[2][3]int]float64, 0)
	euclidianDistancesToBoxes := make(map[float64][2][3]int, 0)
	boxToCircuit := make(map[[3]int]int)
	circuitToBoxes := make(map[int][][3]int)
	// Assign a circuit to each box
	for i, b := range boxes {
		boxToCircuit[b] = i
		circuitToBoxes[i] = [][3]int{b}
	}

	// First pass compute all euclidian distances
	for i, b1 := range boxes {
		for j, b2 := range boxes {
			if j <= i {
				continue
			}
			distance := euclidianDistance(b1, b2)
			boxesToEuclidianDistances[[2][3]int{b1, b2}] = distance
			// Assert that there are not 2 couple of boxes with the same euclidian distance
			if otherBoxes, ok := euclidianDistancesToBoxes[distance]; ok {
				log.Fatal("Shared euclidian distance for ", b1, ", ", b2, " and ", otherBoxes)
			}
			euclidianDistancesToBoxes[distance] = [2][3]int{b1, b2}
			boxesToEuclidianDistances[[2][3]int{b2, b1}] = distance
		}
	}

	// Extract euclidian distances keys
	euclidianDistances := make([]float64, 0, len(euclidianDistancesToBoxes))
	for d := range euclidianDistancesToBoxes {
		euclidianDistances = append(euclidianDistances, d)
	}
	// Sort euclidian distances
	slices.Sort(euclidianDistances)

	// Merge the boxes in circuits for `merges` iterations by taking the smallest euclidian distance at each iteration
	for m := 0; m < merges; m++ {
		smallestEuclidianDistance := euclidianDistances[0]
		euclidianDistances = euclidianDistances[1:]
		boxesToMerge := euclidianDistancesToBoxes[smallestEuclidianDistance]
		mergeBoxes(boxesToMerge, boxToCircuit, circuitToBoxes)
		// fmt.Println(boxesToMerge, boxToCircuit[boxesToMerge[0]], boxToCircuit[boxesToMerge[1]])
	}

	// fmt.Println(merges, circuitToBoxes)

	// Extract sizes of circuits
	sizes := make([]int, 0, len(circuitToBoxes))
	for _, boxes := range circuitToBoxes {
		sizes = append(sizes, len(boxes))
		// fmt.Println(c, len(boxes))
	}
	slices.Sort(sizes)

	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func parseInput(input string) [][3]int {
	lines := strings.Split(input, "\n")
	boxes := make([][3]int, len(lines))
	var err error
	var x, y, z int
	for i, line := range lines {
		splitLine := strings.Split(line, ",")
		x, err = strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err = strconv.Atoi(splitLine[1])
		if err != nil {
			log.Fatal(err)
		}
		z, err = strconv.Atoi(splitLine[2])
		if err != nil {
			log.Fatal(err)
		}
		boxes[i] = [3]int{x, y, z}
	}
	return boxes
}

func euclidianDistance(p [3]int, q [3]int) float64 {
	dx := p[0] - q[0]
	dy := p[1] - q[1]
	dz := p[2] - q[2]
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func mergeBoxes(boxesToMerge [2][3]int, boxToCircuit map[[3]int]int, circuitToBoxes map[int][][3]int) {
	b1, b2 := boxesToMerge[0], boxesToMerge[1]
	c1 := boxToCircuit[b1]
	c2 := boxToCircuit[b2]
	if c1 == c2 {
		// Nothing to do - circuits are already merged
	} else {
		// Merge circuits
		for _, b := range circuitToBoxes[c2] {
			boxToCircuit[b] = c1
			circuitToBoxes[c1] = append(circuitToBoxes[c1], b)
		}
		delete(circuitToBoxes, c2)
	}
}

func runPart2(input string) any {
	return "not implemented"
}
