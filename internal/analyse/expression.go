package analyse

import (
	"strings"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// isCorrectExpression returns true if the expression is correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) isCorrectExpression() bool {
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

// !Tool Methods

// areCorrectNumbers returns true is the numbers are correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) areCorrectNumbers() bool {
	digit, flagDot := uint16(0), false

	for i, r := range *a.expression {
		if !(data.IsFloat(&r)) {
			flagDot = false
			digit = 0
			continue
		}

		if data.IsDot(&r) {
			if flagDot {
				e := rune((*a.expression)[i-1])
				eError = ierr.TwoRune{S: r, E: e, I: i}.Together()
				return false
			}
			flagDot = true
		}

		if digit++; digit == data.DigitLimit {
			eError = ierr.OneRune{R: r, I: i}.Limit()
			return false
		}
	}

	return true
}

// areCorrectOperators returns true if the operators are correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) areCorrectOperators() bool {
	n := *a.lenght - 1

	for i, r := range *a.expression {
		if i != 0 && data.IsOperator(&r) && i != n {
			before := rune((*a.expression)[i-1])
			after := rune((*a.expression)[i+1])

			is := data.IsAfter(&after)

			switch {
			case data.IsNumber(&before) && is:
			case data.IsRight(&before) && is:
			case data.IsLeft(&before) && is:
			case data.IsPow(&before) && is:
			case data.IsPi(&before) && is:
			default:
				eError = ierr.ThreeRune{
					B: before, M: r, A: after, I: i,
				}.Together()
				return false
			}
		}
	}

	return true
}

// areCorrectParentheses returns true if the parentheses are correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) areCorrectParentheses() bool {
	if strings.Contains(*a.expression, string(data.Left)+string(data.Right)) {
		eError = ierr.TwoRune{S: data.Left}
		return false
	}

	var nLeft, nRight int
	for _, r := range *a.expression {
		if data.IsLeft(&r) {
			nLeft++
			continue
		}
		if data.IsRight(&r) {
			nRight++
		}
	}

	if nLeft != nRight {
		eError = ierr.IncompleteParentheses
		return false
	}

	return true
}

// areCorrectDots returns true if the dots are correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) areCorrectDots() bool {
	n := *a.lenght - 1

	for i, r := range *a.expression {
		if data.IsDot(&r) && i != n {
			before := rune((*a.expression)[i-1])
			after := rune((*a.expression)[i+1])

			if !data.IsNumber(&after) {
				eError = ierr.ThreeRune{
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
				eError = ierr.ThreeRune{
					B: before, M: r, A: after, I: i,
				}.Together()
				return false
			}
		}
	}

	return true
}

// areCorrectPowers returns true if the powers are correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) areCorrectPowers() bool {
	n := *a.lenght - 1

	for i, r := range *a.expression {
		if data.IsPow(&r) && i != n {
			before := rune((*a.expression)[i-1])
			after := rune((*a.expression)[i+1])

			is := data.IsAfterPow(&after)

			switch {
			case data.IsNumber(&before) && is:
			case data.IsRight(&before) && is:
			case data.IsPi(&before) && is:
			default:
				eError = ierr.ThreeRune{
					B: before, M: r, A: after, I: i,
				}.Together()
				return false
			}
		}
	}

	return true
}
