package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

func (a *analyse) areCorrectNumbers() bool {
	digit, isDot := uint16(0), false

	for i, r := range *a.expr {
		if !(data.Numbers(&r) || r == data.Dot) {
			isDot = false
			digit = 0
			continue
		}

		if r == data.Dot {
			if isDot {
				*a.err = ierr.Rune{R: &r, I: &i}.Dot()
				return false
			}
			isDot = true
		}

		if digit++; digit == data.DigitLimit {
			*a.err = ierr.Rune{R: &r, I: &i}.Digit()
			return false
		}
	}

	return true
}

func (a analyse) areCorrectOperators() bool { return true }

func (a analyse) areCorrectParentheses() bool { return true }

func (a analyse) areCorrectDots() bool { return true }

func (a *analyse) IsCorrectExpression() bool {
	if !a.areCorrectNumbers() {
		return false
	}

	if !a.areCorrectOperators() {
		return false
	}

	if !a.areCorrectParentheses() {
		return false
	}

	if !a.areCorrectDots() {
		return false
	}

	return true
}
