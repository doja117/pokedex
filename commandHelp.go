package main

import "fmt"

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your avialable commands")
	comamnds := getCommands()
	for key, _ := range comamnds {
		fmt.Printf(" - %s\t%s\n", key, comamnds[key].description)
	}
	fmt.Println("")
	return nil
}
