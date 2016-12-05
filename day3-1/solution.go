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

func NewTriangle(definition string) Triangle {
	parts := strings.Split(definition, " ")
	partValues := []int{}

	for _, part := range parts {
		if part != "" {
			v, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal(err)
			}
			partValues = append(partValues, v)
		}
	}

	return Triangle{partValues[0], partValues[1], partValues[2]}
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

	// Each line of the input is possibly a triangle
	var validTriangleCount int
	triangles := strings.Split(strings.Trim(string(contents), "\n"), "\n")

	for _, triangleDefinition := range triangles {
		// Parse triangle definition
		triangle := NewTriangle(triangleDefinition)

		// Increment valid count if valid
		if triangle.IsValid() {
			validTriangleCount++
		}
	}

	fmt.Printf("Valid triangles: %v\n", validTriangleCount)
}
