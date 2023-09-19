package tools

import "testing"

func TestRandString(t *testing.T) {
	t.Log(RandString(5))
	t.Log(RandString(7))
	t.Log(RandString(9))
	t.Log(RandString(12))
	t.Log(RandString(15))
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandString(10)
	}
}
