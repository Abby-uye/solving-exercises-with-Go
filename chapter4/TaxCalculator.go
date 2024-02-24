package main

import "fmt"

func main() {
	for index := 0; index < 3; index++ {
		var citizen string
		fmt.Scanln("Enter your name")
		var citizenEarnings float32
		fmt.Scan("Enter your total earning for the year")
		var taxForThatYear float32
		if citizenEarnings > 0 && citizenEarnings <= 30000 {
			taxForThatYear = (citizenEarnings / 15) * 100
			fmt.Println("Dear "+citizen+"Your tax for this year is ", taxForThatYear)
		} else {
			taxForThatYear = (citizenEarnings / 20) * 100
			fmt.Println("Dear "+citizen+"Your tax for this year is ", taxForThatYear)

		}

	}
}
