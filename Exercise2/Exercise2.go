package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main(){
	flag.Int("n", 0, "numero esperado de argumentos posicionales")
	flag.Parse()

	bucle := flag.Args()
	

	// I create a bucle for to control the number of cases
	for i := 1; i<= len(bucle); i++ {
		
		age, err := strconv.Atoi(bucle[i-1])
		if err != nil {
			fmt.Println("Error: argumento no es un número entero")
			continue
		}

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