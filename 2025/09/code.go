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
	coords := parseInput(input)
	max_area := -1
	// var max_coords [2][2]int
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			area := computeArea(c1, c2)
			if area > max_area {
				max_area = area
				// max_coords = [2][2]int{c1, c2}
			}
		}
	}
	// fmt.Println(max_coords)
	return max_area
}

func computeArea(c1 [2]int, c2 [2]int) int {
	return absInt((c1[0] - c2[0] + 1) * (c1[1] - c2[1] + 1))
}

func parseInput(input string) [][2]int {
	lines := strings.Split(input, "\n")
	coords := make([][2]int, 0, len(lines))
	for _, l := range lines {
		x_y := strings.Split(l, ",")
		x, err_x := strconv.Atoi(x_y[0])
		if err_x != nil {
			log.Fatal(err_x)
		}
		y, err_y := strconv.Atoi(x_y[1])
		if err_y != nil {
			log.Fatal(err_y)
		}
		coords = append(coords, [2]int{x, y})
	}
	return coords
}

func runPart2(input string) any {
	return "not implemented"
}

// Go doesn't have an abs native function for integers 🙃
func absInt(x int) int {
	if x >= 0 {
		return x
	} else {
		return 0 - x
	}
}
