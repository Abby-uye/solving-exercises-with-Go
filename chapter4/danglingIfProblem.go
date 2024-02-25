package main

import "fmt"

func main() {
	var x int
	var y int
	if x > 5 {
		if y > 5 {
			fmt.Println("X and y are greater than 5")
		}
	} else {
		fmt.Println("X is < or equal to 5")
	}
}
