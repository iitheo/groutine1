package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Go Channels Tutorial")
	//channels are a means of communication between goroutines

	values := make(chan string, 2)
	//unbuffered channel, it is blocking
	//
	defer close(values)

	go SendValue(values)
	go SendValue(values)

	SayHello("Carlos")
	SayHello("Paul")
	SayHello("Nkiru")
	value := <-values // receive value from channel

	SayHello("Chisom")
	fmt.Printf("Value: %s\n",value)
	SayHello("Theo")
}

func SendValue(c chan string){
	fmt.Printf("Executing Goroutine\n")
	time.Sleep(1 * time.Second)
	c <- "Hello World" // send a value to a channel
	fmt.Printf("Finished Executing Goroutine\n")
	time.Sleep(3 * time.Second)
}

func SayHello(name string){
	fmt.Printf("Hello %s!\n",name)
}
