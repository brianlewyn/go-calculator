package basic

import (
	"github.com/brianlewyn/go-calculator/internal/analyse"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/rebuild"
	"github.com/brianlewyn/go-calculator/internal/tokenize"
)

// Calculator represents the basic calculator
type Calculator struct {
	Expr string
}

// New creates an Calculator instance
func New(expr string) *Calculator {
	data.Lenght = len(expr)
	return &Calculator{Expr: expr}
}

// Calculate solves some expression and returns true
// if everything went well and otherwise returns false
func (c *Calculator) Calculate() bool {
	// rebuilder
	rebuilder := rebuild.New(&c.Expr)

	rebuilder.RemoveSpaces()
	rebuilder.AddAsterisks()
	rebuilder.AddZeros()

	// analyser
	analyser := analyse.New(&c.Expr)

	if !analyser.IsCorrectSyntax() {
		return false
	}

	if !analyser.IsCorrectExpression() {
		return false
	}

	// tokenizer
	tokenizer := tokenize.New(&c.Expr)
	tokenizer.LinkedListTokens()

	return true
}

// Answer returns a possible answer
func (c Calculator) Answer() float64 {
	return data.Answer
}

// Error returns a possible error
func (e Calculator) Error() error {
	return data.Error
}

/*
!linked list
list *doubly.Doubly[data.Token]
listTemp *doubly.Doubly[data.Token]
*/
