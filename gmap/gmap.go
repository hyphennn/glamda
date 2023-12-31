package gmap

func Map[K1, K2 comparable, V1, V2 any](m map[K1]V1, fc func(K1, V1) (K2, V2)) map[K2]V2 {
	ret := make(map[K2]V2, len(m))
	for k1, v1 := range m {
		k2, v2 := fc(k1, v1)
		ret[k2] = v2
	}
	return ret
}

func ForEach[K comparable, V any](m map[K]V, fc func(K, V)) {
	for k, v := range m {
		fc(k, v)
	}
}

func Reverse[K, V comparable](m map[K]V) map[V]K {
	ret := make(map[V]K, len(m))
	for k, v := range m {
		ret[v] = k
	}
	return ret
}

func SafeStore[K comparable, V any, M ~map[K]V](m M, k K, v V) {
	if m == nil {
		m = make(map[K]V)
	}
	m[k] = v
}
