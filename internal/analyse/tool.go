package analyse

import "github.com/brianlewyn/go-calculator/internal/data"

// !duplicate

// duplicate represents a parser that checks for duplicates
type duplicate struct {
	expr *string
}

// variables for duplicate struct
var (
	index *int
	start *rune
	end   *rune
)

// areDuplicates finds duplicates, if it is, returns true
func (d *duplicate) areDuplicates(data func(r *rune) bool) bool {
	var isChar bool

	for i, r := range *d.expr {
		if !data(&r) {
			isChar = false
			continue
		}

		if isChar {
			*start = rune((*d.expr)[i-1])
			*end = r
			*index = i
			return false
		}

		isChar = true
	}

	return true
}

// !Rune

// isGoodRune checks if r is: (, ), ., ^, π, √, %, *, +, -, /
func isGoodRune(r *rune) bool {
	switch *r {
	case data.Left:
	case data.Right:
	case data.Dot:
	case data.Pow:
	case data.Pi:
	case data.Root:
	default:
		return data.IsOperator(r)
	}
	return true
}

// isGoodChar checks if r is: 0-9, (, ), ., ^, π, √, %, *, +, -, /
func isGoodChar(r *rune) bool {
	if !data.IsNumber(r) {
		return isGoodRune(r)
	}
	return true
}

// isGoodFirstChar checks if first r is: (, ., √, π, √, 0-9
func isGoodFirstChar(r *rune) bool {
	switch *r {
	case data.Left:
	case data.Dot:
	case data.Pi:
	case data.Root:
	default:
		return data.IsNumber(r)
	}
	return true
}

// isGoodLastChar checks if last r is: ), π, 0-9
func isGoodLastChar(r *rune) bool {
	switch *r {
	case data.Right:
	case data.Pi:
	default:
		return data.IsNumber(r)
	}
	return true
}

// isGoodAfter checks if after is: π, (, √, ., 0-9
func isGoodAfter(after *rune) bool {
	switch *after {
	case data.Pi:
	case data.Left:
	case data.Root:
	case data.Dot:
	default:
		return data.IsNumber(after)
	}
	return true
}

// isGoodAfterPow checks if after is: π, (, √, ., 0-9, +, -
func isGoodAfterPow(after *rune) bool {
	if !isGoodAfter(after) {
		return data.IsMoreLess(after)
	}
	return true
}
