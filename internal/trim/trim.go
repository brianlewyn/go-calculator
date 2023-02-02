package trim

import "github.com/brianlewyn/go-calculator/internal/data"

type trim struct {
	size int
	char rune
	expr *string
}

func newTrim(expr *string, char rune) *trim {
	return &trim{size: len(*expr), char: char, expr: expr}
}

// contains
func (t *trim) contains() bool {
	for _, char := range *t.expr {
		if char == t.char {
			return true
		}
	}
	return false
}

// internalCut
func (t *trim) internalCut(i, j *int) {
	if *j != 0 {
		*t.expr = (*t.expr)[:*i-*j] + (*t.expr)[*i:]
		t.size -= *j
		*i -= *j
		*j = 0
	}
}

// finalCut
func (t *trim) finalCut(i, j *int) {
	if *j++; *i == t.size-1 {
		if *j != 0 {
			*t.expr = (*t.expr)[:*i-*j+1]
		}
	}
}

func (t *trim) remove() {
	if !t.contains() {
		return
	}

	for i, j := 0, 0; i <= t.size-1; i++ {
		if rune((*t.expr)[i]) == t.char {
			t.finalCut(&i, &j)
			continue
		}

		t.internalCut(&i, &j)
	}
}

func Gaps(expr *string) {
	t := newTrim(expr, data.Gap)
	t.remove()
}
