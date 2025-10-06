package main

import (
	"fmt"
	"time"
)

func main() {
	var add int = 2

	for {
		add += add
		fmt.Println(add)
		time.Sleep((100 * time.Millisecond))
	}
}
