package storage

import "context"

type Store struct {
}

type Storage interface {
	Store(ctx context.Context, id string, data []byte) error
	Retrieve(context context.Context, id string) (data []byte, err error)
	Delete(context context.Context, id string) error
}
