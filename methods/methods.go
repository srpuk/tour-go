package methods

import "fmt"

type Vertex struct {
	x1 int
	x2 int
}

func DoMethods() {

	//A method is a function with a special receiver argument.
	//You cannot declare a method with a receiver whose type is defined in another
	// package (which includes the built-in types such as int).

	var x myInt = 1
	y := x.DoMath()
	fmt.Println("receiver function returned :", y)
	var v Vertex
	v.Scale() //methods with pointer receivers take either a value or a pointer as the receiver when they are called:
	fmt.Println("pointer receiver used to change value :", v)
	Scale2(&v) //functions with a pointer argument must take a pointer //The equivalent thing happens in the reverse direction.
	fmt.Println("pointer function used to change value :", v)

}

//	func (a int) DoMath() int {
//		return 0
//	} //will give an error
type myInt int

func (a myInt) DoMath() int {
	return 0
} //will give an error

func (v *Vertex) Scale() {
	v.x1 = 3
	v.x2 = 8
}

func Scale2(v *Vertex) {
	v.x1 = 8
	v.x2 = 3
}

//Value or pointer receiver
//The first is so that the method can modify the value that its receiver points to.
//The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
//In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both
