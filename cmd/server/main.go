package main

import (
	"SPE_MiniProject/internal/fact"
	"SPE_MiniProject/internal/ln"
	"SPE_MiniProject/internal/pow"
	"SPE_MiniProject/internal/sqrt"
	"fmt"
)

func main() {
	banner := `
  _________      .___________        .__          
 /   _____/ ____ |__\_   ___ \_____  |  |   ____  
 \_____  \_/ ___\|  /    \  \/\__  \ |  | _/ ___\ 
 /        \  \___|  \     \____/ __ \|  |_\  \___ 
/_______  /\___  >__|\______  (____  /____/\___  >
        \/     \/           \/     \/          \/ 
	`
	fmt.Println(banner)

	for {
		fmt.Println("Choose your operation[1, 2, 3, 4]")
		fmt.Println("[1] Square Root of a Number (√x)")
		fmt.Println("[2] Factorial of a Number (x!)")
		fmt.Println("[3] Natural Logarithm of a Number(ln(x))")
		fmt.Println("[4] Power Function (x ^ b)")

		fmt.Println("Enter q to exit")

		var choice string
		fmt.Scan(&choice)

		if choice == "1" {
			result, err := sqrt.Sqrt()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			fmt.Printf("Result: %f\n", result)
		} else if choice == "2" {
			result, err := fact.Fact()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			fmt.Printf("Result: %d\n", result)
		} else if choice == "3" {
			result, err := ln.Ln()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			fmt.Printf("Result: %f\n", result)
		} else if choice == "4" {
			result, err := pow.Pow()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			fmt.Printf("Result: %f\n", result)
		} else if choice == "q" {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
