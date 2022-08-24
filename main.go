package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	cons "<project>/pkg/constant"
	"net/http"
)

type createIssue struct {
	Fields fields `json:"fields"`
}
type fields struct {
	Project     key    `json:"project"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Issuetype   name   `json:"issuetype"`
}

type id struct {
	Id string `json:"id"`
}
type key struct {
	Key string `json:"key"`
}
type name struct {
	Name string `json:"name"`
}

type updateIssue struct {
	Update update `json:"update"`
}

type assignee struct {
	Set id `json:"set"`
}
type update struct {
	Assignee []assignee `json:"assignee"`
}

func CreateIssue(summary, description, issuetype string) ([]byte, error) {
	var jsonData = createIssue{
		Fields: fields{
			Project: key{
				Key: cons.JiraProjectId,
			},
			Summary:     summary,
			Description: description,
			Issuetype: name{
				Name: issuetype,
			},
		},
	}
	mainJSON, err := json.MarshalIndent(jsonData, "", " ")
	if err != nil {
		return nil, err
	}
	res, err := postJira("POST", cons.JiraBaseUrl, mainJSON)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateIssue(ticketId string) ([]byte, error) {
	var updateIssueJson = updateIssue{
		Update: update{
			Assignee: []assignee{
				{
					Set: id{
						Id: cons.JiraAssignee,
					},
				},
			},
		},
	}
	mainJSON, errJson := json.MarshalIndent(updateIssueJson, "", " ")
	if errJson != nil {
		return nil, errJson
	}
	url := cons.JiraBaseUrl + "/" + ticketId
	_, err := postJira("PUT", url, mainJSON)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func postJira(method, url string, postBody []byte) ([]byte, error) {
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(postBody)) //url: https://<yourjira>/rest/api/2/issue
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(cons.JiraUsername, cons.JiraToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The Http request with error %s\n", err)
		return nil, err
	}
	data, _ := ioutil.ReadAll(response.Body)

	return data, nil
}
