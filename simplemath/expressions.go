package simplemath

import (
	"errors"
	"math"
)

// Divide performs division operation
func Divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by zero")
	}
	return p1 / p2, nil
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
