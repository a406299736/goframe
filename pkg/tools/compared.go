package tools

import "fmt"

// Ordered 代表所有可比大小排序的类型
type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

func MaxVal4Map[Key comparable, Val Integer](m map[Key]Val) (Key, Val) {
	var maxV Val
	var maxK Key
	for ky, vl := range m {
		if vl >= maxV {
			maxK = ky
			maxV = vl
		}
	}

	return maxK, maxV
}

func MaxVal4Slice[T Ordered](s ...T) (max T, err error) {
	l := len(s)
	if l == 0 {
		return max, fmt.Errorf("切片长度为空")
	}
	if l == 1 {
		return s[0], nil
	}
	for k := range s {
		if s[k] > max {
			max = s[k]
		}
	}
	return
}

func MinVal4Slice[T Ordered](s ...T) (min T, err error) {
	l := len(s)
	if l == 0 {
		return min, fmt.Errorf("切片长度为空")
	}
	if l == 1 {
		return s[0], nil
	}
	min = s[0]
	for k := range s {
		if s[k] < min {
			min = s[k]
		}
	}
	return
}

func MaxMin4Slice[T Ordered](s ...T) (max, min T, err error) {
	l := len(s)
	if l == 0 {
		return max, min, fmt.Errorf("切片长度为空")
	}
	if l == 1 {
		return s[0], s[0], nil
	}

	min, max = s[0], s[0]
	for k := range s {
		if s[k] < min {
			min = s[k]
		}
		if s[k] > max {
			max = s[k]
		}
	}
	return
}

func InSlice[C comparable](val C, s ...C) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if val == v {
			return true
		}
	}

	return false
}

func Slice2MapKey[T comparable](s []T) map[T]struct{} {
	t := make(map[T]struct{})
	for k := range s {
		t[s[k]] = struct{}{}
	}

	return t
}

// SliceUnique 去重
func SliceUnique[T comparable](in []T) (out []T) {
	t := make(map[T]struct{})
	for k := range in {
		if _, ok := t[in[k]]; !ok {
			out = append(out, in[k])
			t[in[k]] = struct{}{}
		}
	}

	return
}
