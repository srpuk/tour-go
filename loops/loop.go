package loops

import (
	"fmt"
	"math"
	"runtime"
)

func DoLooping() {

	for i := 0; i < 5; i++ { //init and post statements are optional.
		defer fmt.Print(i) //no condition means infinite loop //LIFO concept of defer

	}
	fmt.Print("\n")

	if v := math.Pow(2, 2); v < 10 {
		fmt.Println(v) //scope of v is just inside if, but avaialble inside else also
	}

	switch x := runtime.GOOS; x {
	case "darwin":
		fmt.Println(x) //break is provided automatically by GO
	case "linux":
		fmt.Println(x)
	default:
		fmt.Print("Came to default \n")
	}

	switch {
	case math.Pow(2, 2) < 5:
		fmt.Print("Correct \n")

	case math.Pow(2, 2) > 5:
		fmt.Print("Incorrect \n")
	}

}
