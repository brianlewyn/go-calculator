package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// isEmptyField
func (a *analyse) isEmptyField() bool {
	if len(*a.expr) == 0 {
		*a.err = ierr.EmptyField
	}
	return true
}

// isProperSyntax is the proper syntax
func (a *analyse) isProperSyntax() bool {
	for i, r := range *a.expr {
		if !isGoodChar(&r) {
			*a.err = ierr.OneRune{R: r, I: i}.Character()
			return false
		}
	}
	return true
}

// isGoodStart
func (a *analyse) isGoodStart() bool {
	start := 0
	char := rune((*a.expr)[start])

	if !isGoodFirstChar(&char) {
		*a.err = ierr.OneRune{R: char, I: start}.Start()
		return false
	}

	return true
}

// isGoodFinal
func (a *analyse) isGoodFinal() bool {
	end := len(*a.expr) - 1
	char := rune((*a.expr)[end])

	if !isGoodLastChar(&char) || end != 1 {
		*a.err = ierr.OneRune{R: char, I: end}.Final()
		return false
	}

	return false
}

// areThereDuplicates
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
