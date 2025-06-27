package Morals

import (
	"encoding/csv"
	"fmt"
	"gomoral/menus"
	"os"
	"strconv"
	"strings"
)

type Moral struct {
	Name   string
	Weight int
}

var moralsList []Moral // to store morals

// to create a new moral
func CreateMoral() {
	var m Moral

	// Step 1: Get name or cancel
	fmt.Print("Enter the name of the moral (or type 'cancel' to return): ")
	var nameInput string
	fmt.Scanln(&nameInput)
	if strings.ToLower(nameInput) == "cancel" {
		fmt.Println("Addition cancelled.")
		return
	}
	m.Name = nameInput

	// Step 2: Get priority or cancel
	n := len(moralsList)
	fmt.Printf("Enter the priority number for the moral (1 to %d, or type 'cancel' to return): ", n+1)
	var priorityInput string
	fmt.Scanln(&priorityInput)
	if strings.ToLower(priorityInput) == "cancel" {
		fmt.Println("Addition cancelled.")
		return
	}
	priority, err := strconv.Atoi(priorityInput)
	if err != nil || priority < 1 || priority > n+1 {
		fmt.Println("Invalid priority number.")
		return
	}

	// Shift order of priority of existing morals
	for i := range moralsList {
		if int(moralsList[i].Weight) >= priority {
			moralsList[i].Weight++
		}
	}

	m.Weight = priority
	moralsList = append(moralsList, m)
	saveMoralsToCSV()
	fmt.Println("Moral added!")

	for {
		fmt.Print("Add another? Type 'yes' or 'no': ")
		var again string
		fmt.Scanln(&again)
		answer := strings.ToLower(again)
		if answer == "yes" {
			CreateMoral()
			break
		} else if answer == "no" {
			menus.MoralMenu()
		} else {
			fmt.Println("Invalid input. Please type 'yes' or 'no'.")
		}
	}
}

func saveMoralsToCSV() {
	file, err := os.Create("morals.csv") // This will overwrite the file each time
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"Name", "Weight"})

	// Write each moral
	for _, m := range moralsList {
		writer.Write([]string{m.Name, strconv.Itoa(m.Weight)})
	}
}

func LoadMoralsFromCSV() {
	file, err := os.Open("morals.csv")
	if err != nil {
		// If the file doesn't exist, that's fine—just start with an empty list
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Skip header row
	moralsList = []Moral{}
	for i, record := range records {
		if i == 0 {
			continue // skip header
		}
		if len(record) < 2 {
			continue
		}
		weight, err := strconv.Atoi(record[1])
		if err != nil {
			continue
		}
		moralsList = append(moralsList, Moral{
			Name:   record[0],
			Weight: weight,
		})
	}
}

func RemoveMoral() {
	if len(moralsList) == 0 {
		fmt.Println("No morals to remove.")
		return
	}

	// Show current morals and priorities
	ViewMorals()

	fmt.Printf("Enter the priority number of the moral to remove (1 to %d, or type 'cancel' to return): ", len(moralsList))
	var input string
	fmt.Scanln(&input)
	if strings.ToLower(input) == "cancel" {
		fmt.Println("Removal cancelled.")
		return
	}
	removePriority, err := strconv.Atoi(input)
	if err != nil || removePriority < 1 || removePriority > len(moralsList) {
		fmt.Println("Invalid priority number.")
		return
	}

	// Remove the moral with the given priority
	index := -1
	for i, m := range moralsList {
		if m.Weight == removePriority {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Moral not found.")
		return
	}
	moralsList = append(moralsList[:index], moralsList[index+1:]...)

	// Adjust priorities: decrease by 1 for all morals with priority > removed one
	for i := range moralsList {
		if moralsList[i].Weight > removePriority {
			moralsList[i].Weight--
		}
	}

	saveMoralsToCSV()
	fmt.Println("Moral removed!")

	for {
		fmt.Print("Remove another? Type 'yes' or 'no': ")
		var again string
		fmt.Scanln(&again)
		answer := strings.ToLower(again)
		if answer == "yes" {
			RemoveMoral()
			break
		} else if answer == "no" {
			menus.MoralMenu()
		} else {
			fmt.Println("Invalid input. Please type 'yes' or 'no'.")
		}
	}
}
