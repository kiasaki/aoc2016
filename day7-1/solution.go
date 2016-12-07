package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const inputFilename = "input"

func containsAbba(seq string) bool {
	for i := 0; i <= len(seq)-4; i++ {
		// From current index, for the next 4 letters, is the first the same as
		// the fourth and the second the same as the third
		// Also make sure they are not all the same letters
		if seq[i] == seq[i+3] && seq[i+1] == seq[i+2] && seq[i] != seq[i+1] {
			return true
		}
	}
	return false
}

func main() {
	var contents []byte
	var err error

	if contents, err = ioutil.ReadFile(inputFilename); err != nil {
		log.Fatal(err)
	}

	var addressesSupportingTLSCount = 0
	var ipAddresses = strings.Split(strings.Trim(string(contents), "\n"), "\n")

addressLoop:
	for _, address := range ipAddresses {
		parts := strings.Split(strings.Replace(address, "]", "[", -1), "[")
		possiblyTLSAddess := false

		for i, part := range parts {
			if i%2 == 0 {
				// Normal sequence
				// If it contains an ABBA then it's a TLS address
				// (unless an ABBA is also discovered in an Hypernet sequence)
				if containsAbba(part) {
					possiblyTLSAddess = true
				}
			} else {
				// Hypernet sequence
				// Check if it contains an ABBA, if it does, continue
				if containsAbba(part) {
					continue addressLoop
				}
			}
		}

		if possiblyTLSAddess {
			addressesSupportingTLSCount++
		}
	}

	fmt.Printf(
		"Number of addresses supporting TLS: %v/%v\n",
		addressesSupportingTLSCount,
		len(ipAddresses),
	)
}
