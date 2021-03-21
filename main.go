package main

import (
	"fmt"
	"log"
	"regexp"
)

var (
	isBinary = regexp.MustCompile("^[0-1]+$").MatchString
)

func main() {
	var binarys string
	log.Println("Enter the binary:")
	if _, err := fmt.Scanf("%s", &binarys); err != nil {
		log.Fatalf("Could not read from input, error %v", err)
	}

	if !isBinary(binarys) {
		log.Fatalf("%s is not a valid binary!", binarys)
	}

	log.Println("You entered ", binarys)
}
