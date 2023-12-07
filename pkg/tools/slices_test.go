package tools

import (
	"testing"
)

func TestSliceRand(t *testing.T) {
	s := []int{1, 2, 3, 4, 4, 5, 6, 7}
	t.Log(SliceRand(s, 3))
	t.Log(s)
	t.Log(SliceRand([]int{0, 34}, 2))
	t.Log(SliceRand([]int{80}, 2))
}

func TestSliceShuffle(t *testing.T) {
	s := []int{1, 2, 3, 4, 4, 5, 6, 7}
	SliceShuffle(s)
	t.Log(s)
	SliceShuffle(s)
	t.Log(s)
}
