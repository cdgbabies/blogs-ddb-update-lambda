package service

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSMGetParameterAPI interface {
	GetParameter(ctx context.Context,
		params *ssm.GetParameterInput,
		optFns ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
}

func GetSSMClient(ctx context.Context) *ssm.Client {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ssm.NewFromConfig(cfg)
	return client

}
func FindParameter(c context.Context, api SSMGetParameterAPI, input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	return api.GetParameter(c, input)

}
func GetAccessTokenforGithub(ctx context.Context) (*string, error) {
	client := GetSSMClient(ctx)
	tokenParamName := os.Getenv("token_name")
	decrypted := true
	input := &ssm.GetParameterInput{
		Name:           &tokenParamName,
		WithDecryption: &decrypted,
	}

	results, err := FindParameter(ctx, client, input)
	if err != nil {
		panic(err)
	}

	return results.Parameter.Value, err
}
