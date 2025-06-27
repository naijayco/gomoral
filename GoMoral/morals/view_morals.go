package Morals

import "fmt"

func ViewMorals() {
	const cols = 4
	const rows = 4
	const perTable = cols * rows

	total := len(moralsList)
	if total == 0 {
		fmt.Println("No morals to display.")
		return
	}

	for tableStart := 0; tableStart < total; tableStart += perTable {
		tableEnd := tableStart + perTable
		if tableEnd > total {
			tableEnd = total
		}
		fmt.Println("\nMorals Table:")
		// Print rows
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				idx := tableStart + c*rows + r
				if idx < tableEnd {
					m := moralsList[idx]
					fmt.Printf("%2d. %-15s", m.Weight, m.Name)
				}
			}
			fmt.Println()
		}
	}
}
