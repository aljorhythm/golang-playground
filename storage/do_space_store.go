package storage

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type SpaceStore struct {
	*s3.S3
	BucketProperties
}

func (s *SpaceStore) Store(ctx context.Context, id string, data []byte) error {
	return errors.New("unimplemented")
}

func (s *SpaceStore) Retrieve(ctx context.Context, id string) ([]byte, error) {
	panic("implement me")
}

type BucketProperties struct {
	Name     string
	Location string
}

/**
https://docs.digitalocean.com/products/spaces/resources/s3-sdk-examples/
aws region is us-east-1 for digital ocean spaces
 */
func getRegion() *string {
	return aws.String("us-east-1")
}

func NewSpaceStore(ctx context.Context, endpoint string, key string, secret string, bucket BucketProperties) (*SpaceStore, error) {
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String(endpoint),
		Region:      getRegion(),
	}

	newSession, err := session.NewSession(s3Config)

	if err != nil {
		return nil, err
	}

	client := s3.New(newSession)

	return &SpaceStore{
		client,
		bucket,
	}, nil
}

type DigitalOceanSpaceConfig struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Endpoint  string `yaml:"endpoint"`
}
