package benchmarkAnagram

import (
	"anagram/python/test/anagram-golang/benchmark/anaexec"
	"testing"
)

var gs string

// Sub benchmarking
func BenchmarkAnagram(b *testing.B) {
	b.Run("none", benchAnagram)
}

// benchmark for Anagram
func benchAnagram(b *testing.B) {

	for i := 0; i < b.N; i++ {
		anaexec.Anaexec()
	}
}
