package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// areCorrectNumbers
func (a *analyse) areCorrectNumbers() bool {
	digit, isDot := uint16(0), false

	for i, r := range *a.expr {
		if !(data.IsNumber(&r) || r == data.Dot) {
			isDot = false
			digit = 0
			continue
		}

		if r == data.Dot {
			if isDot {
				e := rune((*a.expr)[i-1])
				*a.err = ierr.TwoRune{S: &r, E: &e, I: &i}.Together()
				return false
			}
			isDot = true
		}

		if digit++; digit == data.DigitLimit {
			*a.err = ierr.OneRune{R: &r, I: &i}.Limit()
			return false
		}
	}

	return true
}

// areCorrectOperators
func (a analyse) areCorrectOperators() bool { return true }

// areCorrectParentheses
func (a analyse) areCorrectParentheses() bool { return true }

// areCorrectDots
func (a analyse) areCorrectDots() bool { return true }

// IsCorrectExpression
func (a *analyse) IsCorrectExpression() bool {
	switch {
	case !a.areCorrectNumbers():
	case !a.areCorrectOperators():
	case !a.areCorrectParentheses():
	case !a.areCorrectDots():
	default:
		return true
	}
	return false
}
