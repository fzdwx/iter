package mapx

func Eq[K comparable, V any](m1 map[K]V, m2 map[K]V, cmp func(v1, v2 V) bool) bool {
	if len(m1) == 0 && len(m2) == 0 {
		return true
	}
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if !cmp(m2[k], v) {
			return false
		}
	}
	return true
}

func EqDefault[K, V comparable](m1 map[K]V, m2 map[K]V) bool {
	return Eq(m1, m2, func(v1, v2 V) bool {
		return v1 == v2
	})
}
