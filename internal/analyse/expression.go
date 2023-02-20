package analyse

import (
	"strings"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// IsCorrectExpression check that the expression is correct
func (a *analyse) IsCorrectExpression() bool {
	switch {
	case !a.areCorrectNumbers():
	case !a.areCorrectOperators():
	case !a.areCorrectParentheses():
	case !a.areCorrectDots():
	case !a.areCorrectPowers():
	default:
		return true
	}
	return false
}

// !Tool Functions

// areCorrectNumbers check that the numbers are correct
func (a *analyse) areCorrectNumbers() bool {
	digit, flagDot := uint16(0), false

	for i, r := range *a.expr {
		if !(data.IsFloat(&r)) {
			flagDot = false
			digit = 0
			continue
		}

		if data.IsDot(&r) {
			if flagDot {
				e := rune((*a.expr)[i-1])
				data.Error = ierr.TwoRune{S: r, E: e, I: i}.Together()
				return false
			}
			flagDot = true
		}

		if digit++; digit == data.DigitLimit {
			data.Error = ierr.OneRune{R: r, I: i}.Limit()
			return false
		}
	}

	return true
}

// areCorrectOperators check that the operators are correct
func (a *analyse) areCorrectOperators() bool {
	n := data.Lenght - 1

	for i, r := range *a.expr {
		if i != 0 && data.IsOperator(&r) && i != n {
			before := rune((*a.expr)[i-1])
			after := rune((*a.expr)[i+1])

			isGood := isGoodAfter(&after)

			switch {
			case data.IsNumber(&before) && isGood:
			case data.IsRight(&before) && isGood:
			case data.IsLeft(&before) && isGood:
			case data.IsPow(&before) && isGood:
			case data.IsPi(&before) && isGood:
			default:
				data.Error = ierr.ThreeRune{
					B: before, M: r, A: after, I: i,
				}.Together()
				return false
			}
		}
	}

	return true
}

// areCorrectParentheses check that the parentheses are correct
func (a *analyse) areCorrectParentheses() bool {
	if strings.Contains(*a.expr, string(data.Left)+string(data.Right)) {
		data.Error = ierr.TwoRune{S: data.Left}
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
		data.Error = ierr.IncompleteParentheses
		return false
	}

	return true
}

// areCorrectDots check that the dots are correct
func (a *analyse) areCorrectDots() bool {
	n := data.Lenght - 1

	for i, r := range *a.expr {
		if data.IsDot(&r) && i != n {
			before := rune((*a.expr)[i-1])
			after := rune((*a.expr)[i+1])

			if !data.IsNumber(&after) {
				data.Error = ierr.ThreeRune{
					B: before, M: r, A: after, I: i,
				}.Together()
				return false
			}

			switch {
			case data.IsOperator(&before):
			case data.IsNumber(&before):
			case data.IsLeft(&before):
			case data.IsRoot(&before):
			case data.IsPow(&before):
			default:
				data.Error = ierr.ThreeRune{
					B: before, M: r, A: after, I: i,
				}.Together()
				return false
			}
		}
	}

	return true
}

// areCorrectPowers check that the powers are correct
func (a *analyse) areCorrectPowers() bool {
	n := data.Lenght - 1

	for i, r := range *a.expr {
		if data.IsPow(&r) && i != n {
			before := rune((*a.expr)[i-1])
			after := rune((*a.expr)[i+1])

			isGood := isGoodAfterPow(&after)

			switch {
			case data.IsNumber(&before) && isGood:
			case data.IsRight(&before) && isGood:
			case data.IsPi(&before) && isGood:
			default:
				data.Error = ierr.ThreeRune{
					B: before, M: r, A: after, I: i,
				}.Together()
				return false
			}
		}
	}

	return true
}
