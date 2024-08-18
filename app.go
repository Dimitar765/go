package main

import (
	"fmt"
)

func main() {
	choise := 0
	balance := 2000

	fmt.Println("Hello, succer!")
	for {

		fmt.Println("1. check your balance")
		fmt.Println("2. widraw")
		fmt.Println("3. choise No3")
		fmt.Println("4. exit")
		fmt.Scan(&choise)
		fmt.Println("You choise ", &choise)

		if choise == 1 {
			fmt.Println("You balance is:", balance)
			continue
		} else if choise == 2 {
			withraw := 0
			fmt.Println("How much do you want to withraw?")
			fmt.Scan(&withraw)
			if withraw > balance {
				fmt.Println("You don't have enough money")
				continue
			}
			balance = balance - withraw
			fmt.Println("you widraw", withraw, "your balance is", balance)
		} else if choise == 3 {
			fmt.Println("You choise No3")
		} else {
			fmt.Println("thank you")
			return
		}
	}
}
