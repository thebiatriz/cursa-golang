package main

import "fmt"

func LiteralMapStruct() {
	var profiles map[string]Profile = map[string]Profile{
		"João": {
			1.74, true, "Médico",
		},
		"Maria": {
		1.63, false, "Engenheira",
		},
	}

	fmt.Println("Perfis salvos:", profiles)
}
