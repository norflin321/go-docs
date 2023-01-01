package main

import (
	"fmt"
	"time"
)

func unbufferedChannelExample() {

	// create a new channel
	messages := make(chan string, 0)

	// invoke a function in a goroutine, and it will run ansynchronously
	go func() {
		fmt.Println("lets sleep a bit") // this or next print will be first, we cant know which one
		time.Sleep(5 * time.Second)

		// send a value into a channel
		messages <- "ping"
	}()

	fmt.Println("we just invoked a goroutine") // this or previous print will be first, we cant know which one

	// receive the value from the channel
	// here is execution of our "thread" (actually not),
	// but we will await here until we get some data from the channel
	msg := <-messages
	fmt.Println(msg)

	fmt.Println("program finish")
}

func bufferedChannelExample() {
	channel := make(chan string, 0)

	go func() {
		fmt.Println("before ping send")

		// if we had unbuffered channel, than below code would not execute (i think)
		// since we are not receiving any data anywhere from the channel
		// but if it is buffered channel, than this data is saved and we can receive it later
		// and only after we received it, this go routine will continue execution
		channel <- "ping"
		fmt.Println("after ping send")
	}()

	fmt.Println("lets sleep a bit")
	time.Sleep(2 * time.Second)

	// fmt.Println(<-channel)

	fmt.Println("program finish")
}

// Channel Synchronization
// https://gobyexample.com/channel-synchronization
