package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Golang Program to Jira Issue Creation - Dhyanio

type zero struct {
	Fields first `json:"fields"`
}
type first struct {
	Project     second `json:"project"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Issuetype   third  `json:"issuetype"`
}
type second struct {
	Key string `json:"key"`
}
type third struct {
	Name string `json:"name"`
}

func main() {
	var jsonData zero
	jsonData = zero{
		Fields: first{
			Project: second{
				Key: "DEV",
			},
			Summary:     "Issue from Golang Program",
			Description: "This issue is created from golang Jira rest API Client with basic-auth using Jira secret token",
			Issuetype: third{
				Name: "Bug",
			},
		},
	}
	mainJSON, err := json.MarshalIndent(jsonData, "", " ")
	if err != nil {
		fmt.Printf("error is: %s", err)
	}
	fmt.Println(string(mainJSON))

	const (
		passWD   = "token"
		baseURL  = "baseurl"
		userNAME = "username"
	)
	request, _ := http.NewRequest("POST", baseURL, bytes.NewBuffer(mainJSON))
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(userNAME, passWD)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The Http request with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
