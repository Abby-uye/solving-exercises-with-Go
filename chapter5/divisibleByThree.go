package main

func divisibleByThree() int {
	var sum int
	for number := 1; number <= 30; number++ {
		if number%3 == 0 {
			sum += number
		}
	}
	return sum
}
