package iter

import (
	"slices"
	"testing"
)

// From https://github.com/DeedleFake/xiter/blob/master/transform_test.go

func TestMap(t *testing.T) {
	s := slices.Values([]int{1, 2, 3})
	n := slices.Collect(Map(s, func(v int) float64 { return float64(v * 2) }))
	if [3]float64(n) != [...]float64{2, 4, 6} {
		t.Fatal(n)
	}
}
