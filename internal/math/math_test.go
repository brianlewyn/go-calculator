package math

import (
	"fmt"
	"testing"

	"github.com/brianlewyn/go-calculator/internal/analyse"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
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
func toList(expression string) *plugin.TokenList {
	list, err1 := tokenize.Tokenizer(data.NewInfo(&expression))
	if err1 != nil {
		fmt.Printf("ERROR [1]: %s\n\n", err1)
	}

	err2 := analyse.Analyser(list)
	if err2 != nil {
		fmt.Printf("ERROR [2]: %s\n\n", err2)
	}

	return list
}
