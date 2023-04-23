package collections

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

	s := [4]int{1, 2, 3, 4} //this declaaration is different from others eg: s:="Hello"

	fmt.Println("First slice :", s[0:4])
	fmt.Println("Second slice :", s[0:])
	fmt.Println("Third slice :", s[:4])
	fmt.Println("Fourth slice :", s[:]) // default lowebound and upper bound

	t := []int{3, 5, 6, 7, 9}
	PrintLengCap(t)
	PrintLengCap(t[:3])
	//PrintLengCap(t[:7])
	//panic: runtime error: slice bounds out of range [:7] with capacity 5
	PrintLengCap(t[1:5]) //The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	var p []int
	PrintLengCap(p)
	if p == nil {
		fmt.Println("zero value of a slice is nil")
	}

	makeslice := make([]int, 4) //using make to create slice
	PrintLengCap(makeslice)
	makeslicecap := make([]int, 4, 7) //4 length, 7 capacity
	PrintLengCap(makeslicecap)

	//slice of slices
	q := [][]int{
		[]int{1, 2, 4},
		[]int{3, 4, 5},
	}
	fmt.Println(q)
	q[1][1] = 9000

	fmt.Println(q)

	//usage of append

	var n []int
	n = append(n, 4, 5, 6, 9)
	PrintLengCap(n) //underlying array is flexible

	//range

	for x, y := range n {
		fmt.Println("index : ", x, "value : ", y)
	}

	//map

	var mapexample1 map[string]string
	if mapexample1 == nil {
		fmt.Println("The zero value of a map is nil")
	}
	mapsexample := make(map[string]string, 5)
	mapsexample["Anjana"] = "Thrissur"
	mapsexample["Sreeroop"] = "Kozhikode"

	fmt.Println(mapsexample["Anjana"])

	mapexample2 := map[string]string{ //using literals with key
		"anjana":   "thrissur",
		"sreeroop": "kozhikode",
		"vishruth": "ramanattukar",
	}
	fmt.Println(mapexample2["sreeroop"])

	delete(mapexample2, "vishruth")
	fmt.Println("value for sreeroop after deleteion :", mapexample2["vishruth"])

	elem, ok := mapexample2["sreeroop"] //checking value is present in map
	if ok {
		fmt.Println(elem)
	}

}

func PrintLengCap(s []int) {
	fmt.Println("Size is", len(s), "/n capacity is ", cap(s))
}
