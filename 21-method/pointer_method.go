package main

import (
	"fmt"
	"strings"
)

// adicionar ponteiro para modificar o original, ao invés da cópia
func (p *Person) CapitalizeName() {
	p.Name = strings.ToUpper(p.Name)
}

func PointerMethod() {
	ana := Person{"ana", "lívia"}

	ana.CapitalizeName()
	fmt.Println(ana.FullName())
}
