package main

import (
	"fmt"
)

func main() {
	choise := 0
	fmt.Println("Hello, World!")
	fmt.Println("1. choise No1")
	fmt.Println("2. choise No2")
	fmt.Println("3. choise No3")
	fmt.Scan(&choise)
	fmt.Println("You choise ", &choise)
}
