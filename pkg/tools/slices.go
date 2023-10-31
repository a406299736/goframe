package tools

import (
	"cmp"
	"slices"
)

// SlicesDiff[T cmp.Ordered] s1与s2的差集
func SlicesDiff[T cmp.Ordered](s1, s2 []T) []T {
	d := slices.Compare(s1, s2)
	if d == 0 {
		return make([]T, 0)
	}

	dk := Slice2MapKey(s1)
	for _, val2 := range s2 {
		if _, ok := dk[val2]; ok {
			delete(dk, val2)
		}
	}

	var n []T
	for k := range dk {
		n = append(n, k)
	}

	return n
}
