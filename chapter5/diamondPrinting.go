package main

import "fmt"

func diamondPrinting() {
	for index := 1; index <= 5; index++ {
		for index1 := 5; index1 > 0; index1++ {
			fmt.Print(" ")
		}
		for index2 := 0; index2 < index; index2++ {
			fmt.Print("* ")

		}
		for index1 := 5; index1 > 0; index1++ {
			fmt.Print(" ")
		}

	}
}
