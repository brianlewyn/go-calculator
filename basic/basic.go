package basic

import (
	"github.com/brianlewyn/go-calculator/internal/analyse"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/rebuild"
	"github.com/brianlewyn/go-calculator/internal/tokenize"
)

// Calculate solves a basic math expression and returns the result if it is correct and nil;
//
// otherwise, it returns a zero and an error.
func Calculate(expr string) (float64, error) {
	data := data.New(&expr)

	rebuild.Rebuilder(data)

	if analyse.Analyser(data) {
		return 0, analyse.Error()
	}

	tokenize.Tokenizer(data)

	return 0 /*Answer*/, nil
}

/*
!linked list
list *doubly.Doubly[data.Token]
listTemp *doubly.Doubly[data.Token]
Remember: n^±., ^±n, ^±π, ^±(, ^±√,
*/
