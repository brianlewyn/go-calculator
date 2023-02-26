package basic

import (
	"github.com/brianlewyn/go-calculator/internal/analyse"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/tokenize"
)

// Calculate solves a basic math expression and returns the result if it is correct and nil;
//
// otherwise, it returns a zero and an error.
func Calculate(expr string) (float64, error) {
	list := tokenize.Tokenizer(data.New(&expr))

	if analyse.Analyser(list) {
		return 0, data.Error
	}

	return 0 /*Answer*/, nil
}

/*
!linked list
list *doubly.Doubly[data.Token]
listTemp *doubly.Doubly[data.Token]
*/
