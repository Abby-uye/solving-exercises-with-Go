package main

import "fmt"

func main() {
	var largest int = 0
	for counter := 0; counter < 10; counter++ {
		var number int
		fmt.Scan("Please enter a number")
		if number > largest {
			largest = number
		}
	}
	fmt.Println("The largest number is ", largest)
}
