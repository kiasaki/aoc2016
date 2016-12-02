package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

const inputFilename = "input"

const padSize = 3

var pad = [][]int{
	{1, 4, 7},
	{2, 5, 8},
	{3, 6, 9},
}

func move(x, y int, direction byte) (int, int) {
	switch direction {
	case 'U':
		y--
	case 'D':
		y++
	case 'L':
		x--
	case 'R':
		x++
	default:
		panic("move: invalid direction given")
	}

	// Clamp x,y to pad size
	x = int(math.Min(math.Max(float64(x), 0.0), padSize-1.0))
	y = int(math.Min(math.Max(float64(y), 0.0), padSize-1.0))

	return x, y
}

func main() {
	var pin string
	var contents []byte
	var err error

	if contents, err = ioutil.ReadFile(inputFilename); err != nil {
		log.Fatal(err)
	}

	// Start at the "5" button position
	x := 1
	y := 1

	// Get an array of movement sequences (for each pin digit)
	sequences := strings.Split(strings.Trim(string(contents), "\n"), "\n")

	for _, sequence := range sequences {
		// Move to the right place
		for i := 0; i < len(sequence); i++ {
			x, y = move(x, y, sequence[i])
		}

		// Record pin number
		pin += strconv.Itoa(pad[x][y])
	}

	fmt.Printf("pin is: %v\n", pin)
}
