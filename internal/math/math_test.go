package math

import (
	"fmt"
	"testing"

	"github.com/brianlewyn/go-calculator/internal/analyse"
	"github.com/brianlewyn/go-calculator/internal/doubly"
	"github.com/brianlewyn/go-calculator/internal/tokenize"
)

func TestMath(t *testing.T) {
	result, err := Math(toList("(0.5 + 4.5 - 1) * 10 * √(6-2) / 4^2"))
	if err != nil {
		t.Errorf("RESULT = %f\n", result)
		t.Errorf("ERROR = %f\n", err)
	}
	// t.Fatalf("RESULT = %0.2f\n", result) // RESULT = 5
}

// toList returns the expression in a raw Tokenized Linked List
func toList(expression string) *doubly.Doubly {
	list, err1 := tokenize.Tokenizer(expression)
	if err1 != nil {
		fmt.Printf("ERROR [1]: %s\n\n", err1)
	}

	err2 := analyse.Analyser(list)
	if err2 != nil {
		fmt.Printf("ERROR [2]: %s\n\n", err2)
	}

	return list
}

// go test -bench=BenchmarkMath -benchmem -count=10 -benchtime=100x >> bench.txt
func BenchmarkMath(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Math(toList("(0.5 + 4.5 - 1) * 10 * √(6-2) / 4^2"))
	}
}
