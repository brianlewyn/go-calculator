package rebuild

import (
	"strings"

	"github.com/brianlewyn/go-calculator/internal/data"
)

type rebuild struct {
	size int
	expr *string
}

func New(expr *string) *rebuild {
	return &rebuild{size: len(*expr), expr: expr}
}

// internalCut
func (r *rebuild) internalCut(i, j *int) {
	if *j != 0 {
		*r.expr = (*r.expr)[:*i-*j] + (*r.expr)[*i:]
		r.size -= *j
		*i -= *j
		*j = 0
	}
}

// finalCut
func (r *rebuild) finalCut(i, j *int) {
	if *j++; *i == r.size-1 {
		if *j != 0 {
			*r.expr = (*r.expr)[:*i-*j+1]
		}
	}
}

func (r *rebuild) RemoveGaps() {
	if !strings.Contains(*r.expr, string(data.Gap)) {
		return
	}

	for i, j := 0, 0; i <= r.size-1; i++ {
		if rune((*r.expr)[i]) == data.Gap {
			r.finalCut(&i, &j)
			continue
		}

		r.internalCut(&i, &j)
	}
}

func (r *rebuild) AddAsterisk() {
	if !strings.Contains(*r.expr, string(data.Left)+string(data.Right)) {
		return
	}

	//
}
