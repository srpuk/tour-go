package interfacing

import "fmt"

type Id struct {
	name string
	id   int
}

type str string

func DoInterfacing() {
	fmt.Println("Lets do some interfacing")
	//An interface type is defined as a set of method signatures.
	//implicit implementation

	am := Id{
		"Sreeroop",
		3737,
	}
	fmt.Println(am.SayHello(" Morning"))

	var b Hello
	b = am
	fmt.Println(b.SayHello(" Evening"))

	var s str = "Kemon"
	fmt.Println(s.SayHello(" Morning"))
	b = s
	fmt.Println(b.SayHello(" Evening"))

}

type Hello interface {
	SayHello(string) string
}

func (i Id) SayHello(a string) string {
	return "Hello " + i.name + a
}

func (s str) SayHello(a string) string {
	return "Bye " + string(s) + a
}
