package storage

type InmemoryStore struct {
	dataMap map[string][]byte
}

func (i *InmemoryStore) Retrieve(id string) ([]byte, error) {
	if data, ok := i.dataMap[id]; ok {
		return data, nil
	} else {
		return nil, ERROR_DATA_NOT_FOUND
	}
}

func (i *InmemoryStore) Store(id string, bytes []byte) error {
	i.dataMap[id] = bytes
	return nil
}

func NewInmemoryStore() Storage {
	return &InmemoryStore{
		map[string][]byte{},
	}
}
