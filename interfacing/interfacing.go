package interfacing

import (
	"fmt"
	"io"
	"strings"
	"time"
)

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
	b = &am
	fmt.Println(b.SayHello(" Evening"))

	var s str = "Kemon"
	fmt.Println(s.SayHello(" Morning"))
	b = s
	fmt.Println(b.SayHello(" Evening"))

	//in Go it is common to write methods that gracefully handle being called with a nil receiver

	var a *Id
	fmt.Println("value of a is :", a)
	fmt.Println(a.SayHello("testing nil receiver"))

	//if method of nil interface is called, and no implementations are there
	//it will throw runtime error as program wont understand which method to run
	//eg
	// var nilH Hello
	// nilH.SayHello("Pooy")

	//empty interface - to hold any value

	var emptytest empty
	emptytest = 1
	fmt.Printf("type is %T : %v \n", emptytest, emptytest)

	emptytest = "hello to empty"
	fmt.Printf("type is %T : %v \n", emptytest, emptytest)

	//type assertions

	fmt.Println("type assetion is :", emptytest.(string))

	//if not string then error then panic will happen, which can be managed like below
	if q, w := emptytest.(int); !w {
		fmt.Println("wrong type assertion")
	} else {
		fmt.Println(q)
	}

	//type switches
	TypeSwitch(1)
	TypeSwitch("Hello")

	//One of the most ubiquitous interfaces is Stringer defined by the fmt package.

	// type Stringer interface {
	// 	String() string
	// }

	//Errors

	//The error type is a built-in interface similar to fmt.Stringer
	// type error interface {
	// 	Error() string
	// }

	if err := run(); err != nil {
		fmt.Println(err)
	}

	//io.Reader interface which has method
	//Read([]byte) (int, error)
	r := strings.NewReader("Hello my dear bro")
	br := make([]byte, 8)
	for {
		n, err := r.Read(br) //n is the number of bytes populated, br will be populated
		fmt.Printf("Read : %q \n", br[:n])
		if err == io.EOF { //when end of file is reached
			break
		}
	}

}

type Hello interface {
	SayHello(string) string
}

type empty interface {
}

func (i *Id) SayHello(a string) string {
	if i == nil {
		return "nil reference"
	}
	return "Hello " + i.name + a
}

func (s str) SayHello(a string) string {
	return "Bye " + string(s) + a
}

func TypeSwitch(i interface{}) {

	switch v := i.(type) {
	case string:
		{
			fmt.Println("String ", v)
		}
	case int:
		{
			fmt.Println("Integer ", v)
		}
	}

}

type MyError struct {
	day   time.Time
	issue string
}

func (x *MyError) Error() string { //making MyError struct an implementation of error interface
	return fmt.Sprintf("happend at %v and error is %s", x.day, x.issue)
}

func run() error {
	return &MyError{time.Now(), "Nothing but testing"}
}
