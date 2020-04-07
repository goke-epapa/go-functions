package simplemath

import (
	"errors"
)

// Divide performs division operation
func Divide(p1, p2 float64) (answer float64, err error) {
	if p2 == 0 {
		err = errors.New("cannot divide by zero")
	}
	answer = p1 / p2

	return
}

// Add performs add operation
func Add(p1, p2 float64) float64 {
	return p1 + p2
}

// Subtract performs subtract operation
func Subtract(p1, p2 float64) float64 {
	return p1 - p2
}

// Multiply performs multiply operation
func Multiply(p1, p2 float64) float64 {
	return p1 * p2
}
