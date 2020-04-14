package main

import (
	"go-functions/simplemath"
)

func main() {
	sv := simplemath.NewSemanticVersion(1, 2, 0)
	sv.IncrementMajor()
	sv.IncrementMinor()
	println(sv.String())
}
