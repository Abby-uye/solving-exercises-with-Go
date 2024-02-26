package main

import "fmt"

func twelveDaysOfChristmas() {

	for days := 1; days <= 12; days++ {
		switch days {
		case 1:
			fmt.Println("On the first day of Christmas")
			break
		case 2:
			fmt.Println("On the second day of Christmas,")
			break
		case 3:
			fmt.Println("On the third day of Christmas,")
			break
		case 4:
			fmt.Println("On the fourth day of Christmas,")
			break
		case 5:
			fmt.Println("On the fifth day of Christmas,")
			break
		case 6:
			fmt.Println("On the sixth day of Christmas,")
			break
		case 7:
			fmt.Println("On the seventh day of Christmas,")
			break
		case 8:
			fmt.Println("On the eighth day of Christmas,")
			break
		case 9:
			fmt.Println("On the ninth day of Christmas,")
			break
		case 10:
			fmt.Println("On the tenth day of Christmas,")
			break
		case 11:
			fmt.Println("On the eleventh day of Christmas,")
			break
		case 12:
			fmt.Println("On the twelfth day of Christmas,")

		}
		fmt.Println("my true love sent to me")
		for gift := days; gift > 0; gift-- {

			switch gift {
			case 12:
				fmt.Println("Twelve drummers drumming,")

			case 11:
				fmt.Println("Eleven pipers piping,")

			case 10:
				fmt.Println("Ten lords a-leaping,")

			case 9:
				fmt.Println("Nine ladies dancing,")

			case 8:
				fmt.Println("Eight maids a-milking,")

			case 7:
				fmt.Println("Seven swans a-swimming,")

			case 6:
				fmt.Println("Six geese a-laying,")

			case 5:
				fmt.Println("Five golden rings,")

			case 4:
				fmt.Println("Four calling birds,\n")

			case 3:
				fmt.Println("Three French hens,\n")

			case 2:
				fmt.Println("Two turtle doves,\nand a partridge in a pear tree.")
				break

			case 1:
				fmt.Println("my true love sent to me\na partridge in a pear tree.")

			}
		}
	}
}
