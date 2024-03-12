package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmt float64 = 1000
	returnAmt := 5.0
	var years float64 = 5
	var total = investmentAmt * math.Pow((1+returnAmt/100), years)
	// fmt.Println("Hello, World!")
	fmt.Println("Total amount is: ", total)
}
