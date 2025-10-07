package main

import (
	"fmt"
)

func GoingThroughMatrix() {
	board := [][]string{
		{"X", "-", "-"},
		{"O", "X", "O"},
		{"-", "O", "X"},
	}

	fmt.Println("\nMatrix percorrida individualmente:")
	for i := range board {
		for j := range board[i] {
			fmt.Print(" ", board[i][j])
		}
		fmt.Println()
	}
}

//  [X - -]
// 	[O X O]
// 	[- O X]
