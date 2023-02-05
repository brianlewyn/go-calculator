package basic

import (
	"github.com/brianlewyn/go-calculator/internal/analyse"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/rebuild"
)

// token
type token struct {
	kind  data.TokenKind
	value string
}

// Calculator
type Calculator struct {
	Expr string
	err  *error

	// start and end of temporary list
	// inside original list
	start, end int

	// linked list
	original  *[]string // *doubly.Doubly[token]
	temporary *[]string // *doubly.Doubly[token]
}

func New(expr string) *Calculator {
	return &Calculator{Expr: expr}
}

func (c *Calculator) newOriginalList() {}

func (c *Calculator) newTemporaryList() {}

func (c *Calculator) removeTemporaryFromOriginal() {}

func (c *Calculator) countSubExpr() int { return 0 }

func (c *Calculator) solve() {}

func (c *Calculator) process() {}

func (c *Calculator) Calculate() bool {
	rebuilder := rebuild.New(&c.Expr)

	rebuilder.RemoveGaps()
	rebuilder.AddAsterisk()

	analyser := analyse.New(&c.Expr)

	if !analyser.IsCorrectSyntax() {
		c.err = analyser.Error()
		return false
	}

	if !analyser.IsCorrectExpression() {
		c.err = analyser.Error()
		return false
	}

	// list.Fill(&c.Expr, c.list.Original)

	return true
}

func (c Calculator) Answer() float64 { return 0 }

func (e Calculator) Error() error { return *e.err }
