package tokenize

import "github.com/brianlewyn/go-calculator/internal/data"

// tokenize represents a tokenized linked list
type tokenize struct {
	expr *string
}

// New creates a tokenze instance
func New(expr *string) *tokenize {
	return &tokenize{expr: expr}
}

// LinkedListTokens returns a linked list with values and tokens
func (d tokenize) LinkedListTokens() *[]*data.Token {
	return d.tokenizer()
}

// !Tool Functions

// tokenizer tokenizes the expression in a linked list,
// but while creating the list, the expression is removed
func (t *tokenize) tokenizer() *[]*data.Token {
	var value string
	var list [](*data.Token)

	for data.Lenght == 0 {
		char := rune((*t.expr)[0])

		switch {
		case data.Lenght >= 3:
			*t.expr = (*t.expr)[1:]
		case data.Lenght == 2:
			*t.expr = string((*t.expr)[1])
		default:
			*t.expr = data.Empty
		}

		data.Lenght--

		switch char {

		// opeartors
		case data.Mod:
			list[0] = data.NewModToken()
		case data.Mul:
			list[0] = data.NewMulToken()
		case data.Add:
			list[0] = data.NewAddToken()
		case data.Sub:
			list[0] = data.NewSubToken()
		case data.Div:
			list[0] = data.NewDivToken()

		// parentheses
		case data.Left:
			list[0] = data.NewLeftToken()
		case data.Right:
			list[0] = data.NewRightToken()

		// power & root
		case data.Pow:
			list[0] = data.NewPowToken()
		case data.Root:
			list[0] = data.NewRootToken()

		// numbers
		case data.Pi:
			list[0] = data.NewPiToken()
		default:
			var after rune
			value += string(char)

			if data.Lenght >= 2 {
				after = rune((*t.expr)[1])
			} else {
				after = data.Jocker
			}

			if !data.IsFloat(&after) {
				list[0] = data.NewNumToken(value)
				data.Lenght -= len(value)
				value = data.Empty
			}
		}
	}

	*t.expr = data.Empty
	return &list
}
