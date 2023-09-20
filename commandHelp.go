package main

import "fmt"

func CommandHelp() {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your avialable commands")
	comamnds := getCommands()
	for key, _ := range comamnds {
		fmt.Println(" - ", key)
		fmt.Println(comamnds[key].description)
	}
	fmt.Println("")
}
