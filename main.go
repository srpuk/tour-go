package main

import (
	"fmt"

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

	var v = Vertex{
		1, 2,
	}
	p := &v
	p.X = 4 //should be *p.X but simplified in language

	fmt.Println(v)

	defer fmt.Println("\nDefer statement running from main") // deferred calls are executed in last-in-first-out order.

	maths.DoMaths()
	loops.DoLooping()
	pointers.DoPointers()

}
