package main

import (
	"fmt"
	"time"
)

func doWork(channel chan<- bool) {
	fmt.Print("working...")
    time.Sleep(time.Second)
		fmt.Println("done")
		
		channel <- true
}

func main() {
	// Creating a buffed channel
	messages := make(chan string, 2)

	// go starts a goroutine, and this is an anonimous function that will send data to channel `messages`
	go func() { messages <- "ping" }()

	messages <- "UwU"

	fmt.Println(<-messages) // here we receive from the channel
	fmt.Println(<-messages) // here we receive from the channel

	//
	// Channel synchronization
	// (https://gobyexample.com/channel-synchronization)
	// This is neat example of a async worker

	done := make(chan bool, 1)
	go doWork(done)

	fmt.Println(<-done)
}
