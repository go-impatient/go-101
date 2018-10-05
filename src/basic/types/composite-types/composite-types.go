package main

import (
	"math/cmplx"
	"math"
	"fmt"
)

func euler()  {
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
}
func main() {
	euler()
}



