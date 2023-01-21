package service

import (
	"context"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestUploadFileContents(t *testing.T) {
	// Create a mock implementation of the S3UploadObjectClient interface
	mockS3Client := &mockS3UploadObjectClient{}

	// Test the function with the mock client

	uploadContentsRequest := UploadContentsRequest{
		BucketName: "mybucket",
		Key:        "myfile.txt",
		Contents:   "Hello, world!",
	}
	err := UploadFileContents(context.Background(), mockS3Client, uploadContentsRequest)
	assert.Nil(t, err)

	// Check that the mock client received the expected input
	assert.Equal(t, *mockS3Client.putObjectInput.Bucket, uploadContentsRequest.BucketName)
	assert.Equal(t, *mockS3Client.putObjectInput.Key, uploadContentsRequest.Key)
	data, _ := ioutil.ReadAll(mockS3Client.putObjectInput.Body.(*strings.Reader))
	assert.Equal(t, string(data), uploadContentsRequest.Contents)
}

// mockS3UploadObjectClient is a mock implementation of the S3UploadObjectClient interface
type mockS3UploadObjectClient struct {
	putObjectInput *s3.PutObjectInput
}

func (m *mockS3UploadObjectClient) PutObject(ctx context.Context, input *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	// Save the input for later examination in the test
	m.putObjectInput = input
	return &s3.PutObjectOutput{}, nil
}
