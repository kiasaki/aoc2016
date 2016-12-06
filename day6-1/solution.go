package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type LetterCount struct {
	Letter rune
	Count  int
}

type LetterCounts []LetterCount

func NewLetterCounts() LetterCounts {
	lc := make(LetterCounts, 26)
	for l := 'a'; l <= 'z'; l++ {
		lc[l-'a'] = LetterCount{
			Letter: l,
			Count:  0,
		}
	}
	return lc
}

func (a LetterCounts) Len() int {
	return len(a)
}
func (a LetterCounts) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a LetterCounts) Less(i, j int) bool {
	// Sort by count (desc)
	return a[i].Count > a[j].Count
}

const inputFilename = "input"

func main() {
	var contents []byte
	var err error

	if contents, err = ioutil.ReadFile(inputFilename); err != nil {
		log.Fatal(err)
	}

	// Read all input in a 2D array for easier traversal
	rows := strings.Split(strings.Trim(string(contents), "\n"), "\n")
	messageCharacterTable := make([][]rune, len(rows))
	for i, row := range rows {
		messageCharacterTable[i] = make([]rune, len(row))
		for j, character := range row {
			messageCharacterTable[i][j] = character
		}
	}

	message := ""

	// Compute most frequent letter column by column
	for x := 0; x < len(messageCharacterTable[0]); x++ {
		letterCounts := NewLetterCounts()

		// Increment the count each time a letter is encountered
		for y := 0; y < len(messageCharacterTable); y++ {
			c := messageCharacterTable[y][x]
			letterCounts[c-'a'].Count++
		}

		// Sort letter counts and pick the most occuring
		sort.Sort(letterCounts)
		fmt.Println(letterCounts)
		message += string(letterCounts[0].Letter)
	}

	fmt.Printf("Error corrected message: %v\n", message)
}
