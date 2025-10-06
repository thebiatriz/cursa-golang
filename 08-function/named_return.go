package main

func CalcNamedReturn(value int) (squared, cubed int) {
	squared = value * value
	cubed = value * value * value

	//naked return (evitar com funções longas)
	return
}