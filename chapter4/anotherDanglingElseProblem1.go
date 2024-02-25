package main

import "fmt"

func main() {
	var x int = 5
	var y int = 5

	if y == 8 {
		if x == 5 {
			fmt.Println("@@@@@")
			fmt.Println("$$$$$")
			fmt.Println("&&&&&")
		}
	} else {
		fmt.Println("#####")
	}

	fmt.Println("==================================================")
	if y == 8 {
		if x == 5 {
			fmt.Println("@@@@@")
		}
	} else {
		fmt.Println("#####")
		fmt.Println("$$$$$")
		fmt.Println("&&&&&")
	}
	fmt.Println("==================================================")
	if y == 8 {
		if x == 5 {
			fmt.Println("@@@@@")
			fmt.Println("&&&&&")
		}

	} else {
		fmt.Println("#####")
		fmt.Println("$$$$$")
	}

	fmt.Println("==================================================")
	y = 7
	if y == 7 {
		if x == 5 {
			fmt.Println("#####")
			fmt.Println("$$$$$")
			fmt.Println("&&&&&")

		}
	} else {
		fmt.Println("@@@@@")

	}

}
