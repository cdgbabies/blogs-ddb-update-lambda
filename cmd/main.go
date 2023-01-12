package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cdgbabies/blogs-ddb-update-lambda/service"
	"github.com/google/go-github/v49/github"
)

type handler struct {
	actionService         service.GitHubWorkflowClient
	repoName              string
	ownerName             string
	workflowFileName      string
	workflowDispatchEvent github.CreateWorkflowDispatchEventRequest
}

func (h *handler) handleRequest(ctx context.Context, e events.DynamoDBEvent) error {

	_, err := h.actionService.CreateWorkflowDispatchEventByFileName(ctx, h.ownerName, h.repoName, h.workflowFileName, h.workflowDispatchEvent)
	return err
}

func main() {
	h := handler{
		repoName:         os.Getenv("repo_name"),
		ownerName:        os.Getenv("owner_name"),
		workflowFileName: os.Getenv("workflow_file_name"),
		workflowDispatchEvent: github.CreateWorkflowDispatchEventRequest{
			Ref: os.Getenv("branch"),
		},

		actionService: service.GetGitHubClient(context.Background()),
	}
	lambda.Start(h.handleRequest)
}
