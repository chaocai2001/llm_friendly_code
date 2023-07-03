package demo

import (
	"fmt"

	"github.com/google/uuid"
)

type UUIDTokenCreator struct{}

func (utc *UUIDTokenCreator) CreateToken(data string) string {
	token := uuid.New()
	return fmt.Sprintf("%s-%s", token, data)
}

func NewUUIDTokenCreator() TokenCreator {
	return &UUIDTokenCreator{}
}
