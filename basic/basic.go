package basic

import (
	"github.com/brianlewyn/go-calculator/internal/analyse"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/math"
	"github.com/brianlewyn/go-calculator/internal/tokenize"
)

/*
Calculate solves a basic math expression and
returns the result if it is correct and nil,
otherwise, returns a zero value and an error.
*/
func Calculate(expr string) (float64, data.Error) {
	info := data.NewInfo(&expr)

	list, err := tokenize.Tokenizer(info)
	if err != nil {
		return 0, err
	}

	err = analyse.Analyser(list)
	if err != nil {
		return 0, err
	}

	return math.Math(list), nil
}

/*
!linked list
list *doubly.Doubly[data.Token]
listTemp *doubly.Doubly[data.Token]
*/
