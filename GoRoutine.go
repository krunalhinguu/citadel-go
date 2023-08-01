package main

import (  
	"time"
    "fmt"
)

func hello() {  
    fmt.Println("Hello world goroutine")
}
func main() {  
    go hello()
	time.Sleep(1 * time.Second)
    fmt.Println("main function")
}