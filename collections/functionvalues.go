package collections

import "fmt"

func TestFunctionValues() {

	fmt.Println("Functions are values too. They can be passed around just like other values.")

	hypot := func(x int, y int) int {
		return x * y
	}

	fmt.Println(hypot(3, 4))

	compute(hypot) //passing function as value

	pos, neg := addr(), addr() //initilaizing pos and neg as functions
	fmt.Println(pos(2))        //pos will set sum as 2 //sum is now set as local variable of pos function
	fmt.Println(neg(2))
	fmt.Println(pos(3))
	//fmt.Println(neg(3))
	fmt.Println(pos(5))
	fmt.Println(neg(5))

}

func compute(fn func(s int, t int) int) int {
	fmt.Println("using the function received ", fn(8, 9))
	return fn(8, 9)
}

//function closures

func addr() func(int) int {
	sum := 0
	//A closure is a function value that references variables from outside its body.
	//The function may access and assign to the referenced variables;
	//in this sense the function is "bound" to the variables.

	return func(i int) int {
		sum += i
		return sum
	}
}
