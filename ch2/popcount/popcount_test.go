package popcount

import (
	"testing"

	"github.com/anuar45/kern/ch2/popcount"
)

func BenchPopCount(b *testing.B) {
	popcount.PopCount(uint64(b.N))
}

func BenchPopCountFor(b *testing.B) {
	popcount.PopCountFor(uint64(b.N))
}

func BenchPopCount64(b *testing.B) {
	popcount.PopCount64(uint64(b.N))
}

func BenchPopCounSpect(b *testing.B) {
	popcount.PopCountSpec(uint64(b.N))
}
