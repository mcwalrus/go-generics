package generics

func Apply[T any](items []T, fn func(item T) T) []T {
	s := make([]T, len(items))
	for i, v := range items {
		s[i] = fn(v)
	}
	return s
}

func Convert[T, R any](items []T, fn func(item T) R) []R {
	s := make([]R, len(items))
	for i, v := range items {
		s[i] = fn(v)
	}
	return s
}

func Filter[T any](items []T, fn func(item T) bool) []T {
	s := make([]T, 0, len(items))
	for _, item := range items {
		if fn(item) {
			s = append(s, item)
		}
	}
	return s
}

func Reduce[T any](items []T, fn func(a, b T) T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	var val = items[0]
	for i := 1; i < len(items); i++ {
		val = fn(val, items[i])
	}
	return val
}
