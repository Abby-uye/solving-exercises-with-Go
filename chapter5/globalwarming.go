package main

import "fmt"

func climateChange() {
	var correctAnswerCount int
	var answer string
	fmt.Println(`What is the primary cause of global warming?

a) Volcanic activity
b) Solar radiation
c) Human activities, such as burning fossil fuels
d) Deforestation`)
	fmt.Scanln(&answer)
	if answer == "c" {
		correctAnswerCount++
	}
	fmt.Println(`Which greenhouse gas is most responsible for trapping heat in the Earth's atmosphere?

a) Oxygen (O2)
b) Methane (CH4)
c) Nitrogen (N2)
d) Carbon dioxide (CO2)`)
	fmt.Scanln(&answer)
	if answer == "D" {
		correctAnswerCount++
	}
	fmt.Println(`How does deforestation contribute to global warming?

a) It releases oxygen into the atmosphere
b) It increases the Earth's albedo
c) It reduces the absorption of solar radiation
d) It decreases the concentration of greenhouse gases
`)
	fmt.Scanln(&answer)
	if answer == "C" {
		correctAnswerCount++
	}
	fmt.Println(`What is the main impact of global warming on sea levels?

a) Sea levels rise due to the melting of polar ice caps
b) Sea levels fall due to increased precipitation
c) Sea levels remain unchanged
d) Sea levels rise due to increased salinity`)
	fmt.Scanln(&answer)
	if answer == "a" {
		correctAnswerCount++
	}
	fmt.Println(`Which of the following is a potential consequence of global warming?

a) Decreased frequency of extreme weather events
b) Expansion of polar ice caps
c) Ocean acidification
d) Reduced biodiversity`)
	fmt.Scanln(&answer)
	if answer == "a" {
		correctAnswerCount++
	}
	if correctAnswerCount == 5 {
		fmt.Println("Excellent")
	} else if correctAnswerCount == 4 {
		fmt.Println("Very good")
	} else if correctAnswerCount < 4 {
		fmt.Println("Time to brush up on your knowledge of global warming")
	}
}
