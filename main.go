package main

import (
	"fmt"
	"log"
)

func main() {
	var binarys string
	log.Println("Enter the binary:")
	if _, err := fmt.Scanf("%s", &binarys); err != nil {
		log.Fatalf("Could not read from input, error %v", err)
	}

	log.Println("You entered ", binarys)
}
