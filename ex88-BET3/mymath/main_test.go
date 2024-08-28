package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data []int
		ans  float64
	}

	tests := []test{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{-1, -2, -3, -4, -5}, -3},
		{[]int{-2, -1, 0, 1, 2}, 0},
		{[]int{0, 0, 0, 0, 0}, 0},
	}

	for _, v := range tests {
		computed := CenteredAvg(v.data)
		if computed != v.ans {
			t.Error("got", computed, "want", v.ans, v.data)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5}))
	// Output:
	// 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6})
	}
}
