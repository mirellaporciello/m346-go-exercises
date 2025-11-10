package main
 
import (

	"fmt"

	"math/rand"

	"os"

	"time"

)
 
func main() {

	rand.Seed(time.Now().UnixNano()) // seed the random generator

	var eyes = rand.Intn(6) + 1

	var when = time.Now()
 
	// open or create files

	eyesFile, err := os.Create("eyes.txt")

	if err != nil {

		fmt.Println("Error creating eyes.txt:", err)

		return

	}

	defer eyesFile.Close()
 
	logFile, err := os.OpenFile("dice.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {

		fmt.Println("Error opening dice.log:", err)

		return

	}

	defer logFile.Close()
 
	// write to files using fmt.Fprintln

	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")

	fmt.Fprintln(logFile, "the dice was rolled at", when)
 
	// optional: also print to console

	fmt.Println("Results written to eyes.txt and dice.log")

}

 