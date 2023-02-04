package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// isGoodStart
func (a *analyse) isGoodStart() bool {
	start := 0
	char := rune((*a.expr)[start])
	if !data.IsFirstChar(&char) {
		*a.err = ierr.Rune{R: &char, I: &start}.Start()
		return false
	}
	return true
}

// isGoodFinal
func (a *analyse) isGoodFinal() bool {
	end := len(*a.expr) - 1
	char := rune((*a.expr)[end])
	if !data.IsLastChar(&char) {
		*a.err = ierr.Rune{R: &char, I: &end}.Final()
		return false
	}
	return false
}

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

// areCorrectOperators
func (a analyse) areCorrectOperators() bool { return true }

// areCorrectParentheses
func (a analyse) areCorrectParentheses() bool { return true }

// areCorrectDots
func (a analyse) areCorrectDots() bool { return true }

// IsCorrectExpression
func (a *analyse) IsCorrectExpression() bool {
	switch {
	case a.isGoodStart():
	case a.isGoodFinal():
	case a.areCorrectNumbers():
	case a.areCorrectOperators():
	case a.areCorrectParentheses():
	case a.areCorrectDots():
	default:
		return false
	}
	return true
}
