package main

import "math"

//mais de um valor de retorno
func Calc(value int) (int, float64) {
	var squared int = value * value
	var cubed = math.Pow(float64(value), 3.0)

	return squared, cubed
}