package main

import (
	"errors"
	"fmt"
	"net/http"
)

func GetPockemon(name string) {
	res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/", name))
	if err != nil {
		fmt.Println(res)
	} else {
		errors.New("Error Occured ")
		fmt.Printf(err.Error())
	}
}
