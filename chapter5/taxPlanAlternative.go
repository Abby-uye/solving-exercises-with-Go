package main

import "fmt"

func fairTax() {
	var totalExpense float32
	fmt.Println("Enter your housing expense: ")
	var housingExpense float32
	fmt.Scanln(&housingExpense)
	totalExpense += housingExpense
	fmt.Println("Enter your food expense: ")
	var foodExpense float32
	fmt.Scan(&foodExpense)
	totalExpense += foodExpense
	fmt.Println("Enter your clothing expense: ")
	var clothingExpense float32
	fmt.Scan(&clothingExpense)
	totalExpense += clothingExpense
	fmt.Println("Enter your transportation expense: ")
	var transportationExpense float32
	fmt.Scan(&transportationExpense)
	totalExpense += transportationExpense
	fmt.Println("Enter your healthcare expense:")
	var healthCareExpense float32
	fmt.Scan(&healthCareExpense)
	totalExpense += healthCareExpense
	fmt.Println("Enter your education Expense")
	var educationExpense float32
	fmt.Scan(&educationExpense)
	totalExpense += educationExpense
	fmt.Println("Enter your vacation expense")
	var vacationExpense float32
	fmt.Scan(&vacationExpense)
	totalExpense += vacationExpense

	var tax float32
	tax = (totalExpense / 23) * 100
	fmt.Println("Your tax for the year is: ", tax)
}
