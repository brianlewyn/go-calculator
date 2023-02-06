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

	for i, char := range *r.expr {
		if i != 0 && i != r.size-1 {
			after := rune((*r.expr)[i+1])
			if data.IsRight(&char) && data.IsLeft(&after) {
				*r.expr = (*r.expr)[:i+1] + string(data.Mul) + (*r.expr)[i+1:]
				r.size++
			}
		}
	}
}

func (r *rebuild) AddZero() {
	zero := "0"

	for i, char := range *r.expr {
		if data.IsMoreLess(&char) {
			if i == 0 {
				*r.expr = zero + *r.expr
				r.size++
				continue
			}

			if i != r.size-1 {
				before := rune((*r.expr)[i-1])
				if data.IsLeft(&before) {
					*r.expr = (*r.expr)[:i] + zero + (*r.expr)[i:]
					r.size++
				}
			}
		}
	}
}
