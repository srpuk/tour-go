package maths

import (
	"fmt"
	"math"
	"math/cmplx"
)

var MaxInt uint64 = (1 << 64) - 1
var Ufloat float32 = 12.45

const Pi = 3.14

func DoMaths() {

	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt(7))
	fmt.Println(cmplx.Sqrt(-5 + 12i))
	fmt.Println(MaxInt)        //shift 1<<2 means 100 in binary
	fmt.Println(int32(Ufloat)) //type conversion
	x := 65.6
	fmt.Printf("Printing type of Ufloat %T \n", x)
	fmt.Println("const value", Pi)

}
