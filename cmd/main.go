package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cdgbabies/blogs-ddb-update-lambda/service"
	"github.com/google/go-github/v49/github"
)

type handler struct {
	actionService         service.GitHubWorkflowClient
	dynamodbClient        service.DynamoDbReadOperationClient
	s3Client              service.S3UploadObjectClient
	repoName              string
	ownerName             string
	workflowFileName      string
	workflowDispatchEvent github.CreateWorkflowDispatchEventRequest
}

func (h *handler) handleRequest(ctx context.Context, e events.DynamoDBEvent) error {

	blogs, err := service.QueryDynamoDB(ctx, h.dynamodbClient, os.Getenv("table_name"))
	if err != nil {
		return err
	}
	body, err := json.Marshal(blogs)
	if err != nil {
		return err
	}
	err = service.UploadFileContents(ctx, h.s3Client, service.UploadContentsRequest{
		BucketName: os.Getenv("bucket_name"),
		Key:        os.Getenv("file_name"),
		Contents:   string(body),
	})
	if err != nil {
		return err
	}
	_, err = h.actionService.CreateWorkflowDispatchEventByFileName(ctx, h.ownerName, h.repoName, h.workflowFileName, h.workflowDispatchEvent)
	return err
}

func main() {
	region := os.Getenv("region_name")
	h := handler{
		repoName:         os.Getenv("repo_name"),
		ownerName:        os.Getenv("owner_name"),
		workflowFileName: os.Getenv("workflow_file_name"),
		workflowDispatchEvent: github.CreateWorkflowDispatchEventRequest{
			Ref: os.Getenv("branch"),
		},
		dynamodbClient: service.NewDynamoDbClient(context.Background(), region),
		s3Client:       service.NewS3Client(context.Background(), region),

		actionService: service.GetGitHubClient(context.Background()),
	}
	lambda.Start(h.handleRequest)
}
