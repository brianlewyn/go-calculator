package analyse

import "github.com/brianlewyn/go-calculator/ierr"

// duplicate represents a scanner that looks for duplicates
type duplicate struct {
	expression *string
}

// findDuplicates returns true if finds duplicates, otherwise returns false.
func (d *duplicate) findDuplicates(i int, r rune, flag *bool, data func(r *rune) bool) bool {
	if !data(&r) {
		*flag = false
		return false
	}

	if *flag {
		s := rune((*d.expression)[i-1])
		eError = ierr.TwoRune{S: s, E: r, I: i}.Together()
		return true
	}

	*flag = true
	return false
}
