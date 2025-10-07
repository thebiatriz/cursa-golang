package main

import "fmt"

func ModifySlice() {
	var names = [3]string{
		"Ana",
		"Jose",
		"Maria",
	}
	
	var slice1 []string = names[0:2]
	fmt.Println("Slice do indíce 0 ao 1:", slice1)

	slice1[0] = "Rogério"
	fmt.Println("Slice do indíce 0 ao 1 substituido:", slice1)

	fmt.Println("Array original (resultado):", names)

	slice2 := names[0:3]
	slice2[2] = "Beatriz" 
	fmt.Println("Slice do indíce 2 substituído:", slice2)
}
