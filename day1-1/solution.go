package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

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

func move(headed int, distance int, x int, y int) (int, int) {
	switch headed {
	case DNorth:
		return x, y + distance
	case DEast:
		return x + distance, y
	case DSouth:
		return x, y - distance
	case DWest:
		return x - distance, y
	default:
		panic("move: invalid headed value")
	}
}

func main() {
	var contents []byte
	var err error

	if contents, err = ioutil.ReadFile(inputFilename); err != nil {
		log.Fatal(err)
	}

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

		x, y = move(headed, distance, x, y)
	}

	distance := math.Abs(float64(x + y))
	fmt.Printf("Headquter's (x %v)  (y %v) (distance %.0f)\n", x, y, distance)
}
