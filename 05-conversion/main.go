package main

import "fmt"

func main() {
	var value1 int = 46
	var value2 float64 = 6.4

	// dessa forma não funciona
	// var value3 float64 = value1

	// modo correto de converter
	var value3 float64 = float64(value1)
	var value4 int = int(value2)

	fmt.Printf("Primeiro valor convertido para float: %.2f\n", value3)
	fmt.Println("Segundo valor convertido para int:", value4)
}