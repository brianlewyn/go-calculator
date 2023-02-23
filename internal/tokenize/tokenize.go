package tokenize

import "github.com/brianlewyn/go-calculator/internal/data"

// Tokenizer tokenizes the expression in a linked list,
//
// but while creating the list, the expression is removed
func Tokenizer() *[]*data.Token {
	var value string
	var list [](*data.Token)

	for data.Lenght > 0 {
		first := rune((*data.Expression)[0])

		switch first {

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
			if data.IsFloat(&first) {
				var after rune
				value += string(first)

				if data.Lenght >= 2 {
					after = rune((*data.Expression)[1])
				} else {
					after = data.Jocker
				}

				if !data.IsFloat(&after) {
					list[0] = data.NewNumToken(value)
					value = data.Empty
				}
			}
		}

		*data.Expression = (*data.Expression)[1:]
		data.Lenght--
	}

	*data.Expression = data.Empty
	return &list
}
