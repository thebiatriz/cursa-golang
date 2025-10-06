package main

import "fmt"

func main() {
	// usa-se o for também para while
	// somar potencias de 2

	var add int = 2

	for add < 600 {
		add += add
	}

	fmt.Println(add)
}
