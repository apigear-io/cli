package helper

import "fmt"

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
