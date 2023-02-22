package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// isCorrectSyntax returns true if the syntax is correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error
func (a *analyse) isCorrectSyntax() bool {
	switch {
	case !a.isEmptyField():
	case !a.isProperSyntax():
	case !a.isGoodStart():
	case !a.isGoodFinal():
	default:
		return a.areThereDuplicates()
	}
	return false
}

// !Tool Methods

// isEmptyField returns true if the field is empty, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error
func (a *analyse) isEmptyField() bool {
	if data.Lenght == 0 {
		data.Error = ierr.EmptyField
	}
	return true
}

// isProperSyntax returns true if is the proper syntax, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error
func (a *analyse) isProperSyntax() bool {
	for i, r := range *a.expr {
		if !data.IsRuneSyntax(&r) {
			data.Error = ierr.OneRune{R: r, I: i}.Character()
			return false
		}
	}
	return true
}

// isGoodStart returns true if is a good start for the expression, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error
func (a *analyse) isGoodStart() bool {
	start := 0
	char := rune((*a.expr)[start])

	if !data.IsFirst(&char) {
		data.Error = ierr.OneRune{R: char, I: start}.Start()
		return false
	}

	return true
}

// isGoodFinal returns true if is a good final for the expression, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error
func (a *analyse) isGoodFinal() bool {
	end := len(*a.expr) - 1
	char := rune((*a.expr)[end])

	if !data.IsLast(&char) || end != 1 {
		data.Error = ierr.OneRune{R: char, I: end}.Final()
		return false
	}

	return false
}

// areThereDuplicates returns true if there are duplicate characters, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error
func (a *analyse) areThereDuplicates() bool {
	d := &duplicate{expr: a.expr}

	switch {
	case !d.findDuplicates(data.IsOperator):
	case !d.findDuplicates(data.IsPow):
	case !d.findDuplicates(data.IsPi):
	case !d.findDuplicates(data.IsDot):
	default:
		return true
	}

	data.Error = ierr.TwoRune{
		S: *d.start, E: *d.end, I: *d.index,
	}.Together()
	return false
}
