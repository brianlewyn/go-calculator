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
	data.Lenght = len(expr)

	rebuild.Rebuilder(&expr)

	if analyse.Analyser(&expr) {
		return 0, data.Error
	}

	tokenize.Tokenizer(&expr)

	return data.Answer, data.Error
}

/*
!linked list
list *doubly.Doubly[data.Token]
listTemp *doubly.Doubly[data.Token]
Remember: n^±., ^±n, ^±π, ^±(, ^±√,
*/
