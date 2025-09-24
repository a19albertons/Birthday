package main

import (
	"fmt"
)

func main(){
	// Notify, create the var and request the number of cases
	fmt.Println("Number of cases:")
	var bucle int
	fmt.Scanln(&bucle)

	// Create a loop to control the number of cases
	for i := 0; i< bucle; i++ {
		
		// Now, the program requests your age 
		fmt.Println("Your age:")
		var age int
		fmt.Scanln(&age)

		// Declare the candle variable
		candleCount := 0

		// Start counting to every situation in binary which you get one rather than zero
		for age >= 2 {

			// Control whether your the rest is 1 to increase the candle variable
			if age % 2 == 1 {
				candleCount++
			}

			// Reduce the age its half
			age = age / 2
		}

		// Control whether last number is 1 or 0
		if age == 1 {
			candleCount++
		}

		// Print the number of candles
		fmt.Println(candleCount)
	}
}