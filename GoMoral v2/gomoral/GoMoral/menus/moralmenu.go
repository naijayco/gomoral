package menus

import (
	"fmt"
	"gomoral/morals"
	"os"
	"strings"
)

func MoralMenu() {
	for {
		fmt.Println("\nMorals")
		fmt.Println("1. View Morals")
		fmt.Println("2. Add/Remove Morals")
		fmt.Println("3. Back")

		var choice int
		fmt.Println("Choose an option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			morals.ViewMorals()
		case 2:
			MoralMenu_AddRemoveMoral()

			var choicetwo string

			fmt.Println("Return? (Y/N): ")
			fmt.Scanln(&choicetwo)
			answer := strings.ToLower(choicetwo)

			if answer == "yes" {
				break
			} else if answer == "no" {
				fmt.Println("Goodbye!")
				os.Exit(0)
			} else {
				fmt.Println("Invalid input. Please type 'yes' or 'no'")
			}
		case 3:
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func MoralMenu_AddRemoveMoral() {
	for {
		fmt.Println("1. Add Morals")
		fmt.Println("2. Remove Morals")
		fmt.Println("3. Back")

		var choice int

		fmt.Println("Choose an option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\nCreate Moral")
			morals.CreateMoral()
		case 2:
			morals.RemoveMoral()
		case 3:
			return
		default:
			fmt.Println("Invalid option, please try again.")

		}
	}

}
