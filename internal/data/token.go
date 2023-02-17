package data

// TokenKind is the data type of the value
type TokenKind uint8

// token
type Token struct {
	Kind  TokenKind
	Value string
}

const (
	_ = TokenKind(iota)

	ModToken // Mod = '%'
	MulToken // Mul = '*'
	AddToken // Add = '+'
	SubToken // Sub = '-'
	DivToken // Div = '/'

	LeftToken  // Left = '('
	RightToken // Right = ')'

	PowToken  // Pow = '^' Remember: n^±., ^±n, ^±π, ^±(, ^±√,
	RootToken // Root = '√'

	PiToken  // Pi number = 'π'
	NumToken // Number = n
)

func NewToken(kind TokenKind, value string) *Token {
	return &Token{Kind: kind, Value: value}
}
