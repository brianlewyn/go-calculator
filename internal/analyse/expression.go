package analyse

import (
	"strings"

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
				*a.err = ierr.TwoRune{S: r, E: e, I: i}.Together()
				return false
			}
			isDot = true
		}

		if digit++; digit == data.DigitLimit {
			*a.err = ierr.OneRune{R: r, I: i}.Limit()
			return false
		}
	}

	return true
}

// areCorrectOperators
func (a *analyse) areCorrectOperators() bool {
	n := len(*a.expr) - 1

	for i, r := range *a.expr {
		if i != 0 && data.IsOperator(&r) && i != n {
			before := rune((*a.expr)[i-1])
			after := rune((*a.expr)[i+1])

			switch {
			case data.IsNumber(&before) && isGoodAfter(&after):
			case data.IsPi(&before) && isGoodAfter(&after):
			case data.IsRight(&before) && isGoodAfter(&after):
			case data.IsLeft(&before) && data.IsMoreLess(&r) && isGoodAfter(&after):
			default:
				*a.err = ierr.ThreeRune{
					B: before, M: r, A: after, I: i,
				}.Together()
				return false
			}
		}
	}

	return true
}

// areCorrectParentheses
func (a *analyse) areCorrectParentheses() bool {
	if strings.Contains(*a.expr, string(data.Left)+string(data.Right)) {
		*a.err = ierr.TwoRune{S: data.Left}
		return false
	}

	var nLeft, nRight int
	for _, r := range *a.expr {
		if data.IsLeft(&r) {
			nLeft++
			continue
		}
		if data.IsRight(&r) {
			nRight++
		}
	}

	if nLeft != nRight {
		*a.err = ierr.IncompleteParentheses
		return false
	}

	return true
}

// areCorrectDots
func (a analyse) areCorrectDots() bool {
	return true
}

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
