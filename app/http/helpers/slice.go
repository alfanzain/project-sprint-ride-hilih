package helpers

type Slice[T comparable] []T

func (s Slice[T]) Includes(item T) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}
