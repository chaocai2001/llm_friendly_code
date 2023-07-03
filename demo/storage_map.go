package demo

import (
	"errors"
)

type MapStorage struct {
	data map[string]string
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		data: make(map[string]string),
	}
}

func (ms *MapStorage) RetiveData(token string) (string, error) {
	value, ok := ms.data[token]
	if !ok {
		return "", errors.New("token not found")
	}
	return value, nil
}

func (ms *MapStorage) StoreData(token string, data string) error {
	ms.data[token] = data
	return nil
}
