package menus

import (
	"fmt"
)

func MainMenu() {
	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Morals")
		fmt.Println("2. Exit")

		var choice int
		fmt.Println("Choose an option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			MoralMenu()
		case 2:
			fmt.Println("Goodbye!")
			return
		}
	}
}
