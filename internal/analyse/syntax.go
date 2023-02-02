package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

func symbols(r *rune) bool {
	if data.Numbers(r) {
		return true
	}

	if data.Runes(r) {
		return true
	}

	return false
}

func (a *analyse) IsCorrectSyntax() bool {
	for _, r := range *a.expr {
		if !symbols(&r) {
			a.err = ierr.Syntax{S: &r}.Wrap()
			return true
		}
	}
	return false
}
