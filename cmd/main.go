package main

import (
	"fmt"
	"os"
)

func main() {

	tasks := make(map[string]string)

	choice := -1
	for true {
		fmt.Println()
		fmt.Println("What would you like to do?")
		fmt.Println("1. Show tasks")
		fmt.Println("2. Add a task")
		fmt.Println("0. Exit")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Couldn't parse your choice, try again")
			continue
		}

		switch choice {
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		case 1:
			fmt.Println("Tasks")
			for t, _ := range tasks {
				fmt.Println(t)
			}
		case 2:
			fmt.Println("Add a task placeholder")
		default:
			fmt.Println("Invalid choice")
		}
	}

}
