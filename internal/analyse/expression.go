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

func (a analyse) inconsistentOperators() {}

func (a analyse) inconsistentParentheses() {}

func (a analyse) inconsistentBlankSpakces() {}

func (a analyse) inconsistentDecimalPoints() {}

func (a *analyse) IsCorrectExpression() bool {

	if !a.areCorrectNumbers() {
		return false
	}

	a.inconsistentOperators()
	a.inconsistentParentheses()
	a.inconsistentBlankSpakces()
	a.inconsistentDecimalPoints()

	return true
}
