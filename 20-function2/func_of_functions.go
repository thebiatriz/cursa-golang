package main

import "fmt"

// João - Olá João
func greet(name string) func() string {
	return func() string {
		return fmt.Sprintf("Olá, %s!", name)
	}

}

func FuncOfFunctions() {
	greetJoao := greet("João")
	fmt.Println(greetJoao())

	greetMaria := greet(("Maria"))
	fmt.Println(greetMaria())
}
