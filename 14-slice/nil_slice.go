package main

import "fmt"

func NilSlice() {
	var number []int

	fmt.Println("Valor zero do pedaço é nil? -", number == nil)
}