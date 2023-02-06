package data

// TokenKind is the data type of the value
type TokenKind uint8

const (
	_ = TokenKind(iota)

	ModToken // Mod = '%'
	MulToken // Mul = '*'
	AddToken // Add = '+'
	SubToken // Sub = '-'
	DivToken // Div = '/'

	NumToken // Number = n
	PiToken  // Pi number = 'π'

	LeftToken  // Left = '('
	RightToken // Right = ')'

	PowToken  // Pow = '^' Remember: n^±., ^±n, ^±π, ^±(, ^±√,
	RootToken // Root = '√'
)

// DigitLimit is the limit of digits of a float64 and dot
const DigitLimit uint16 = 617
