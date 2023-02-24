package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// isCorrectExpression returns true if the expression is correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) isCorrectExpression() bool {
	var nLeft, nRight int
	var nDigit uint16
	var fDot bool

	for i, r := range *a.expression {
		switch {
		case !a.areCorrectNumbers(i, r, &nDigit, &fDot):
		case !a.areCorrectOperators(i, r):
		case !a.areCorrectParentheses(i, r, &nLeft, &nRight):
		case !a.areCorrectDots(i, r):
		case !a.areCorrectPowers(i, r):
		default:
			continue
		}
		return false
	}

	return true
}

// !Tool Methods

// areCorrectNumbers returns true is the numbers are correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) areCorrectNumbers(i int, r rune, n *uint16, f *bool) bool {
	if !(data.IsFloat(&r)) {
		*f = false
		*n = 0
		return true
	}

	if data.IsDot(&r) {
		if *f {
			e := rune((*a.expression)[i-1])
			eError = ierr.TwoRune{S: r, E: e, I: i}.Together()
			return false
		}
		*f = true
	}

	if *n++; *n == data.DigitLimit {
		eError = ierr.OneRune{R: r, I: i}.Limit()
		return false
	}

	return true
}

// areCorrectOperators returns true if the operators are correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) areCorrectOperators(i int, r rune) bool {
	if !data.IsOperator(&r) || i == 0 || i == *a.lenght-1 {
		return true
	}

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

	return true
}

// areCorrectParentheses returns true if the parentheses are correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) areCorrectParentheses(i int, r rune, nLeft, nRight *int) bool {
	if data.IsLeft(&r) {
		*nLeft++
	}

	if data.IsRight(&r) {
		*nRight++
	}

	if i != *a.lenght-1 {
		return true
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
func (a *analyse) areCorrectDots(i int, r rune) bool {
	if !data.IsDot(&r) || i == *a.lenght-1 {
		return true
	}

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

	return true
}

// areCorrectPowers returns true if the powers are correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) areCorrectPowers(i int, r rune) bool {
	if !data.IsPow(&r) || i == *a.lenght-1 {
		return true
	}

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

	return true
}
