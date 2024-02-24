//package main
//
//import "fmt"
//
//func main() {
//	fmt.Print("How many trips did you go ")
//	var numberOfTrips int
//	fmt.Scan(&numberOfTrips)
//	var totalNumberOfGallons float64
//	var totalNumberOfMiles float64
//	for index := 1; index <= numberOfTrips; index++ {
//		var miles float64
//		var numberOfGallons float64
//
//		fmt.Print("Enter the miles for  trip ", index, ": ")
//		fmt.Scan(&miles)
//
//		fmt.Print("Enter the number of gallons used for  trip ", index, ": ")
//		fmt.Scan(&numberOfGallons)
//
//		milesPerGallon := miles / numberOfGallons
//		fmt.Println("The number of miles per gallon for trip ", index, " is ", milesPerGallon)
//
//		totalNumberOfGallons += numberOfGallons
//		totalNumberOfMiles += miles
//	}
//
//	if totalNumberOfGallons != 0 {
//		totalMilesPerGallon := totalNumberOfMiles / totalNumberOfGallons
//		fmt.Println("The total miles per gallon obtained is ", totalMilesPerGallon)
//	} else {
//		fmt.Println("No gallons entered. Cannot calculate total miles per gallon.")
//	}
//}
