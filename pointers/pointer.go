package pointers

import "fmt"

type Vertex struct {
	x int
	y int
}

var q *int

func DoPointers() {
	fmt.Println("\nLets do pointers")

	var v = Vertex{
		1, 2, //literals
	}
	x := 2
	q = &x
	fmt.Println("value of q is ", q)
	p := &v //p will be iniialized as *Vertex
	fmt.Printf("type of p is %T", p)
	p.x = 4 //should be *p.X but simplified in language

	fmt.Println(v)
}
