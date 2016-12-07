package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const inputFilename = "input"

func extractABAs(seq string) []string {
	abas := []string{}

	for i := 0; i <= len(seq)-3; i++ {
		if seq[i] == seq[i+2] && seq[i] != seq[i+1] {
			abas = append(abas, seq[i:i+3])
		}
	}

	return abas
}

func supportsSSL(abas, babs []string) bool {
	for _, aba := range abas {
		targetBab := string(aba[1]) + string(aba[0]) + string(aba[1])
		for _, bab := range babs {
			if bab == targetBab {
				return true
			}
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

	var addressesSupportingSSLCount = 0
	var ipAddresses = strings.Split(strings.Trim(string(contents), "\n"), "\n")

	for _, address := range ipAddresses {
		parts := strings.Split(strings.Replace(address, "]", "[", -1), "[")
		supernetABAs := []string{}
		hypernetBABs := []string{}

		for i, part := range parts {
			if i%2 == 0 {
				// Supernet sequence
				supernetABAs = append(supernetABAs, extractABAs(part)...)
			} else {
				// Hypernet sequence
				hypernetBABs = append(hypernetBABs, extractABAs(part)...)
			}
		}

		if supportsSSL(supernetABAs, hypernetBABs) {
			addressesSupportingSSLCount++
		}
	}

	fmt.Printf(
		"Number of addresses supporting SSL: %v/%v\n",
		addressesSupportingSSLCount,
		len(ipAddresses),
	)
}
