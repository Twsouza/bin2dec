package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
)

var (
	isBinary = regexp.MustCompile("^[0-1]+$").MatchString
)

func main() {
	var binarys string
	fmt.Println("Enter the binary:")
	if _, err := fmt.Scanf("%s", &binarys); err != nil {
		log.Fatalf("Could not read from input, error %v", err)
	}

	if !isBinary(binarys) {
		log.Fatalf("%s is not a valid binary!", binarys)
	}

	decimal := binToDec(binarys)

	fmt.Println("Decimal is:", decimal)
}

func binToDec(bin string) int {
	var decimal int
	for i := len(bin) - 1; i >= 0; i-- {
		b, _ := strconv.Atoi(fmt.Sprintf("%c", bin[i]))
		s := math.Pow(2, float64(len(bin)-1-i))

		decimal += b * int(s)
	}

	return decimal
}
