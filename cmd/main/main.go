package main

import (
	"fmt"
	"project/internal/ch1"
)

var optionsChapter1 = map[string]func(){
	"echo1":    ch1.Echo1,
	"echo2":    ch1.Echo2,
	"echo3":    ch1.Echo3,
	"dup1":     ch1.Dup1,
	"dup3":     ch1.Dup3,
	"map1":     ch1.Map1,
	"fetch":    ch1.Fecth,
	"fetchall": ch1.Fetchall,
	"server":   ch1.Server,
}

func main() {
	fmt.Println("=== Welcome! ===")
	printMenu()
}

func printMenu() {
	var choice int

	// TODO: Print always in the same order
	for {
		fmt.Println("Chapter 1: ")
		mappedOptions := make(map[int]string)
		optionCounter := 1

		for key := range optionsChapter1 {
			fmt.Printf("    %d. %s\n", optionCounter, key)
			mappedOptions[optionCounter] = key
			optionCounter++
		}
		fmt.Print("> ")

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Scanln() // clear the invalid input
			fmt.Println("Invalid input, please enter a number:")
			continue
		}

		if choice <= 0 || choice > optionCounter-1 {
			fmt.Println("Out of range, choose again:")
		} else {
			optionsChapter1[mappedOptions[choice]]()
			break
		}
	}
}
