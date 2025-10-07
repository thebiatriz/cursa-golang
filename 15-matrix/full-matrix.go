package main

import (
	"fmt"
	"strings"
)

func FullMatrix() {
	board := [][]string{
		{"X", "-", "-"},
		{"O", "X", "O"},
		{"-", "O", "X"},
	}

	fmt.Println("\nMatrix completa com Println:")
	for i := range board {
		fmt.Println(board[i])
	}

	fmt.Println("\nMatrix completa com JOIN:")
	for i := range board {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

//  [X - -]
// 	[O X O]
// 	[- O X]
