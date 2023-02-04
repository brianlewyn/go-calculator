package analyse

type duplicate struct {
	expr  *string
	start *rune
	end   *rune
	index *int
}

func (d *duplicate) duplicates(data func(r *rune) bool) bool {
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
