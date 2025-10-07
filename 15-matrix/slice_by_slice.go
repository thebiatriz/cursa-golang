package main

import (
	"fmt"
)

func SliceBySlice() {
	board := [][]string{
		{"X", "-", "-"},
		{"O", "X", "O"},
		{"-", "O", "X"},
	}

	fmt.Println("Primeira linha da matrix tabuleiro:", board[0])
}

//  [X - -]
// 	[O X O]
// 	[- O X]
