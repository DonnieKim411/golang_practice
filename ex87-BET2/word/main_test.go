package word

import (
	"fmt"
	"puppy/ex87-BET2/quote"
	"testing"
)

func TestCount(t *testing.T) {
	n := Count("one two three")

	if n != 3 {
		t.Error("got", n, "want", 3)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two two three three three")
	cnt := 1
	for k, v := range m {
		if v != cnt {
			t.Error("For", k, "got", v, "want", cnt)
		}
		cnt++
	}

}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	//Output:
	//3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
