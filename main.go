package main

import "fmt"

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

}
