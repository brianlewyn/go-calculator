package basic

import (
	"github.com/brianlewyn/go-calculator/internal/analyse"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/decode"
	"github.com/brianlewyn/go-calculator/internal/rebuild"
)

// Calculator
type Calculator struct {
	Expr string
	err  *error

	// start and end of temporary list
	// inside original list
	// start, end int

	// linked list
	original *[]*data.Token // *doubly.Doubly[data.Token]
	// temporary *[]*data.Token // *doubly.Doubly[data.Token]
}

func New(expr string) *Calculator {
	return &Calculator{Expr: expr}
}

// func (c *Calculator) newOriginalList() {}

// func (c *Calculator) newTemporaryList() {}

// func (c *Calculator) removeTemporaryFromOriginal() {}

// func (c *Calculator) countSubExpr() int { return 0 }

// func (c *Calculator) solve() {}

// func (c *Calculator) process() {}

func (c *Calculator) Calculate() bool {
	// rebuilder
	rebuilder := rebuild.New(&c.Expr)

	rebuilder.RemoveGaps()
	rebuilder.AddAsterisk()
	rebuilder.AddZero()

	// analyser
	analyser := analyse.New(&c.Expr)

	if !analyser.IsCorrectSyntax() {
		c.err = analyser.Error()
		return false
	}

	if !analyser.IsCorrectExpression() {
		c.err = analyser.Error()
		return false
	}

	// decoder
	decoder := decode.New(&c.Expr)
	decoder.FillAndTokenize()
	decoder.LinkedList(c.original)

	return true
}

func (c Calculator) Answer() float64 { return 0 }

func (e Calculator) Error() error { return *e.err }
