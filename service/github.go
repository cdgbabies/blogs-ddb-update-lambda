package service

import (
	"context"

	"github.com/google/go-github/v49/github"
	"golang.org/x/oauth2"
)

type GitHubWorkflowClient interface {
	CreateWorkflowDispatchEventByFileName(ctx context.Context, owner, repo, workflowFileName string, event github.CreateWorkflowDispatchEventRequest) (*github.Response, error)
}

func GetGitHubClient(ctx context.Context) *github.ActionsService {
	token, _ := GetAccessTokenforGithub(ctx)

	if *token == "" {
		return github.NewClient(nil).Actions
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc).Actions
}
