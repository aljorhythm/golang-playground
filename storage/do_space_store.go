package storage

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"log"
)

type SpaceStore struct {
	*minio.Client
	BucketProperties
}

func (s *SpaceStore) Store(ctx context.Context, id string, data []byte) error {
	reader := bytes.NewReader(data)
	uploadInfo, err := s.Client.PutObject(ctx, s.Name, id, reader, int64(len(data)), minio.PutObjectOptions{})

	if err != nil {
		log.Panicf("bucket %s id %s upload info %#v error %s", s.Name, id, uploadInfo, err)
		return err
	}

	log.Printf("bucket %s id %s updload info %#v", s.Name, id, uploadInfo)
	return nil
}

func (s *SpaceStore) Retrieve(ctx context.Context, id string) ([]byte, error) {
	panic("implement me")
}

type BucketProperties struct {
	Name     string
	Location string
}

func ListSomeObjects(ctx context.Context, endpoint string, minioOpts *minio.Options) ([]minio.ObjectInfo, error) {
	client, err := minio.New(endpoint, minioOpts)

	if err != nil {
		log.Panicf("error connecting to space endpoint %s", endpoint)
		return nil, ERROR_FAIL_TO_CONNECT_SPACE_STORE
	}

	c := client.ListObjects(ctx, "", minio.ListObjectsOptions{
		WithVersions: false,
		WithMetadata: false,
		Prefix:       "",
		Recursive:    false,
		MaxKeys:      0,
		StartAfter:   "",
		UseV1:        false,
	})

	list := []minio.ObjectInfo{}

	for item, done := <-c; done ; {
	list := append(list, item)
	   if len(list) > 10 {
		   break
	   }
	}

	if err != nil {
		log.Panicf("error listing objects %#v", err)
	} else {
		log.Printf("objects %s", )
	}
	return list, err
}

func NewSpaceStore(ctx context.Context, endpoint string, minioOpts *minio.Options, bucket BucketProperties) (Storage, error) {
	client, err := minio.New(endpoint, minioOpts)
	if err != nil {
		return nil, ERROR_FAIL_TO_CONNECT_SPACE_STORE
	}

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
