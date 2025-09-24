package main

import (
	"fmt"
)

func main(){
	// I notify, create the var and request the number of cases
	fmt.Println("Number of cases:")
	var bucle int
	fmt.Scanln(&bucle)

	// I create a bucle for to control the number of cases
	for i := 0; i< bucle; i++ {
		
		// Now, I request your age 
		fmt.Println("Your age:")
		var age int
		fmt.Scanln(&age)

		// I declare the candle variable
		candleCount := 0

		// I start counting to every situation in binary which you get one rather than zero
		for age >= 2 {

			// I control whether your the rest is 1 to increase the candle variable
			if age % 2 == 1 {
				candleCount++
			}

			// I reduce the age its half
			age = age / 2
		}

		// I control whether last number is 1 or 0
		if age == 1 {
			candleCount++
		}

		// I print the number of candles
		fmt.Println(candleCount)
	}
}