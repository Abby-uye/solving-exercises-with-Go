package main

import "fmt"

func main() {
	var astericsk int
	fmt.Println("Enter a number of asterisk that will define the base of a triangle note: It has to be between 1 and 10")
	fmt.Scan(&astericsk)
	for base := 0; base < astericsk; base++ {
		for base1 := 0; base1 < base; base1++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
