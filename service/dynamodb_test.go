package service

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// mockDynamoDbReadOperationClient is a mock implementation of the DynamoDbReadOperationClient interface
type mockDynamoDbReadOperationClient struct{}

func (m *mockDynamoDbReadOperationClient) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	// Return a mocked response
	response := &dynamodb.QueryOutput{
		Items: []map[string]types.AttributeValue{
			{

				"sk":          &types.AttributeValueMemberS{Value: "123"},
				"pk":          &types.AttributeValueMemberS{Value: "blogs"},
				"description": &types.AttributeValueMemberS{Value: "Description"},
				"title":       &types.AttributeValueMemberS{Value: "Test Title"},
				"createdDate": &types.AttributeValueMemberS{Value: "2022-01-01T00:00:00Z"},
			},
		},
	}

	return response, nil
}

func TestQueryDynamoDB(t *testing.T) {
	ctx := context.Background()
	client := &mockDynamoDbReadOperationClient{}
	tableName := "testTable"

	blogs, err := QueryDynamoDB(ctx, client, tableName)

	if err != nil {
		t.Errorf("QueryDynamoDB returned an error: %v", err)
	}

	if len(blogs) != 1 {
		t.Errorf("Expected 1 blog, got %v", len(blogs))
	}
}
