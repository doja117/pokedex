package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ResponseData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func ReturnJSONData(url string) (ResponseData, error) {

	response, err := http.Get(url)
	if err != nil {
		return ResponseData{}, err

	} else {
		defer response.Body.Close()
		if response.StatusCode != 200 {
			fmt.Println("Status Code showing error .....", response.StatusCode)
			return ResponseData{}, errors.New(fmt.Sprintf("Error Occured", response.StatusCode))
		}
		var data ResponseData
		if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
			return data, errors.New(fmt.Sprintf("error decoding json file with error ....", err))
		}

		return data, nil
	}
}
