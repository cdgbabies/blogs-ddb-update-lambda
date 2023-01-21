package service

import (
	"context"

	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/aws-sdk-go/aws"
)

type mockSSMClient struct {
	response ssm.GetParameterOutput
	err      error
}

func (m *mockSSMClient) GetParameter(ctx context.Context, input *ssm.GetParameterInput, options ...func(*ssm.Options)) (*ssm.GetParameterOutput, error) {
	return &m.response, m.err
}

func TestGetAccessTokenforGithub(t *testing.T) {
	// Test case 1: Successful retrieval of token
	response := ssm.GetParameterOutput{
		Parameter: &types.Parameter{
			Value: aws.String("abc123"),
		},
	}
	mockClient := &mockSSMClient{response: response}

	resp, err := FindParameter(context.TODO(), mockClient, &ssm.GetParameterInput{Name: aws.String("TestName")})
	if err != nil {
		t.Errorf("Error getting parameter")
	}
	if *resp.Parameter.Value != "abc123" {
		t.Errorf("Expected: Parameter value to be 'abc123', but got '%s'", *resp.Parameter.Value)
	}

}
