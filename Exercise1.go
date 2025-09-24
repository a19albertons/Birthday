package main

import (
	"fmt"
)

func main(){
	fmt.Println("Number of cases:")
	var bucle int
	fmt.Scanln(&bucle)
	for i := 0; i< bucle; i++ {
		fmt.Println("Your age:")
		var age int
		fmt.Scanln(&age)
		candleCount := 0
		for age >= 2 {
			if age % 2 == 1 {
				candleCount++
			}
			age = age / 2
		}
		if age == 1 {
			candleCount++
		}
		fmt.Println(candleCount)
	}
}