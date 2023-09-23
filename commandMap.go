package main

import (
	"errors"
	"fmt"
)

var url string = "https://pokeapi.co/api/v2/location-area"

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

	r, err := getJSONData(url)
	url = r.Next
	if err == nil {
		for _, v := range r.Results {
			fmt.Println(v.Name)
		}
		return nil
	} else {
		return err
	}
}

func CommandMapB() error {
	r, err := getJSONData(url)
	if err == nil {
		str, ok := r.Previous.(string)
		if ok {
			if str != "" {
				url = str
				fmt.Println(url, str)
				CommandMap()
			}
		} else {
			return errors.New("Wrong data found")
		}

	}
	return err

}
