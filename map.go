package generics

func MapKeys[K comparable, V any](items map[K]V) []K {
	keys := make([]K, 0, len(items))
	for k := range items {
		keys = append(keys, k)
	}
	return keys
}

func MapValues[K comparable, V any](items map[K]V) []V {
	values := make([]V, 0, len(items))
	for _, v := range items {
		values = append(values, v)
	}
	return values
}

func MapSlice[T comparable](items []T) map[T]struct{} {
	m := make(map[T]struct{}, len(items))
	for _, item := range items {
		m[item] = struct{}{}
	}
	return m
}

func MapApply[K comparable, V any](items map[K]V, fn func(key K, item V) V) map[K]V {
	m := make(map[K]V, len(items))
	for k, v := range items {
		m[k] = fn(k, v)
	}
	return m
}

func MapConvert[K comparable, V, R any](items map[K]V, fn func(key K, item V) R) map[K]R {
	m := make(map[K]R, len(items))
	for k, v := range items {
		m[k] = fn(k, v)
	}
	return m
}

func MapFilter[K comparable, V any](itemsMap map[K]V, fn func(key K, item V) bool) map[K]V {
	m := make(map[K]V, len(itemsMap))
	for k, v := range itemsMap {
		if fn(k, v) {
			m[k] = v
		}
	}
	return m
}

func MapReduceKeys[K comparable, V any](itemsMap map[K]V, fn func(a, b K) K) K {
	if len(itemsMap) == 0 {
		var zero K
		return zero
	}
	var key K
	for fKey := range itemsMap {
		key = fKey
		break
	}
	var fKey = key
	for k := range itemsMap {
		if k != fKey {
			key = fn(key, k)
		}
	}
	return key
}

func MapReduceValues[K comparable, V any](itemsMap map[K]V, fn func(a, b V) V) V {
	if len(itemsMap) == 0 {
		var zero V
		return zero
	}
	var fKey K
	for key := range itemsMap {
		fKey = key
		break
	}
	var val = itemsMap[fKey]
	for k, v := range itemsMap {
		if k != fKey {
			val = fn(val, v)
		}
	}
	return val
}
