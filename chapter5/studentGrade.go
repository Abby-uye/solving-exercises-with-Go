package main

import "fmt"

func studentGrade() {
	var aGrade int
	var bGrade int
	var cGrade int
	var dGrade int
	for student := 0; student < 5; student++ {
		fmt.Println("Enter your name: ")
		var name string
		fmt.Scanln(&name)
		fmt.Println("Enter your Grade: ")
		var letterGrade string
		fmt.Scanln(&letterGrade)

		switch letterGrade {
		case "A":
			aGrade++
		case "B":
			bGrade++
		case "C":
			cGrade++
		case "D":
			dGrade++
		}
	}
	fmt.Println("The number of students with A grade is: ", aGrade)
	fmt.Println("The number of students with B grade is: ", bGrade)
	fmt.Println("The number of students with B grade is: ", cGrade)
	fmt.Println("The number of students with D grade is: ", dGrade)
}
