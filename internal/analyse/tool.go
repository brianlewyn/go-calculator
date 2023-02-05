package analyse

import "github.com/brianlewyn/go-calculator/internal/data"

// duplicate

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

// is rune methods

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

func isGoodChar(r *rune) bool {
	if !data.IsNumber(r) {
		return isGoodRune(r)
	}
	return true
}

func isGoodFirstChar(r *rune) bool {
	switch *r {
	case data.Left:
	case data.Add:
	case data.Sub:
	case data.Dot:
	case data.Pi:
	case data.Root:
	default:
		return data.IsNumber(r)
	}
	return true
}

func isGoodLastChar(r *rune) bool {
	if !data.IsRight(r) {
		return data.IsNumber(r)
	}
	return true
}
