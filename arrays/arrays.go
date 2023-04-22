package arrays

import "fmt"

func DoArrays() {
	fmt.Println("Lets do arrays")

	var x [4]int //initalization one way
	x[0] = 8
	x[1] = 90

	prime := [4]int{1, 2, 3, 4} //initalization second way

	for i := 0; i < 4; i++ {
		fmt.Print(prime[i])

	}

	var a []string
	a = append(a, "Hello")
	a = append(a, "Hai")

	for _, x := range a {
		fmt.Println(x)
	}

	//Changing the elements of a slice modifies the corresponding elements
	//of its underlying array.Other slices that share the same underlying array will see those changes.

	test := [4]int{3, 5, 7, 9}
	b := test[0:3] //lower bound and higher bound will be defaulted to 0 and size of slice respeictively
	c := test[1:4] //start included end not included
	fmt.Println(b, c)
	test[1] = 999
	fmt.Println(b, c)

}
