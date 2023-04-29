package concurrency

import (
	"fmt"
	"time"
)

func DoConcurrency() {
	ch := make(chan string)
	ch2 := make(chan string)
	fmt.Println("Lets do some concurrency")

	Say("sir") //simple go func

	go HeyChannel(" Sreeroop", ch)
	x := <-ch
	fmt.Println(x) //using a channel to receive data

	go SayLoop(" people", ch2)
	for i := range ch2 {
		fmt.Println(i) //waiting for all the values
	}

	//Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	//ch := make(chan string,100)

	//Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

	//The select statement lets a goroutine wait on multiple communication operations.

	//A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

	fib := make(chan int)
	quit := make(chan int)

	go func() { //this func is listening to the fib channel
		for i := 0; i < 10; i++ {
			fmt.Println(<-fib)
		}
		quit <- 0 //here updating the quit channel

	}()
	fibanocci(fib, quit)

}

func Say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello ", s)
	}
}
func HeyChannel(s string, c chan string) {
	c <- "Hello" + s

}
func SayLoop(s string, c chan string) {
	for j := 0; j < 3; j++ {
		c <- "Hello" + s
	}
	close(c) //required if receiver must be told there are no more values coming

}

func fibanocci(f, q chan int) {
	x, y := 0, 1
	for {
		select {
		case f <- x:
			x, y = y, x+y
		case <-q:
			fmt.Println("quit")
			return //important for exiting the for loop
		default:
			fmt.Println("waiting...")
			time.Sleep(500 * time.Millisecond)

		}
	}
}
