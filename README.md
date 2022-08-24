# go-jira

## Golang jira library!

- Create New Jira issue
- Update existing jira issue

```golang
package main

import (
	jiraUtil "github.com/dhyanio/go-jira"
)

func main() {
  // Create Jira Issue
	jiraRes, err := jiraUtil.CreateIssue(summary, resdes, "Task")
	if err != nil {
		return "", err
	}
  
  // Parse into JSON
	jiraResJson, err := util.ParseEvent(string(jiraRes))
	if err != nil {
		return "", err
	}
  
  // Get Jira TicketId/IssueId
	ticketId := jiraResJson["key"].(string)
	_, updateIssueErr := jiraUtil.UpdateIssue(ticketId)
	if updateIssueErr != nil {
		return "", updateIssueErr
	}
}

  
```
