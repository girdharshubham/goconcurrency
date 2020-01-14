# Concurrency

## Goroutines
A goroutine is a function that is capable of running concurrently with other functions. To create a goroutine use the
keyword `go` followed by a function invocation. Each Goroutine is created initially with a stack size of **2KB**

## Channels
Channels can be thought os as typed FIFO queues through which you can send and receive values with the channel operator
, <-. The data flows in the direction of the <- operator

## Running the code
 **Build the code**: ```go build main.go```  
 **Running the code**: ```./main```