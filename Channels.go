package main

import (
	"fmt"
)

func removeNumerals(input string, resultChan chan<- string) {
	var numeralFreeString string
	for _, char := range input {
		if !('0' <= char && char <= '9') {
			numeralFreeString += string(char)
		}
	}
	resultChan <- numeralFreeString
}

func main() {
	inputStrings := []string{"gopher123", "alpha99beta", "1cita2del3"}

	resultChan := make(chan string)
	defer close(resultChan)

	for _, input := range inputStrings {
		go removeNumerals(input, resultChan)
	}

	var numeralFreeStrings []string
	for range inputStrings {
		numeralFreeStrings = append(numeralFreeStrings, <-resultChan)
	}

	fmt.Println("Input Strings:", inputStrings)
	fmt.Println("Numeral-free Strings:", numeralFreeStrings)
}
