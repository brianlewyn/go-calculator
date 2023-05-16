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
func Calculate(expression string) (float64, data.Error) {
	list, err := tokenize.Tokenizer(expression)
	if err != nil {
		return 0, err
	}

	err = analyse.Analyser(list)
	if err != nil {
		return 0, err
	}

	return math.Math(list)
}

/*
!linked list
list *doubly.Doubly[data.Token]
listTemp *doubly.Doubly[data.Token]
*/
