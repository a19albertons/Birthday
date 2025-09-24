package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main(){
	// Enable parameter "n" and create help
	flag.Int("n", 0, "number of times and numbers")
	flag.Parse()

	bucle := flag.Args()
	

	// Create a loop to control the number of cases
	for i := 1; i<= len(bucle); i++ {
		
		age, err := strconv.Atoi(bucle[i-1])
		if err != nil {
			fmt.Println("Error: this number is invalid")
			continue
		}

		// Declare the candle variable
		candleCount := 0

		// Start counting every situation in binary which you get one rather than zero
		for age >= 2 {

			// This conditional controls whether your the rest is 1 or 0 to increase the candle variable
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