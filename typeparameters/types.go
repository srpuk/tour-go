package typeparameters

import "fmt"

func Index[T comparable](sl []T, s T) []T { //both sl and s of same type
	//sl should be comaprable //constraints //instead of comparable 'any' also can be used

	sl = append(sl, s)
	return sl
}

func DoTypeTest() {
	fmt.Println("Lets do type parameters")

	var s = []string{
		"Hello", "hai",
	}
	d := "bye"
	fmt.Println(Index(s, d))

	g := Generictype[int]{1, 2, 3} //declaring Generictype
	g.Print("testetst")

}

type Generictype[T any] []T

func (p Generictype[T]) Print(a string) {

	fmt.Println(a, p)
}
