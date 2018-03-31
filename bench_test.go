package main 

import "testing"

func BenchmarkAnswer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		actual, err := answer()
		if err != nil {
			b.Fatal(err)
		}
		if 42 != actual {
			b.Error("should be 42")
		}
	}
}