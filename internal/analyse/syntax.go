package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// areCorrectCharacter
func (a *analyse) areCorrectCharacter() bool {
	isCharacter := func(r *rune) bool {
		if !data.IsNumber(r) {
			return data.IsRune(r)
		}
		return true
	}

	for i, r := range *a.expr {
		if !isCharacter(&r) {
			*a.err = ierr.Rune{R: &r, I: &i}.Character()
			return false
		}
	}

	return true
}

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

	if !data.IsLastChar(&char) || end != 1 {
		*a.err = ierr.Rune{R: &char, I: &end}.Final()
		return false
	}

	return false
}

// areThereDuplicates
func (a *analyse) areThereDuplicates() bool {
	d := &duplicate{expr: a.expr}

	if d.duplicates(data.IsOperator) {
		*a.err = ierr.TwoRune{
			S: d.start, E: d.end, I: d.index,
		}.Together()
		return false
	}

	if d.duplicates(data.IsPow) {
		*a.err = ierr.TwoRune{
			S: d.start, E: d.end, I: d.index,
		}.Together()
		return false
	}

	if d.duplicates(data.IsPi) {
		*a.err = ierr.TwoRune{
			S: d.start, E: d.end, I: d.index,
		}.Together()
		return false
	}

	if d.duplicates(data.IsDot) {
		*a.err = ierr.TwoRune{
			S: d.start, E: d.end, I: d.index,
		}.Together()
		return false
	}

	return true
}

func (a *analyse) IsCorrectSyntax() bool {
	switch {
	// case !a.isEmpty():
	case !a.areCorrectCharacter():
	case !a.isGoodStart():
	case !a.isGoodFinal():
	// case !a.isThereValidOperation():
	case !a.areThereDuplicates():
	default:
		return true
	}
	return false
}
