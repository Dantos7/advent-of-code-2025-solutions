package main

import (
	"log"
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
	vertices := parseInput(input)
	start := "you"
	end := "out"

	cache := make(map[string]int)
	count := countPathsToEnd(start, end, vertices, cache)
	return count
}

type Edge struct {
	From string
	To   string
}

func parseInput(input string) map[string][]Edge {
	vertices := make(map[string][]Edge)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, " ")
		id := parts[0][:len(parts[0])-1]
		edges := make([]Edge, 0, len(parts)-1)
		for _, to := range parts[1:] {
			edge := Edge{From: id, To: to}
			edges = append(edges, edge)
		}
		vertices[id] = edges
	}
	return vertices
}

func countPathsToEnd(start string, end string, vertices map[string][]Edge, cache map[string]int) int {
	count := 0
	if val, ok := cache[start]; ok {
		return val
	}
	if start == end {
		cache[start] = 1
		return 1
	}
	edges, ok := vertices[start]
	if !ok {
		// Dead end
		return 0
	} else {
		for _, e := range edges {
			count += countPathsToEnd(e.To, end, vertices, cache)
		}
	}
	cache[start] = count
	return count
}

func runPart2(input string) any {
	vertices := parseInput(input)

	cache := make(map[string]int)
	count_svr_fft := countPathsToEnd("svr", "fft", vertices, cache)
	// fmt.Println(count_svr_fft)

	cache = make(map[string]int)
	count_fft_dac := countPathsToEnd("fft", "dac", vertices, cache)
	// fmt.Println(count_fft_dac)

	cache = make(map[string]int)
	count_dac_out := countPathsToEnd("dac", "out", vertices, cache)
	// fmt.Println(count_dac_out)

	count_svr_fft_dac_out := count_svr_fft * count_fft_dac * count_dac_out
	// fmt.Println(count_svr_fft_dac_out)

	cache = make(map[string]int)
	count_svr_dac := countPathsToEnd("svr", "dac", vertices, cache)
	// fmt.Println(count_svr_dac)

	cache = make(map[string]int)
	count_dac_fft := countPathsToEnd("dac", "fft", vertices, cache)
	// fmt.Println(count_dac_fft)

	cache = make(map[string]int)
	count_fft_out := countPathsToEnd("fft", "out", vertices, cache)
	// fmt.Println(count_fft_out)

	count_svr_dac_fft_out := count_svr_dac * count_dac_fft * count_fft_out
	// fmt.Println(count_svr_fft_dac_out)

	return count_svr_fft_dac_out + count_svr_dac_fft_out
}
