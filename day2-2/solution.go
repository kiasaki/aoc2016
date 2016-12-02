package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const inputFilename = "input"

const padSize = 5
const nilDigit = "_"

// Surroung the "diamond" pad with nil digits for easy overflow/clamp checks
var pad = [][]string{
	{"_", "_", "_", "_", "_", "_", "_"},
	{"_", "_", "_", "1", "_", "_", "_"},
	{"_", "_", "2", "3", "4", "_", "_"},
	{"_", "5", "6", "7", "8", "9", "_"},
	{"_", "_", "A", "B", "C", "_", "_"},
	{"_", "_", "_", "D", "_", "_", "_"},
	{"_", "_", "_", "_", "_", "_", "_"},
}

func move(x, y int, direction byte) (int, int) {
	switch direction {
	case 'U':
		if pad[y-1][x] != nilDigit {
			y--
		}
	case 'D':
		if pad[y+1][x] != nilDigit {
			y++
		}
	case 'L':
		if pad[y][x-1] != nilDigit {
			x--
		}
	case 'R':
		if pad[y][x+1] != nilDigit {
			x++
		}
	default:
		panic("move: invalid direction given")
	}

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
	y := 3

	// Get an array of movement sequences (for each pin digit)
	sequences := strings.Split(strings.Trim(string(contents), "\n"), "\n")

	for _, sequence := range sequences {
		// Move to the right place
		for i := 0; i < len(sequence); i++ {
			x, y = move(x, y, sequence[i])
		}

		// Record pin number
		pin += pad[y][x]
	}

	fmt.Printf("pin is: %v\n", pin)
}
