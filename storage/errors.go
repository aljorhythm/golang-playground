package storage

import "errors"

var (
	ERROR_DATA_NOT_FOUND = errors.New("ERROR_DATA_NOT_FOUND")
	ERROR_STORE_DATA_FAILED = errors.New("ERROR_STORE_DATA_FAILED")
)