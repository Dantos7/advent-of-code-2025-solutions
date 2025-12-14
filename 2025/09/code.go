package main

import (
	"fmt"
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
	coords := parseInput(input)
	max_area := -1
	// var maxCoords [2][2]int
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			area := computeArea(c1, c2)
			if area > max_area {
				max_area = area
				// maxCoords = [2][2]int{c1, c2}
			}
		}
	}
	// fmt.Println(maxCoords)
	return max_area
}

func computeArea(c1 [2]int, c2 [2]int) int {
	return (absInt(c1[0]-c2[0]) + 1) * (absInt(c1[1]-c2[1]) + 1)
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
	coords := parseInput(input)

	horizontalEdges := getHorizontalEdges(coords)
	verticalEdges := getVerticalEdges(coords)

	max_area := -1
	// var maxCoords [2][2]int
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			area := computeArea(c1, c2)
			if area > max_area {
				if isRectangleValid(c1, c2, horizontalEdges, verticalEdges, coords) {
					max_area = area
					// maxCoords = [2][2]int{c1, c2}
					// fmt.Println("NEW MAX AREA", max_area, c1, c2)
				}
			}
		}
	}
	// fmt.Println(maxCoords)

	return max_area
}

type HorizontalEdge struct {
	segment [2]int
	y       int
}

type VerticalEdge struct {
	x       int
	segment [2]int
}

func getHorizontalEdges(coords [][2]int) []HorizontalEdge {
	horizontalEdges := make([]HorizontalEdge, 0, len(coords)/2)

	for i, c1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			c2 := coords[j]
			if j == i+1 {
				if c1[1] == c2[1] {
					var segment [2]int
					if c1[0] < c2[0] {
						segment = [2]int{c1[0], c2[0]}
					} else {
						segment = [2]int{c2[0], c1[0]}
					}
					horizontalEdges = append(horizontalEdges, HorizontalEdge{segment, c1[1]})
				}
			}
		}
	}

	// Add the last edge
	c1 := coords[0]
	c2 := coords[len(coords)-1]
	if c1[1] == c2[1] {
		var segment [2]int
		if c1[0] < c2[0] {
			segment = [2]int{c1[0], c2[0]}
		} else {
			segment = [2]int{c2[0], c1[0]}
		}
		horizontalEdges = append(horizontalEdges, HorizontalEdge{segment, c1[1]})
	}

	slices.SortFunc(horizontalEdges, func(a, b HorizontalEdge) int {
		if a.y < b.y {
			return -1
		} else if a.y > b.y {
			return 1
		} else {
			return 0
		}
	})

	return horizontalEdges
}

func getVerticalEdges(coords [][2]int) []VerticalEdge {
	verticalEdges := make([]VerticalEdge, 0, len(coords)/2)

	for i, c1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			c2 := coords[j]
			if j == i+1 {
				if c1[0] == c2[0] {
					var segment [2]int
					if c1[1] < c2[1] {
						segment = [2]int{c1[1], c2[1]}
					} else {
						segment = [2]int{c2[1], c1[1]}
					}
					verticalEdges = append(verticalEdges, VerticalEdge{c1[0], segment})
				}
			}
		}
	}

	// Add the last edge
	c1 := coords[0]
	c2 := coords[len(coords)-1]
	if c1[0] == c2[0] {
		var segment [2]int
		if c1[1] < c2[1] {
			segment = [2]int{c1[1], c2[1]}
		} else {
			segment = [2]int{c2[1], c1[1]}
		}
		verticalEdges = append(verticalEdges, VerticalEdge{c1[0], segment})
	}

	slices.SortFunc(verticalEdges, func(a, b VerticalEdge) int {
		if a.x < b.x {
			return -1
		} else if a.x > b.x {
			return 1
		} else {
			return 0
		}
	})

	return verticalEdges
}

