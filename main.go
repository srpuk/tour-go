package main

import (
	"fmt"

	"airnudge.com/learn/collections"
	"airnudge.com/learn/interfacing"
	"airnudge.com/learn/loops"
	"airnudge.com/learn/maths"
	"airnudge.com/learn/methods"
	"airnudge.com/learn/pointers"
	"airnudge.com/learn/typeparameters"
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
	collections.DoArrays()
	collections.TestFunctionValues()
	methods.DoMethods()
	interfacing.DoInterfacing()
	typeparameters.DoTypeTest()

}
