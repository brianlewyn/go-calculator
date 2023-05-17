package basic

import (
	"github.com/brianlewyn/go-calculator/internal/analyse"
	"github.com/brianlewyn/go-calculator/internal/math"
	"github.com/brianlewyn/go-calculator/internal/tokenize"
)

// Calculate solves a basic mathematical expression and returns the result and nil,
// otherwise it returns a zero value and an error.
func Calculate(expression string) (float64, error) {
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
