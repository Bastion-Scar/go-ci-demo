package calc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMultiply(t *testing.T) {
	got := Multiply(1, 2)

	require.Equal(t, 2, got)
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Multiply(531, 12145)
	}
}
