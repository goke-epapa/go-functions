package main

import (
	"fmt"
	"go-functions/simplemath"
)

func main() {
	answer, err := simplemath.Divide(6, 2)

	if err != nil {
		fmt.Printf("an error occured %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}
}
