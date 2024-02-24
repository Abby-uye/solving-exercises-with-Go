package main

import "fmt"

func main() {
	var accountNumber string
	fmt.Print("Please enter your account number")
	fmt.Scanln(&accountNumber)
	var balanceAtTheBeginningOfTheMonth float64
	fmt.Print("Enter your balance at the beginning of the month: ")
	fmt.Scan(&balanceAtTheBeginningOfTheMonth)
	var totalItemsCharged float64
	fmt.Print("Enter The total item charged this month: ")
	fmt.Scan(&totalItemsCharged)
	var totalCreditsApplied float64
	fmt.Print("Enter totalCredits Applied this month: ")
	fmt.Scan(&totalCreditsApplied)
	var allowedCreditLimit float64
	fmt.Print("Enter allowed credit limit: ")
	fmt.Scan(allowedCreditLimit)
	var creditBalance = (balanceAtTheBeginningOfTheMonth + totalItemsCharged) - totalCreditsApplied
	if creditBalance > allowedCreditLimit {
		fmt.Print("Credit Limit Exceeded")
	}

}
