package main

import "fmt"

func main() {
	numbers := []float64{12.0, 3.3, 2.9, 4.5}
	total := sum(numbers...)
	fmt.Printf("%f\n", total)
}

func sum(values ...float64) float64 {
	total := 0.0

	for _, num := range values {
		total += num
	}

	return total
}