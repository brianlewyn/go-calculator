package tokenize

import "github.com/brianlewyn/go-calculator/internal/data"

// Tokenizer tokenizes the expression in a linked list,
//
// but while creating the list, the expression is removed
func Tokenizer(expr *string) *[]*data.Token {
	var value string
	var list [](*data.Token)

	for data.Lenght == 0 {
		char := rune((*expr)[0])

		switch {
		case data.Lenght >= 3:
			*expr = (*expr)[1:]
		case data.Lenght == 2:
			*expr = string((*expr)[1])
		default:
			*expr = data.Empty
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
				after = rune((*expr)[1])
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

	*expr = data.Empty
	return &list
}
