package concurrency

import (
	"fmt"
	"time"
)

/*
	A channel type is represented with the keyword `chan` followed by the type that are passed on the channel.
	Channel Direction: Send Only.
*/
func Pinger(channel chan<- string) {
	for i := 0; ; i++ {
		/*
			The <- operator is used to send and receive messages on the channel.
			Here's how to send a message on the channel.
		*/
		channel <- "Ping"
	}
}

/*
	Channel Direction: Send Only.
*/
func Ponger(channel chan<- string) {
	for i := 0; ; i++ {
		channel <- "Pong"
	}
}

/*
	Channel Direction: Receive only.
*/
func Printer(channel <-chan string) {
	for {
		/*
			Here's how to consume from the channel
		*/
		msg := <-channel
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}
