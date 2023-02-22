package rebuild

import (
	"strings"

	"github.com/brianlewyn/go-calculator/internal/data"
)

// rebuild represents a rebuilder for expression
type rebuild struct {
	expr *string
}

// removeSpaces removes all spaces inside the expression
func (r *rebuild) removeSpaces() {
	if !strings.Contains(*r.expr, string(data.Gap)) {
		return
	}

	for i, j := 0, 0; i <= data.Lenght-1; i++ {
		if rune((*r.expr)[i]) == data.Gap {

			// final cutout
			if j++; i == data.Lenght-1 {
				*r.expr = (*r.expr)[:i-j]
			}
			continue
		}

		// internal cutout
		if j != 0 {
			*r.expr = (*r.expr)[:i-j] + (*r.expr)[i:]
			data.Lenght -= j
			i -= j
			j = 0
		}
	}
}

// addAsterisks adds asterisks between the right and left parentheses
func (r *rebuild) addAsterisks() {
	if !strings.Contains(*r.expr, string(data.Left)+string(data.Right)) {
		return
	}

	for i, char := range *r.expr {
		if i != 0 && i != data.Lenght-1 {
			after := rune((*r.expr)[i+1])

			if data.IsRight(&char) && data.IsLeft(&after) {
				*r.expr = (*r.expr)[:i+1] + string(data.Mul) + (*r.expr)[i+1:]
				data.Lenght++
			}
		}
	}
}

// addZeros adds zeros between the left parentheses and the plus or minus operator
func (r *rebuild) addZeros() {
	zero := "0"

	for i, char := range *r.expr {
		if data.IsMoreLess(&char) {
			if i == 0 {
				*r.expr = zero + *r.expr
				data.Lenght++
				continue
			}

			if i != data.Lenght-1 {
				before := rune((*r.expr)[i-1])

				if data.IsLeft(&before) {
					*r.expr = (*r.expr)[:i] + zero + (*r.expr)[i:]
					data.Lenght++
				}
			}
		}
	}
}

// Rebuilder rebuilds the basic math expression to a simpler form
func Rebuilder(expr *string) {
	rebuilder := &rebuild{expr: expr}
	rebuilder.removeSpaces()
	rebuilder.addAsterisks()
	rebuilder.addZeros()
}
