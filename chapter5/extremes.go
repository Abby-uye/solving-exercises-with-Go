package main

import (
	"fmt"
	"math"
)

func extreme(numberOfNumbers int) {
	var minimum int = math.MaxInt
	var maximum int = 0
	var number int
	for count := 0; count < numberOfNumbers; count++ {
		fmt.Scan(&number)
		if number > maximum {
			maximum = number
		}
		if number < minimum {
			minimum = number
		}
	}
	fmt.Println("The maximum number is: ", maximum)
	fmt.Println("The minimum number is: ", minimum)
}
