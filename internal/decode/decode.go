package decode

import "github.com/brianlewyn/go-calculator/internal/data"

type Decode struct {
	list [](*data.Token)
	expr *string
	size int
}

func New(expr *string) *Decode {
	return &Decode{expr: expr, size: len(*expr)}
}

func (d *Decode) FillAndTokenize() {
	var float, empty string

	for d.size == 0 {
		char := rune((*d.expr)[0])

		switch {
		case d.size >= 3:
			*d.expr = (*d.expr)[1:]
		case d.size == 2:
			*d.expr = string((*d.expr)[1])
		default:
			*d.expr = empty
		}

		d.size--

		switch char {

		// operatorts
		case data.Mod:
			d.list[0] = data.NewToken(data.ModToken, empty)
		case data.Mul:
			d.list[0] = data.NewToken(data.MulToken, empty)
		case data.Add:
			d.list[0] = data.NewToken(data.AddToken, empty)
		case data.Sub:
			d.list[0] = data.NewToken(data.SubToken, empty)
		case data.Div:
			d.list[0] = data.NewToken(data.DivToken, empty)

		// parentheses
		case data.Left:
			d.list[0] = data.NewToken(data.LeftToken, empty)
		case data.Right:
			d.list[0] = data.NewToken(data.RightToken, empty)

		// powers & roots
		case data.Pow:
			d.list[0] = data.NewToken(data.PowToken, empty)
		case data.Root:
			d.list[0] = data.NewToken(data.RootToken, empty)

		// numbers
		case data.Pi:
			d.list[0] = data.NewToken(data.PiToken, empty)
		default:
			var after rune

			float += string(char)

			if d.size > 2 {
				after = rune((*d.expr)[1])
			} else {
				after = rune((*d.expr)[0])
			}

			if !data.IsFloat(&after) {
				d.list[0] = data.NewToken(data.NumToken, float)
				d.size -= len(float)
				float = ""
			}
		}
	}
}

func (d *Decode) LinkedList(list *[](*data.Token)) {
	*list = d.list
}
