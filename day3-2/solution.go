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

type ByValue []int

type Triangle struct {
	A, B, C int
}

func (t Triangle) IsValid() bool {
	sum := t.A + t.B + t.C
	max := int(math.Max(math.Max(float64(t.A), float64(t.B)), float64(t.C)))
	return sum-max > max
}

func main() {
	var contents []byte
	var err error

	if contents, err = ioutil.ReadFile(inputFilename); err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(strings.Trim(string(contents), "\n"), "\n")
	values := make([]int, len(rows)*3)

	// Create one array of values based on columns not rows
	for i, row := range rows {
		parts := strings.Split(row, " ")

		j := 0
		for _, part := range parts {
			if part != "" {
				v, err := strconv.Atoi(part)
				if err != nil {
					log.Fatal(err)
				}
				values[(j*len(rows))+i] = v
				j++
			}
		}
	}

	triangles := []Triangle{}
	for i := 0; i < len(values); i += 3 {
		triangles = append(triangles, Triangle{values[i], values[i+1], values[i+2]})
	}

	var validTriangleCount int
	for _, triangle := range triangles {
		if triangle.IsValid() {
			validTriangleCount++
		}
	}

	fmt.Printf("Valid triangles: %v\n", validTriangleCount)
}
