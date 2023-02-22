package analyse

// duplicate represents a scanner that looks for duplicates
type duplicate struct {
	expr  *string
	start *rune
	end   *rune
	index *int
}

// findDuplicates returns true if finds duplicates, otherwise returns false.
func (d *duplicate) findDuplicates(data func(r *rune) bool) bool {
	var isChar bool

	for i, r := range *d.expr {
		if !data(&r) {
			isChar = false
			continue
		}

		if isChar {
			*d.start = rune((*d.expr)[i-1])
			*d.end = r
			*d.index = i
			return false
		}

		isChar = true
	}

	return true
}
