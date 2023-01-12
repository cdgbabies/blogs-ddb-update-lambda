package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-github/v49/github"
)

type MockGitActionService struct {
	event    events.DynamoDBEvent
	response *github.Response
}

func (service *MockGitActionService) CreateWorkflowDispatchEventByFileName(ctx context.Context, owner, repo, workflowFileName string, event github.CreateWorkflowDispatchEventRequest) (*github.Response, error) {

	return service.response, nil
}

func TestHandleRequest(t *testing.T) {
	h := handler{
		repoName:         "repo_name",
		ownerName:        "owner_name",
		workflowFileName: "workflow_file_name",
		workflowDispatchEvent: github.CreateWorkflowDispatchEventRequest{
			Ref: "branch",
		},
		actionService: &MockGitActionService{
			event:    events.DynamoDBEvent{},
			response: &github.Response{},
		},
	}
	err := h.handleRequest(context.Background(), events.DynamoDBEvent{})
	if err != nil {
		t.Fatalf("Encountered Error : %s", err)
	}

}
