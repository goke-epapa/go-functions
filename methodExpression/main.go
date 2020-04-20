package main

import (
	"fmt"
	"go-functions/simplemath"
	"math"
)

// MathExpr is an expression type
type MathExpr = string

const (
	// AddExpr ...
	AddExpr = MathExpr("add")

	// SubtractExpr ...
	SubtractExpr = MathExpr("subtract")

	// MultiplyExpr ...
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	Expr := mathExpression(MultiplyExpr)
	println(Expr(2.0, 3.9))

	fmt.Printf("%f\n", double(3, 2, mathExpression(MultiplyExpr)))

	fmt.Println("----------------")

	square := powerOfTwo()

	fmt.Println(square())
	fmt.Println(square())
	fmt.Println(square())
	fmt.Println(square())
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	default:
		return func(f1 float64, f2 float64) float64 {
			return 0
		}
	}
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}

func powerOfTwo() func() int64 {
	x := 1.0

	return func() int64 {
		x++
		return int64(math.Pow(x, 2))
	}
}
