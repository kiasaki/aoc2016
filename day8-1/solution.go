package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const inputFilename = "input"
const screenWidth = 50
const screenHeight = 6

func MustParseInt(text string) int {
	val, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return val
}

func CopyScreen(s [][]bool) [][]bool {
	screen := make([][]bool, screenWidth)
	for x := 0; x < screenWidth; x++ {
		screen[x] = make([]bool, screenHeight)
		copy(screen[x], s[x])
	}
	return screen
}

func main() {
	var contents []byte
	var err error

	if contents, err = ioutil.ReadFile(inputFilename); err != nil {
		log.Fatal(err)
	}

	screen := make([][]bool, screenWidth)
	for i := 0; i < screenWidth; i++ {
		screen[i] = make([]bool, screenHeight)
	}

	instructions := strings.Split(strings.Trim(string(contents), "\n"), "\n")

	rectRegExp := regexp.MustCompile(`^rect (\d+)x(\d+)$`)
	rotateRowRegExp := regexp.MustCompile(`^rotate row y=(\d+) by (\d+)$`)
	rotateColRegExp := regexp.MustCompile(`^rotate column x=(\d+) by (\d+)$`)

	for _, ins := range instructions {
		if ms := rectRegExp.FindAllStringSubmatch(ins, -1); len(ms) > 0 {
			// Turn on a square of pixels from 0,0
			width := MustParseInt(ms[0][1])
			height := MustParseInt(ms[0][2])
			for x := 0; x < width; x++ {
				for y := 0; y < height; y++ {
					screen[x][y] = true
				}
			}
		} else if ms := rotateRowRegExp.FindAllStringSubmatch(ins, -1); len(ms) > 0 {
			// Push row
			y := MustParseInt(ms[0][1])
			by := MustParseInt(ms[0][2])
			newScreen := CopyScreen(screen)
			for x := 0; x < screenWidth; x++ {
				newScreen[(x+by)%screenWidth][y] = screen[x][y]
			}
			screen = newScreen
		} else if ms := rotateColRegExp.FindAllStringSubmatch(ins, -1); len(ms) > 0 {
			// Push column
			x := MustParseInt(ms[0][1])
			by := MustParseInt(ms[0][2])
			newScreen := CopyScreen(screen)
			for y := 0; y < screenHeight; y++ {
				newScreen[x][(y+by)%screenHeight] = screen[x][y]
			}
			screen = newScreen
		}
	}

	pixelsLit := 0
	for x := 0; x < screenWidth; x++ {
		for y := 0; y < screenHeight; y++ {
			if screen[x][y] {
				pixelsLit++
			}
		}
	}

	fmt.Printf("Pixels lit: %v\n", pixelsLit)
}
