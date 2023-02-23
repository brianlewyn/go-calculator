package rebuild

import (
	"github.com/brianlewyn/go-calculator/internal/data"
)

// rebuild represents a rebuilder for expressionession
type rebuild struct {
	expression *string
	lenght     *int
}

// trim removes all spaces inside the expression
func (r *rebuild) trim() {
	var expression string
	var lenght int = *r.lenght

	for *r.lenght > 0 {
		first := rune((*r.expression)[0])

		if !data.IsGap(&first) {
			expression += string(first)
			lenght++
		}

		*r.expression = (*r.expression)[1:]
		*r.lenght--
	}

	r.expression = &expression
	r.lenght = &lenght
}

// addAsterisks adds asterisks between the right and left parentheses
func (r *rebuild) addAsterisks(i int, current rune) {
	if i < *r.lenght-1 {
		if data.IsRight(&current) {
			after := rune((*r.expression)[i+1])

			if data.IsLeft(&after) {
				*r.expression = (*r.expression)[:i+1] + string(data.Mul) + (*r.expression)[i+1:]
				*r.lenght++
			}
		}
	}
}

// addZeros adds zeros between the left parentheses and the plus or minus operator
func (r *rebuild) addZeros(i int, current rune) {
	if data.IsMoreLess(&current) {
		if i == 0 {
			*r.expression = "0" + *r.expression
			*r.lenght++
			return
		}

		if i < *r.lenght-1 {
			before := rune((*r.expression)[i-1])

			if data.IsLeft(&before) {
				*r.expression = (*r.expression)[:i] + "0" + (*r.expression)[i:]
				*r.lenght++
			}
		}
	}
}

// Rebuilder rebuilds the basic math expressionession to a simpler form
func Rebuilder(data *data.Data) {
	rebuilder := &rebuild{
		expression: data.Expression(),
		lenght:     data.Lenght(),
	}

	rebuilder.trim()

	for i := 0; i < *rebuilder.lenght; i++ {
		current := rune((*rebuilder.expression)[i])
		rebuilder.addAsterisks(i, current)
		rebuilder.addZeros(i, current)
	}
}
