package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

const (
	DNorth int = iota
	DEast
	DSouth
	DWest
)

const DCount = 4

const inputFilename = "input"

func steer(direction byte, headed int) int {
	if direction == 'L' {
		if headed == DNorth {
			return DWest
		} else {
			return headed - 1
		}
	} else {
		return (headed + 1) % DCount
	}
}

func move(headed int, x int, y int) (int, int) {
	switch headed {
	case DNorth:
		return x, y + 1
	case DEast:
		return x + 1, y
	case DSouth:
		return x, y - 1
	case DWest:
		return x - 1, y
	default:
		panic("move: invalid headed value")
	}
}

func containsPoint(points []Point, point Point) bool {
	for _, p := range points {
		if p.X == point.X && p.Y == point.Y {
			return true
		}
	}
	return false
}

func success(x, y int) {
	distance := math.Abs(float64(x + y))
	fmt.Printf("Headquter's (x %v)  (y %v) (distance %.0f)\n", x, y, distance)
}

func main() {
	var contents []byte
	var err error

	if contents, err = ioutil.ReadFile(inputFilename); err != nil {
		log.Fatal(err)
	}

	visited := []Point{}
	headed := DNorth
	x := 0
	y := 0

	// Get an array of instructions (right or left + distance)
	directions := strings.Split(strings.Trim(string(contents), "\n"), ", ")

	for _, direction := range directions {
		// 1. Steer the direction we are headed too (Left or Right)
		headed = steer(direction[0], headed)

		// 2. Move given distance
		distance, err := strconv.Atoi(direction[1:])
		if err != nil {
			log.Fatal(err)
		}

		// Step corner by corner so we get to build up the list of visited places
		for i := 0; i < distance; i++ {
			x, y = move(headed, x, y)

			// Check if we already been here
			if containsPoint(visited, Point{x, y}) {
				success(x, y)
				return
			}

			visited = append(visited, Point{x, y})
		}
	}

	fmt.Println("Can't find intersection we visit twice :(")
}
