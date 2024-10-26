package main

import (
	"fmt"
	"time"
)

// Go routines are lightweight threads of execution
// They are used to run functions concurrently
func main() {
	//dones := make([]chan bool, 4)
	done := make(chan bool)
	go greet("Hello, World!", done)
	go greet("How are you?", done)
	go slowGreet("Slow greet!", done)
	go greet("Fine!", done)
	fmt.Println("Waiting for slow greet to finish...", done)

	for _ = range done {
		<-done
	}
}

func greet(phrase string, doneChan chan bool) {
	fmt.Println(phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
	doneChan <- true
}
