package main

import "fmt"

type Profile struct {
	Height     float64
	Ativo      bool
	Profession string
}

func MapStruct() {
	var profiles map[string]Profile = make(map[string]Profile)

	profiles["João"] = Profile{
		1.74, true, "Médico",
	}

	profiles["Maria"] = Profile{
		1.63, false, "Engenheira",
	}

	fmt.Println("Perfis salvos:", profiles)
}
