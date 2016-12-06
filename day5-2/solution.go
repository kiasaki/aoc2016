package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	var index = 0
	var charactersFound = 0
	var input = "cxdnnyjw"
	var password = make([]rune, 8)

	for charactersFound < 8 {
		checksum := md5.Sum([]byte(input + strconv.Itoa(index)))
		hexChecksum := fmt.Sprintf("%x", checksum)

		if string(hexChecksum[0:5]) == "00000" {
			character := hexChecksum[6]
			position, err := strconv.Atoi(string(hexChecksum[5]))
			if err != nil {
				index++
				continue
			}

			if position < len(password) && password[position] == rune(0) {
				fmt.Println(hexChecksum)
				password[position] = rune(character)
				charactersFound++
			}
		}

		index++
	}

	fmt.Printf("Password is: %v\n", string(password))
}
