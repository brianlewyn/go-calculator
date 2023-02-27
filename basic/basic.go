package basic

import (
	"github.com/brianlewyn/go-calculator/internal/analyse"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/tokenize"
)

/*
Calculate solves a basic math expression

returns the result if it is correct and nil,
otherwise, returns a zero value and an error.
*/
func Calculate(expr string) (float64, error) {
	dataExpr := data.New(&expr)

	list, err := tokenize.Tokenizer(dataExpr)
	if err != nil {
		return 0, err
	}

	err = analyse.Analyser(list)
	if err != nil {
		return 0, err
	}

	return 0 /*Answer*/, nil
}

/*
!linked list
list *doubly.Doubly[data.Token]
listTemp *doubly.Doubly[data.Token]
*/
