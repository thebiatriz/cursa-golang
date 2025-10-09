package main

import "fmt"

type Integer int

func (i Integer) Quad() Integer {
	return i * i
}

type Other string

func (o Other) Quad() Integer {
	return Integer(len(o) * len(o))
}

type Power interface {
	Quad() Integer
}

func describe(power Power) {
	fmt.Printf("valor: %v / tipo: %T\n", power, power)
}

func InterfaceValues() {
	var pot Power = Integer(3) //valor(3) e tipo(Integer)

	var out Power = Other("abcd")

	fmt.Println("Valor da potência:", pot.Quad())
	describe(pot)

	fmt.Println("String armazenada:", out.Quad())
	describe(out)
}
