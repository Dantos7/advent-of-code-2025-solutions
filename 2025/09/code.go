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
	coords := parseInput(input)
	// var rows, cols int
	// if len(coords) < 10 {
	// 	// Example
	// 	rows = 9
	// 	cols = 15
	// } else {
	// 	rows = 100000
	// 	cols = 100000
	// }
	var maxCoords [2][2]int

	horizontalEdges := getHorizontalEdges(coords)
	verticalEdges := getVerticalEdges(coords)

	max_area := -1
	// var maxCoords [2][2]int
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			area := computeArea(c1, c2)
			if area > max_area {
				c3 := [2]int{c1[0], c2[1]}
				c4 := [2]int{c2[0], c1[1]}
				if isInside(c3, horizontalEdges, verticalEdges) && isInside(c4, horizontalEdges, verticalEdges) {
					max_area = area
					maxCoords = [2][2]int{c1, c2}
				}
			}
		}
	}
	fmt.Println(maxCoords)
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

// TODO: !BUGGED -> always false
func isInside(coord [2]int, horizontalEdges []HorizontalEdge, verticalEdges []VerticalEdge) bool {
	// A coord is inside if an edge can be reached from each direction (up,down,right,left)

	// Up
	up := hasEdgeUp(coord, horizontalEdges, verticalEdges)
	// fmt.Println("UP", coord, up)
	if !up {
		return false
	}

	down := hasEdgeDown(coord, horizontalEdges, verticalEdges)
	// fmt.Println("DOWN", coord, down)
	if !down {
		return false
	}

	left := hasEdgeLeft(coord, horizontalEdges, verticalEdges)
	// fmt.Println("LEFT", coord, left)
	if !left {
		return false
	}

	right := hasEdgeRight(coord, horizontalEdges, verticalEdges)
	// fmt.Println("RIGHT", coord, right)

	return right
}

func hasEdgeUp(coord [2]int, horizontalEdges []HorizontalEdge, verticalEdges []VerticalEdge) bool {
	// Checks if the coordinate has an edge above
	verticalUpFound := false
	horizontalUpFound := false
	for _, e := range verticalEdges {
		if coord[0] == e.x && coord[1] >= e.segment[0] {
			verticalUpFound = true
			break
		} else if coord[0] < e.x {
			// Since the array is sorted, we can exit here
			break
		}
	}
	if !verticalUpFound {
		for _, e := range horizontalEdges {
			// fmt.Println("HORIZONTAL", e, e.y)
			if coord[1] >= e.y && e.segment[0] <= coord[0] && coord[0] <= e.segment[1] {
				horizontalUpFound = true
				break
			} else if coord[1] < e.y {
				// Since the array is sorted, we can exit here
				break
			}
		}
	}
	return verticalUpFound || horizontalUpFound
}

func hasEdgeDown(coord [2]int, horizontalEdges []HorizontalEdge, verticalEdges []VerticalEdge) bool {
	// Checks if the coordinate has an edge below
	verticalDownFound := false
	horizontalDownFound := false
	for _, e := range verticalEdges {
		if coord[0] == e.x && coord[1] <= e.segment[1] {
			verticalDownFound = true
			break
		} else if coord[0] < e.x {
			// Since the array is sorted, we can exit here
			break
		}
	}
	if !verticalDownFound {
		for i := len(horizontalEdges) - 1; i >= 0; i-- {
			e := horizontalEdges[i]
			// fmt.Println("Horizontal", e, coord[1] <= e.y, e.segment[0] <= coord[0], coord[0] <= e.segment[1])
			if coord[1] <= e.y && e.segment[0] <= coord[0] && coord[0] <= e.segment[1] {
				horizontalDownFound = true
				break
			} else if coord[1] > e.y {
				// Since the array is sorted, we can exit here
				break
			}
		}
	}
	return verticalDownFound || horizontalDownFound
}

func hasEdgeLeft(coord [2]int, horizontalEdges []HorizontalEdge, verticalEdges []VerticalEdge) bool {
	// Checks if the coordinate has an edge left
	verticalLeftFound := false
	horizontalLeftFound := false
	for _, e := range verticalEdges {
		if coord[0] >= e.x && e.segment[0] <= coord[1] && coord[1] <= e.segment[1] {
			verticalLeftFound = true
			break
		} else if coord[0] < e.x {
			// Since the array is sorted, we can exit here
			break
		}
	}
	if !verticalLeftFound {
		for _, e := range horizontalEdges {
			// fmt.Println("HORIZONTAL", e, coord[1] == e.y, coord[0] >= e.segment[0])
			if coord[1] == e.y && coord[0] >= e.segment[0] {
				horizontalLeftFound = true
				break
			} else if coord[1] < e.y {
				// Since the array is sorted, we can exit here
				break
			}
		}
	}
	return verticalLeftFound || horizontalLeftFound
}

func hasEdgeRight(coord [2]int, horizontalEdges []HorizontalEdge, verticalEdges []VerticalEdge) bool {
	// Checks if the coordinate has an edge right
	verticalRightFound := false
	horizontalRightFound := false
	for i := len(verticalEdges) - 1; i > 0; i-- {
		e := verticalEdges[i]
		// fmt.Println("Vertical", e, coord[0] <= e.x, e.segment[0] <= coord[0], coord[0] <= e.segment[1])
		if coord[0] <= e.x && e.segment[0] <= coord[1] && coord[1] <= e.segment[1] {
			verticalRightFound = true
			break
		} else if coord[0] > e.x {
			// Since the array is sorted, we can exit here
			break
		}
	}
	if !verticalRightFound {
		for _, e := range horizontalEdges {
			if coord[1] == e.y && coord[1] <= e.segment[1] {
				horizontalRightFound = true
				break
			} else if coord[1] < e.y {
				// Since the array is sorted, we can exit here
				break
			}
		}
	}
	return verticalRightFound || horizontalRightFound
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
