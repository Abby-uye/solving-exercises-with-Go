package main

import "fmt"

func main() {
	var weeklySalary float32 = 200
	var totalAmountSold float32
	var numberOfItemsSold int
	fmt.Scan(&numberOfItemsSold)
	if numberOfItemsSold == 0 {
		fmt.Println("Unfortunately ")
	} else {
		for i := 0; i < numberOfItemsSold; i++ {
			fmt.Println("Enter the name of the item you sold:")
			fmt.Println("Enter the amount of the item sold")
			var itemValue float32
			fmt.Scan(&itemValue)
			totalAmountSold += itemValue
		}
		var salary float32
		salary = (totalAmountSold/9)*1000 + weeklySalary
		fmt.Println("Your pay for the week is ", salary)
	}
}
