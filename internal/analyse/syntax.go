package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// isEmptyField returns true if the field is empty, otherwise returns false.
// But if there are any errors, an error is created and stored
func (a *analyse) isEmptyField() bool {
	if len(*a.expr) == 0 {
		*a.err = ierr.EmptyField
	}
	return true
}

// isProperSyntax returns true if is the proper syntax, otherwise returns false
// But if there are any errors, an error is created and stored
func (a *analyse) isProperSyntax() bool {
	for i, r := range *a.expr {
		if !isGoodChar(&r) {
			*a.err = ierr.OneRune{R: r, I: i}.Character()
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
		*a.err = ierr.OneRune{R: char, I: start}.Start()
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
		*a.err = ierr.OneRune{R: char, I: end}.Final()
		return false
	}

	return false
}

// areThereDuplicates returns true if there are duplicate characters, otherwise returns false
// But if there are any errors, an error is created and stored
func (a *analyse) areThereDuplicates() bool {
	d := &duplicate{expr: a.expr}

	if d.duplicates(data.IsOperator) {
		*a.err = ierr.TwoRune{
			S: *d.start, E: *d.end, I: *d.index,
		}.Together()
		return false
	}

	if d.duplicates(data.IsPow) {
		*a.err = ierr.TwoRune{
			S: *d.start, E: *d.end, I: *d.index,
		}.Together()
		return false
	}

	if d.duplicates(data.IsPi) {
		*a.err = ierr.TwoRune{
			S: *d.start, E: *d.end, I: *d.index,
		}.Together()
		return false
	}

	if d.duplicates(data.IsDot) {
		*a.err = ierr.TwoRune{
			S: *d.start, E: *d.end, I: *d.index,
		}.Together()
		return false
	}

	return true
}

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
