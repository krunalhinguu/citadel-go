package main

import (
	"fmt"
	"time"
)

func printAlphabet(letter rune) {
	fmt.Printf("%c ", letter)
}

func main() {
	var start, end int

	fmt.Print("Enter the start position (between 1 and 26): ")
	_, err := fmt.Scan(&start)
	if err != nil || start < 1 || start > 26 {
		fmt.Println("Invalid input for the start position.")
		return
	}

	fmt.Print("Enter the end position (between 1 and 26): ")
	_, err = fmt.Scan(&end)
	if err != nil || end < 1 || end > 26 {
		fmt.Println("Invalid input for the end position.")
		return
	}

	if start > end {
		fmt.Println("Start position cannot be greater than the end position.")
		return
	}

	// Create a channel to signal the completion of goroutines
	done := make(chan struct{})

	// Launch a goroutine for each alphabet between start and end
	for i := start; i <= end; i++ {
		go func(letter rune) {
			printAlphabet(letter)
			done <- struct{}{}
		}(rune('a' - 1 + i))
	}

	// Wait for all goroutines to complete
	for i := start; i <= end; i++ {
		<-done
	}

	// Add some sleep to ensure all goroutines have enough time to complete
	time.Sleep(100 * time.Millisecond)
}

	
/* 

The alphabets are not printed in order because goroutines are executed concurrently, and the scheduling of goroutines is managed by the Go runtime. When goroutines are launched to print the alphabets, there's no guarantee about the order in which they will complete their execution. The Go scheduler may interleave the execution of different goroutines on different CPU cores, and their execution time may vary based on factors such as system load, CPU scheduling decisions, and other runtime conditions.

Since the goroutines are launched concurrently and independently, they are not synchronized to ensure any specific ordering. As a result, the output of the program may appear non-sequential, as the order of execution is non-deterministic.

*/ 