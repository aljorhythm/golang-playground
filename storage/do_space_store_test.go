package storage

import (
	"context"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"os"
	"testing"

	"github.com/minio/minio-go/v7"
)

func readTestConfig(t *testing.T) DigitalOceanSpaceConfig {
	dat, err := os.ReadFile(".config.yml")

	if err != nil {
		t.Fatal("unable to read config.yml")
	}

	config := DigitalOceanSpaceConfig{}
	err = yaml.Unmarshal(dat, &config)

	if err != nil {
		t.Fatal("unable to parse config.yml")
	}
	return config
}

func TestDigitalOceanSpaceStorage(t *testing.T) {
	if os.Getenv("EXTERNAL_TEST") == "" {
		t.Skipf("env var EXTERNAL_TEST not set, skipping")
	}

	ctx := context.Background()

	config := readTestConfig(t)
	accessKey := config.AccessKey
	secKey := config.SecretKey
	endpoint := config.Endpoint

	t.Run("list buckets", func(t *testing.T) {
		buckets, err := ListSomeObjects(ctx,
			endpoint,
			&minio.Options{
				Creds:  credentials.NewStaticV4(accessKey, secKey, ""),
				Secure: true,
			})

		assert.NoError(t, err)
		t.Logf("buckets %#v", buckets)
	})

	t.Run("SpaceStore", func(t *testing.T) {
		bucketProps := BucketProperties{
			Name:     "yumseng",
			Location: "us-east-1",
		}
		store, err := NewSpaceStore(ctx, endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKey, secKey, ""),
			Secure: true,
		}, bucketProps)

		assert.NoError(t, err)
		assert.NotNil(t, store)

		objectId := "object-1"
		data := []byte("abcde")
		t.Run("store bytes", func(t *testing.T) {
			err := store.Store(ctx, objectId, data)
			assert.NoError(t, err)
		})
	})
}
