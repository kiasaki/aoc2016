package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	var index = 0
	var input = "cxdnnyjw"
	var password = ""

	for len(password) < 8 {
		checksum := md5.Sum([]byte(input + strconv.Itoa(index)))
		hexChecksum := fmt.Sprintf("%x", checksum)

		if string(hexChecksum[0:5]) == "00000" {
			fmt.Println(hexChecksum)
			password += string(hexChecksum[5])
		}

		index++
	}

	fmt.Printf("Password is: %v\n", password)
}
