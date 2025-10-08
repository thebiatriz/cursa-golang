package main

import "fmt"

func LiteralMap() {
	var grades map[string]int = map[string]int{
		"Ana":   9,
		"Maria": 10,
	}

	fmt.Println(grades)
}
