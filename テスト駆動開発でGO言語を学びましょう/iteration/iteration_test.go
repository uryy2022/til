package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Repeat("a")
	}
}

func ExampleRepeat() {
	Repeats := Repeat("a")
	fmt.Printf("result=%q", Repeats)
	// Output: result="aaaaa"

}
