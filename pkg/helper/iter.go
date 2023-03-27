package helper

type Iterator[T any] interface {
	Next() (T, bool)
	HasNext() bool
}

type iterator[T any] struct {
	n     int
	items []T
}

func NewIterator[T any](items []T) Iterator[T] {
	return &iterator[T]{
		n:     0,
		items: items,
	}
}

func (i *iterator[T]) Next() (T, bool) {
	if i.n >= len(i.items) {
		var item T
		return item, false
	}
	item := i.items[i.n]
	i.n++
	return item, true
}

func (i *iterator[T]) HasNext() bool {
	return i.n < len(i.items)
}
