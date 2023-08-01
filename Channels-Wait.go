package main

import (
	"fmt"
	"sync"
)

func removeNumerals(input string, wg *sync.WaitGroup) {
	defer wg.Done()
	var numeralFreeString string
	for _, char := range input {
		if !('0' <= char && char <= '9') {
			numeralFreeString += string(char)
		}
	}
	fmt.Println(numeralFreeString)
}

func main() {
	inputStrings := []string{"gopher123", "alpha99beta", "1cita2del3"}

	var wg sync.WaitGroup

	for _, input := range inputStrings {
		wg.Add(1)
		go removeNumerals(input, &wg)
	}

	wg.Wait()
}
