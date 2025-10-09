package main

import "fmt"

type Person struct {
	Name    string
	Surname string
}

// método é uma função que está associada a um tipo específico
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

func FullNameFunction(p Person) string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

func CreateMethod() {
	someone := Person{"José", "de Alencar"}

	maria := Person{"Maria", "Almeida"}

	fmt.Println("Pessoa cadastrada:", someone.FullName())
	fmt.Println("Pessoa cadastrada (sem receptor):", FullNameFunction(maria))

}
