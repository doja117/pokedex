package main

import (
	"errors"
	"fmt"
)

func getJSONData(url string) (ResponseData, error) {
	r, err := ReturnJSONData(url)
	if err != nil {
		return ResponseData{}, errors.New(fmt.Sprintf("Error occured ", err))
	} else {

		for _, v := range r.Results {
			fmt.Println(v.Name)
		}
		return r, nil
	}
}

func CommandMap() error {
	//fmt.Println("function called")
	url := "https://pokeapi.co/api/v2/location-area"
	r, err := getJSONData(url)
	if err == nil {
		for _, v := range r.Results {
			fmt.Println(v.Name)
		}
		return nil
	} else {
		return err
	}
}
