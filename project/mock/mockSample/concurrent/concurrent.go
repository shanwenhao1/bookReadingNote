//go:generate mockgen -destination mock/concurrent_mock.go bookReadingNote/project/mock/mockSample/concurrent Math

// Package concurrent demonstrates how to use gomock with goroutines.
package concurrent

type Math interface {
	Sum(a, b int) int
}
