package nlp

// https://www.linkedin.com/learning/go-essential-training-16567666/benchmarking-and-profiling?autoSkip=true&autoplay=true&resume=false&u=55937129
//
import (
	"testing"

	"github.com/stretchr/testify/require"
)

var benchText = "Don't communicate by sharing memory, share memory by communicating."

func BenchmarkTokenize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokens := Tokenize(benchText)
		require.Equal(b, 10, len(tokens))
	}
}

func TestTokenize(t *testing.T) {
	text := "Who's on first?"
	expected := []string{"who", "s", "on", "first"}
	tokens := Tokenize(text)
	require.Equal(t, expected, tokens)
}
