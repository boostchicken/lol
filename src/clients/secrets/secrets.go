package secrets // import "github.com/boostchicken/lol/clients/secrets"

import (
	"context"
	"log"

	sm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

var secretsClient *sm.Client

func init() {
	makeClient()
}

var region string = "us-east-2"
var endpoint string = "https://localhost.localstack.cloud:4566"

func makeClient() error {
	r := "us-east-2"
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(getAWSLocalstackConfig(region)), config.WithRegion(r), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "test")))
	if err != nil {
		log.Fatal(err)
	}

	secretsClient = sm.NewFromConfig(cfg)
	return nil
}

func WithRegion(r string) error {
	region = r
	return makeClient()
}

func WithEndpoint(e string) error {
	endpoint = e
	return makeClient()
}

func getAWSLocalstackConfig(region string) aws.EndpointResolverWithOptionsFunc {
	return aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: endpoint}, nil
	})
}

func getSecretString(name string) (res *sm.GetSecretValueOutput, err error) {
	return secretsClient.GetSecretValue(context.TODO(), &sm.GetSecretValueInput{SecretId: aws.String(name)})
}

func GetDSN() (secret *string, err error) {
	res, err := getSecretString("boost-lol-dev")
	if err != nil {
		return nil, err
	}
	return res.SecretString, err
}
