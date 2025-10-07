package main

import "fmt"

func DefineStruct() {

	// exemplo: jogo de duas dimensões. armazenar posição do jogador

	type Position struct {
		X int
		Y int
	}

	pos := Position{40, 67}

	pos.Y = 51

	fmt.Println("Valor do position:", pos)
}
