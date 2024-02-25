package main

import "fmt"

func main() {
	var x int = 9
	var y int = 9
	if x < 10 {
		if y > 10 {
			fmt.Println("*****")
		}
	} else {
		fmt.Println("#####")
		fmt.Println("$$$$$")
	}
	var m int = 9
	var n int = 11
	if m < 10 {
		if n > 11 {
			fmt.Println("*****")
		}
	} else {
		fmt.Println("#####")
		fmt.Println("$$$$$")
	}
}
