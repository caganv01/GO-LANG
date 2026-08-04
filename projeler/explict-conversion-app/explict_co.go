package main

import (
	"fmt"
)

func main() {
	toplamPort := 1000
	saniyeBasina := 40.0

	var z float64 = float64(toplamPort)

	var result float64 = z / saniyeBasina

	fmt.Printf("Result is = %2.2f", result)
}
