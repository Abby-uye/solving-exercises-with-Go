package main

import "fmt"

func main() {
	var largest int
	var secondLargest int
	for counter := 0; counter < 10; counter++ {
		var number int
		fmt.Scan("Please enter a number")
		if number > largest {
			secondLargest = largest
			largest = number

		}
	}
	fmt.Println("The largest number is ", largest)
	fmt.Println("The second largest number ", secondLargest)
}
