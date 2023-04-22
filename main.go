package main

import (
	"fmt"

	"airnudge.com/learn/arrays"
	"airnudge.com/learn/loops"
	"airnudge.com/learn/maths"
	"airnudge.com/learn/pointers"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Print("Hello")

	defer fmt.Println("\nDefer statement running from main") // deferred calls are executed in last-in-first-out order.

	maths.DoMaths()
	loops.DoLooping()
	pointers.DoPointers()
	arrays.DoArrays()

}
