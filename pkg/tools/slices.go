package tools

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"time"
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

// ToIntSlice 粗暴转化
func ToIntSlice(s []string) []int {
	i := make([]int, len(s))
	for k, v := range s {
		i[k], _ = strconv.Atoi(v)
	}
	return i
}

// ToStringSlice 粗暴转化
func ToStringSlice(s []int) []string {
	i := make([]string, len(s))
	for k, v := range s {
		i[k] = strconv.Itoa(v)
	}
	return i
}

// SliceSplit[T any] 将一个切片分割成多个切片
func SliceSplit[T any](s []T, size int) [][]T {
	if size <= 0 {
		return nil
	}

	var chunks [][]T
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

// MapSliceColumn[T any, K comparable] 从字典切片中获取某个key值
func MapSliceValues[T any, K comparable](maps []map[K]T, key K) []T {
	var column []T
	for _, m := range maps {
		if val, ok := m[key]; ok {
			column = append(column, val)
		}
	}
	return column
}

// InSlice[C comparable] 判断一个元素是否在切片
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

// Slice2MapKey[T comparable] 切片元素转成字典key:struct{}
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

// Shuffle[T any] 乱序
func SliceShuffle[T any](in []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(in), func(i, j int) {
		in[i], in[j] = in[j], in[i]
	})
}

// MaxVal4Map[Key comparable, Val cmp.Ordered] 获取字典最大值和对应key
func MaxVal4Map[Key comparable, Val cmp.Ordered](m map[Key]Val) (Key, Val) {
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

// MaxVal4Slice[T cmp.Ordered] 切片最大值
func MaxVal4Slice[T cmp.Ordered](s ...T) (max T, err error) {
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

// MinVal4Slice[T cmp.Ordered] 切片最小值
func MinVal4Slice[T cmp.Ordered](s ...T) (min T, err error) {
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

// MaxMin4Slice[T cmp.Ordered] 切片最大最小值
func MaxMin4Slice[T cmp.Ordered](s ...T) (max, min T, err error) {
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
