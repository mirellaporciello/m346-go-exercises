package main

import "fmt"

func main() {
firstName := "Mirella"
    lastName := "Porciello"
    dayOfBirth := 12
    monthOfBirth := 4
    yearOfBirth := 2008
    numberOfSiblings := 1
    heightInMeters := 1.50
    zodiacSign := 'W'
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
