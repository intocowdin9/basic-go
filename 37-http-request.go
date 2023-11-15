package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// start 36-web-api-json.go to activated server
var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

// fetch all users
func fetchUsers() []*student {
	var err error
	var client = &http.Client{}
	var data []*student

	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	return data
}

// fetch user by ID
func fetchUser(ID string) student {
	var err error
	var client = &http.Client{}
	var data student

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}

	return data
}

func main() {
	var users = fetchUsers()
	var user1 = fetchUser("E001")

	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}

	// fetch user1
	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}
