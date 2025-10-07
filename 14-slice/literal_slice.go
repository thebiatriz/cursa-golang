package main

import "fmt"

func LiteralSlice() {
	var names = []string{
		"Ana",
		"Jose",
		"Maria",
	}

	names = names[:2]

	fmt.Println(names)
	fmt.Printf("len = %d, cap = %d\n", len(names), cap(names))
	fmt.Printf("Tipo do slice: %T\n", names)
}
