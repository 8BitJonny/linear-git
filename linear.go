package main

import (
	"context"
	"github.com/machinebox/graphql"
)

type Issue struct {
	BranchName string
}
type Resp struct {
	Issue Issue
}

func GetBranchName(issueId string, token string) (string, error) {
	client := graphql.NewClient("https://api.linear.app/graphql")

	// make a request
	req := graphql.NewRequest(`
		query ($issueId: String!) {
		  issue(id: $issueId) {
			branchName
		  }
		}
	`)

	// set any variables
	req.Var("issueId", issueId)

	req.Header.Set("Authorization", "Bearer "+token)

	var respData Resp
	if err := client.Run(context.Background(), req, &respData); err != nil {
		return "", err
	}
	return respData.Issue.BranchName, nil
}
