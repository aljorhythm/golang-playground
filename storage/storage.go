package storage

type Store struct {
}

type Storage interface {
	Store(id string, data []byte) error
	Retrieve(id string) ([]byte, error)
}
