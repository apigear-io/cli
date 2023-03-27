package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func MakeIdGenerator(prefix string) func() string {
	id := 0
	return func() string {
		id++
		return fmt.Sprintf("%s-%d", prefix, id)
	}
}

func MakeIntIdGenerator() func() uint64 {
	var id uint64
	return func() uint64 {
		id++
		return id
	}
}

func NewUUID() string {
	return uuid.NewString()
}
