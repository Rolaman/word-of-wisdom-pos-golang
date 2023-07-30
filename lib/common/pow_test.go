package common

import (
	"crypto/rand"
	"fmt"
	"testing"
)

var diffs = []byte{1, 2, 3}

func BenchmarkSolver(b *testing.B) {
	for _, diff := range diffs {
		b.Run(fmt.Sprintf("difficulty_%d", diff), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buf := make([]byte, 8)
				_, err := rand.Read(buf)
				if err != nil {
					b.Fatalf("can't fill slice: %v", err)
				}
				solver, err := NewSolver(buf, diff)
				if err != nil {
					b.Fatalf("can't init solver: %v", err)
				}
				_, err = solver.FindNonce()
				if err != nil {
					b.Fatalf("can't find solution: %v", err)
				}
			}
		})
	}
}
