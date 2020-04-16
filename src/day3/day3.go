package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Direction string

const (
	UP    Direction = "U"
	RIGHT Direction = "R"
	DOWN  Direction = "D"
	LEFT  Direction = "L"
)

type Coord struct {
	X int
	Y int
}

func main() {
	wire1, wire2 := getInput()
	wire1Coords := make(map[Coord]bool)
	position := Coord{0, 0}
	var intersections []Coord

	for _, command := range wire1 {
		// populate wire1Coords with all points wire1 covers
		addPath(&position, &wire1Coords, command)
	}

	position = Coord{0, 0}

	for _, command := range wire2 {
		// we just need to look for intersections, so we don't need to populate
		// and preserve an entire collection of wire2's exploits
		// we can just get the affected coords one command at a time and
		// check if any of those overlap with wire1
		coords := make(map[Coord]bool)

		addPath(&position, &coords, command)

		for coord, _ := range coords {
			_, exists := wire1Coords[coord]

			if exists {
				intersections = append(intersections, coord)
			}
		}
	}

	fmt.Println(findClosest(intersections))
}

func findClosest(intersections []Coord) float64 {
	// skip the first intersection which is the origin
	var distances []float64

	for i := 1; i < len(intersections); i++ {
		intersection := intersections[i]
		// -12 and 12 are the same distance away, so get absolute value
		// to avoid adding negative numbers
		x := math.Abs(float64(intersection.X))
		y := math.Abs(float64(intersection.Y))
		distance := x + y
		distances = append(distances, distance)
	}

	sort.Float64s(distances)
	return distances[0]
}

// updates the position to the correct coordinate based on the incoming
// command
// adds all touched points to the coords map
func addPath(position *Coord, coords *map[Coord]bool, command string) {
	// split the command up based on alpha/numerics
	a := regexp.MustCompile("([A-Z]+)([0-9]+)")
	test := a.FindStringSubmatch(command)
	direction := Direction(test[1])
	distance, _ := strconv.Atoi(test[2])

	var i int = 0

	for i < distance {
		point := *position

		switch direction {
		case UP:
			point.Y += 1

		case RIGHT:
			point.X += 1

		case DOWN:
			point.Y -= 1

		case LEFT:
			point.X -= 1
		}

		(*coords)[point] = true
		(*position) = point

		i++
	}
}

func getInput() ([]string, []string) {
	file, err := ioutil.ReadFile("src/day3/input.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(file), "\n")
	wire1 := strings.Split(inputs[0], ",")
	wire2 := strings.Split(inputs[1], ",")

	return wire1, wire2
}

// 280
