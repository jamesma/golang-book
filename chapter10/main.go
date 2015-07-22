package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)

		// wait between 0 and 250 ms
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// func pinger(c chan<- string) is more precise, c can only be sent to
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" // send "ping" to the channel
	}
}

// func ponger(c chan<- string) is more precise, c can only be sent to
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong" // send "pong" to the channel
	}
}

// func printer(c <-chan string) is more precise, c can only be received
func printer(c chan string) {
	for {
		msg := <-c // receive a message from the channel and store it in msg
		fmt.Println(msg)

		// the two lines above can be shortened to fmt.Println(<-c)

		time.Sleep(time.Second * 1)
	}
}

func sleep(secondsDuration int) {
	for i := 0; i < secondsDuration; i++ {
		<-time.After(time.Second)
	}
}

func main() {

	// goroutines
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input1 string
	fmt.Scanln(&input1)

	// channels
	// prints "ping" and "pong" forever until NL
	c := make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input2 string
	fmt.Scanln(&input2)

	// select, works like switch but for channels
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	// select picks the first channel that is ready and receives from it
	// (or sends to it). if more than one of the channels are aready then it
	// randomly picks one to receive from. if none of the channels are ready,
	// the statement blocks until one becomes available.

	var input3 string
	fmt.Scanln(&input3)

	// select can be used to implement a timeout
	// time.After creates a channel and after the given duration it will
	// send the current time on it
	select {
	case msg1 := <-c1:
		fmt.Println("Message 1", msg1)
	case msg2 := <-c2:
		fmt.Println("Message 2", msg2)
	case <-time.After(time.Second): // result of time.After not stored
		fmt.Println("Timeout")
	}

	// buffered channels
	// bufferedChannel := make(chan int, 1) // capacity of 1
	// normally channels are synchronous - both sides of the channel will
	// wait until the other side is ready.
	// a buffered channel is asynchronous - sending or receiving a message
	// will not wait unless the channel is already full.

	fmt.Println("Sleeping for 10 seconds.")
	sleep(10)
	fmt.Println("Slept for 10 seconds.")
}
