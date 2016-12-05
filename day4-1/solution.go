package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

const inputFilename = "input"

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
	if a[i].Count == a[j].Count {
		// Sort by letter (alphabetically) when same letter count
		return a[i].Letter < a[j].Letter
	} else {
		// Normally, sort by count (descending)
		return a[i].Count > a[j].Count
	}
}

type Room struct {
	NameLetters LetterCounts
	Number      int
	Checksum    string
}

func NewRoom(definition string) Room {
	// Get all room parts (two last items being number and checksum)
	parts := strings.Split(
		strings.Replace(definition[:len(definition)-1], "[", "-", -1), "-",
	)

	// Parse room number as int
	number, err := strconv.Atoi(parts[len(parts)-2])
	if err != nil {
		log.Fatal(err)
	}

	// Create a room object just missing room letter counts
	room := Room{
		NameLetters: NewLetterCounts(),
		Number:      number,
		Checksum:    parts[len(parts)-1],
	}

	for _, part := range parts[:len(parts)-2] {
		for _, letter := range part {
			room.NameLetters[letter-'a'].Count++
		}
	}

	return room
}

func (r Room) IsValid() bool {
	letterCounts := r.NameLetters
	sort.Sort(letterCounts)

	checksum := ""
	checksum += string(letterCounts[0].Letter)
	checksum += string(letterCounts[1].Letter)
	checksum += string(letterCounts[2].Letter)
	checksum += string(letterCounts[3].Letter)
	checksum += string(letterCounts[4].Letter)

	fmt.Println(checksum, r.Checksum, r.NameLetters)

	return checksum == r.Checksum
}

func main() {
	var contents []byte
	var err error

	if contents, err = ioutil.ReadFile(inputFilename); err != nil {
		log.Fatal(err)
	}

	// Parse rooms
	rows := strings.Split(strings.Trim(string(contents), "\n"), "\n")
	rooms := []Room{}
	for _, row := range rows {
		rooms = append(rooms, NewRoom(row))
	}

	// Filter and sum valid room numbers
	var validRoomNumberSum int
	for _, room := range rooms {
		if room.IsValid() {
			validRoomNumberSum += room.Number
		}
	}

	fmt.Printf("Valid room number sum: %v\n", validRoomNumberSum)
}
