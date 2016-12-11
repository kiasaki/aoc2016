package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const inputFilename = "input"

func main() {
	var contents []byte
	var err error

	if contents, err = ioutil.ReadFile(inputFilename); err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(strings.Trim(string(contents), "\n"), "\n")

	fmt.Printf("XXX: %v\n", "")
}
