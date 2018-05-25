package concat_test

import (
	"os"
	"strings"
	"testing"
)

var args = []string{"hello", "world", "buddy", "how", "are", "you", "7", "8", "9"}

func concat(args []string) {
	r, sep := "", ""
	for _, a := range args {
		r += sep + a
		sep = " "
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(os.Args[1:])
	}


func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}
