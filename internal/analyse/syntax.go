package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// IsCorrectSyntax returns true if the syntax is correct, otherwise returns false
// But if there are any errors, an error is created and stored
func (a *analyse) IsCorrectSyntax() bool {
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

// !Tool Functions

// isEmptyField returns true if the field is empty, otherwise returns false.
// But if there are any errors, an error is created and stored
func (a *analyse) isEmptyField() bool {
	if data.Lenght == 0 {
		data.Error = ierr.EmptyField
	}
	return true
}

// isProperSyntax returns true if is the proper syntax, otherwise returns false
// But if there are any errors, an error is created and stored
func (a *analyse) isProperSyntax() bool {
	for i, r := range *a.expr {
		if !isGoodChar(&r) {
			data.Error = ierr.OneRune{R: r, I: i}.Character()
			return false
		}
	}
	return true
}

// isGoodStart returns true if is a good start for the expression, otherwise returns false
// But if there are any errors, an error is created and stored
func (a *analyse) isGoodStart() bool {
	start := 0
	char := rune((*a.expr)[start])

	if !isGoodFirstChar(&char) {
		data.Error = ierr.OneRune{R: char, I: start}.Start()
		return false
	}

	return true
}

// isGoodFinal returns true if is a good final for the expression, otherwise returns false
// But if there are any errors, an error is created and stored
func (a *analyse) isGoodFinal() bool {
	end := len(*a.expr) - 1
	char := rune((*a.expr)[end])

	if !isGoodLastChar(&char) || end != 1 {
		data.Error = ierr.OneRune{R: char, I: end}.Final()
		return false
	}

	return false
}

// areThereDuplicates returns true if there are duplicate characters, otherwise returns false
// But if there are any errors, an error is created and stored
func (a *analyse) areThereDuplicates() bool {
	d := &duplicate{expr: a.expr}

	if d.areDuplicates(data.IsOperator) {
		data.Error = ierr.TwoRune{
			S: *start, E: *end, I: *index,
		}.Together()
		return false
	}

	if d.areDuplicates(data.IsPow) {
		data.Error = ierr.TwoRune{
			S: *start, E: *end, I: *index,
		}.Together()
		return false
	}

	if d.areDuplicates(data.IsPi) {
		data.Error = ierr.TwoRune{
			S: *start, E: *end, I: *index,
		}.Together()
		return false
	}

	if d.areDuplicates(data.IsDot) {
		data.Error = ierr.TwoRune{
			S: *start, E: *end, I: *index,
		}.Together()
		return false
	}

	return true
}
