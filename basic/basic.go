package basic

import (
	"github.com/brianlewyn/go-calculator/internal/base"
	"github.com/brianlewyn/go-calculator/internal/correct"
)

type MathExpr struct {
	Expr   string
	list   base.List
	result base.Result
}

func NewMathExpr(expr string) *MathExpr {
	return &MathExpr{Expr: expr}
}

func (e *MathExpr) completeMathExpr() {}

func (e *MathExpr) newOriginalList() {}

func (e *MathExpr) newTemporaryList() {}

func (e *MathExpr) removeTemporaryFromOriginal() {}

func (e *MathExpr) countSubExpr() int { return 0 }

func (e *MathExpr) solve() {}

func (e *MathExpr) answer() float64 { return 0 }

func (e *MathExpr) process() {}

func (e *MathExpr) Calculate() (float64, error) {

	err := correct.IsCorrect(&e.Expr)
	if err != nil {
		return 0, err
	}

	return e.answer(), nil
}