func isInside(coord [2]int, horizontalEdges []HorizontalEdge, verticalEdges []VerticalEdge) bool {
	// Odd-even method: A coord is inside if
	// - it belongs to an edge
	// - from left to it, an odd number of edges is crossed
	// 		Horizontal edges can be ignored (as the vertices belongs to 2 vertical edges anyway)

	// Check belonging
	for _, e := range horizontalEdges {
		if e.y == coord[1] && e.segment[0] <= coord[0] && coord[0] <= e.segment[1] {
			return true
		} else if e.y > coord[1] {
			// Early exit since edges are ordered
			break
		}
	}
	for _, e := range verticalEdges {
		if e.x == coord[0] && e.segment[0] <= coord[1] && coord[1] <= e.segment[1] {
			// It belongs to this vertical edge
			return true
		} else if e.x > coord[0] {
			// Early exit since edges are ordered
			break
		}
	}

	// Count intersections
	count := 0
	ignoredEdges := make(map[int]struct{})
	for i := 0; i < len(verticalEdges); i++ {
		e := verticalEdges[i]
		if _, ok := ignoredEdges[i]; ok {
			continue
		}
		if e.x < coord[0] && e.segment[0] < coord[1] && coord[1] < e.segment[1] {
			// It belongs to this vertical edge
			count += 1
		} else if e.x < coord[0] && e.segment[0] == coord[1] {
			// Find the next vertical edge. If they go in the same direction count 2 and skip both, otherwise count 1 and skip both
			for j := 0; j < len(verticalEdges); j++ {
				if _, ok := ignoredEdges[j]; ok {
					continue
				}
				if j != i && verticalEdges[j].segment[0] == coord[1] {
					// Found, don't count edge, because same direction
					// Remove the edge to not count it again
					ignoredEdges[j] = struct{}{}
					break
				} else if j != i && verticalEdges[j].segment[1] == coord[1] {
					// Found, count edge because different direction
					// Remove the edge to not count it again
					ignoredEdges[j] = struct{}{}
					count += 1
					break
				}
			}
		} else if e.x < coord[0] && e.segment[1] == coord[1] {
			// Find the next vertical edge. If they go in the same direction count 2 and skip both, otherwise count 1 and skip both
			for j := 0; j < len(verticalEdges); j++ {
				if _, ok := ignoredEdges[j]; ok {
					continue
				}
				if j != i && verticalEdges[j].segment[1] == coord[1] {
					// Found, don't count edge, because same direction
					// Remove the edge to not count it again
					ignoredEdges[j] = struct{}{}
					break
				} else if j != i && verticalEdges[j].segment[0] == coord[1] {
					// Found, count edge because different direction
					// Remove the edge to not count it again
					ignoredEdges[j] = struct{}{}
					count += 1
					break
				}
			}
		} else if e.x > coord[0] {
			// Early exit since edges are ordered
			break
		}
	}
	return count%2 == 1
}

func visualizeMatrix(redCoordsSet map[[2]int]bool, greenCoordsSet map[[2]int]bool, rows int, cols int) {
	// Coordinates are cartesian (x,y) -> opposite as standard array indexing

	fmt.Println()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if redCoordsSet[[2]int{j, i}] {
				fmt.Print("#")
			} else if greenCoordsSet[[2]int{j, i}] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Go doesn't have an abs native function for integers 🙃
func absInt(x int) int {
	if x >= 0 {
		return x
	} else {
		return 0 - x
	}
}

func isRectangleValid(c1, c2 [2]int, horizontalEdges []HorizontalEdge, verticalEdges []VerticalEdge, coords [][2]int) bool {
	// Rectangle is valid if
	//    - All vertices are inside
	//    - No red edge crosses any edge (in opposite directions - horizontal/vertical and vertical/horizontal)

	c3 := [2]int{c1[0], c2[1]}
	c4 := [2]int{c2[0], c1[1]}

	if !isInside(c3, horizontalEdges, verticalEdges) || !isInside(c4, horizontalEdges, verticalEdges) {
		return false
	}

	min_y := min(c1[1], c2[1])
	max_y := max(c1[1], c2[1])
	min_x := min(c1[0], c2[0])
	max_x := max(c1[0], c2[0])

	// Check that no point is inside the rectangle (if it is inside, then an edge has crossed)
	for _, c := range coords {
		if min_x < c[0] && c[0] < max_x && min_y < c[1] && c[1] < max_y {
			return false
		}
	}

	// Check for full crossing of horizontal edges (no edge vertex inside)
	for _, e := range horizontalEdges {
		if min_y < e.y && e.y < max_y {
			if e.segment[0] <= min_x && max_x <= e.segment[1] {
				return false
			}
		}
	}

	// Check for full crossing of vertical edges (no edge vertex inside)
	for _, e := range verticalEdges {
		if min_x < e.x && e.x < max_x {
			if e.segment[0] <= min_y && max_y <= e.segment[1] {
				return false
			}
		}
	}

	return true
}
