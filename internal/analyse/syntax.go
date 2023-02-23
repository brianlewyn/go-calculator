package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// isCorrectSyntax returns true if the syntax is correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) isCorrectSyntax() bool {
	switch {
	case a.isEmptyField():
	case !a.isProperSyntax():
	default:
		return a.areThereDuplicates()
	}
	return false
}

// !Tool Methods

// isEmptyField returns true if the field is empty, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) isEmptyField() bool {
	if *a.lenght == 0 {
		eError = ierr.EmptyField
	}
	return true
}

// isProperSyntax returns true if is the proper syntax, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) isProperSyntax() bool {
	for i, r := range *a.expression {
		switch i {
		case 0:
			a.isGoodStart(r)
		case *a.lenght:
			a.isGoodFinal(r)
		default:
			if !data.IsRuneSyntax(&r) {
				eError = ierr.OneRune{R: r, I: i}.Character()
				return false
			}
		}
	}
	return true
}

// isGoodStart returns true if is a good start for the expression, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) isGoodStart(r rune) bool {
	if start := 0; !data.IsFirst(&r) {
		eError = ierr.OneRune{R: r, I: start}.Start()
		return false
	}
	return true
}

// isGoodFinal returns true if is a good final for the expression, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) isGoodFinal(r rune) bool {
	if !data.IsLast(&r) {
		eError = ierr.OneRune{R: r, I: *a.lenght}.Final()
		return false
	}
	return false
}

// areThereDuplicates returns true if there are duplicate characters, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error()
func (a *analyse) areThereDuplicates() bool {
	var isOperator, isDot, isPow, isPi bool
	d := &duplicate{expression: a.expression}

	for i, r := range *a.expression {
		switch {
		case d.findDuplicates(i, r, &isOperator, data.IsOperator):
		case d.findDuplicates(i, r, &isDot, data.IsDot):
		case d.findDuplicates(i, r, &isPow, data.IsPow):
		case d.findDuplicates(i, r, &isPi, data.IsPi):
		default:
			continue
		}
		return true
	}

	return false
}
