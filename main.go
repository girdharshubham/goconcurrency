package main

import (
	"fmt"
	"goconcurrency/concurrency"
)

func main() {
	/*
		Channels provide a way for two goroutines to communicate with one another and synchronize their execution.
		Channels make goroutines share memory by communicating.
		A channel can be viewed as an internal FIFO queue within a program.
	 */
	channel := make(chan string)
	go concurrency.Pinger(channel)
	go concurrency.Printer(channel)
	go concurrency.Ponger(channel)

	var input string
	fmt.Scanln(&input)

}
