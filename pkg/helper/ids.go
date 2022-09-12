package helper

import "fmt"

func MakeIdGenerator(prefix string) func() string {
	id := 0
	return func() string {
		id++
		return fmt.Sprintf("%s-%d", prefix, id)
	}
}
