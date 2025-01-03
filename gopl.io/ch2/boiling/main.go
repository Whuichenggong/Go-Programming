package main

import (
	"fmt"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c float64
	c = (f - 32) * 5 / 9
	fmt.Printf("%g°F = %g°C\n", f, c)
}
