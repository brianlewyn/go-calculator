package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

func isSymbol(r *rune) bool {
	if data.Numbers(r) {
		return true
	}

	if data.Runes(r) {
		return true
	}

	return false
}

func (a *analyse) IsCorrectSyntax() bool {
	for i, r := range *a.expr {
		if !isSymbol(&r) {
			*a.err = ierr.Rune{R: &r, I: &i}.Character()
			return false
		}
	}
	return true
}
