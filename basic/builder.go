package basic

import (
	"github.com/brianlewyn/go-calculator/internal/analyse"
	"github.com/brianlewyn/go-calculator/internal/base"
	"github.com/brianlewyn/go-calculator/internal/trim"
)

type Calculator struct {
	Expr   string
	list   *base.List
	result *base.Result
	err    *error
}

func New(expr string) *Calculator {
	return &Calculator{Expr: expr}
}

func (c *Calculator) completeMathExpr() {}

func (c *Calculator) newOriginalList() {}

func (c *Calculator) newTemporaryList() {}

func (c *Calculator) removeTemporaryFromOriginal() {}

func (c *Calculator) countSubExpr() int { return 0 }

func (c *Calculator) solve() {}

func (c Calculator) Answer() float64 { return 0 }

func (c *Calculator) process() {}

func (c *Calculator) Calculate() bool {
	trim.Gaps(&c.Expr)

	analyser := analyse.New(&c.Expr)

	if !analyser.IsCorrectSyntax() {
		c.err = analyser.Error()
		return false
	}

	if !analyser.IsCorrectExpression() {
		c.err = analyser.Error()
		return false
	}

	return true
}

func (e Calculator) Error() error {
	return *e.err
}
