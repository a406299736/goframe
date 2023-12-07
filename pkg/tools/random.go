package tools

import (
	"math/rand"
	"slices"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// RandString 获取n个长度字符串
func RandString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())

	sb := strings.Builder{}
	sb.Grow(n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

// RandomInt 随机数， 左闭右开[0, 2) => 0,1
func RandomInt(start int, end int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Intn(end - start)
	random = start + random
	return random
}

// SliceRand[T any] 随机从原来切片中取N个元素，生成新的切片
func SliceRand[T any](s []T, n int) []T {
	if n <= 0 || n > len(s) {
		return nil
	}
	s1 := slices.Clone(s)
	SliceShuffle(s1)

	return s1[:n]
}
