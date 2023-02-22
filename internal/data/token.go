package data

// TokenKind is the data type of the value
type TokenKind uint8

// Token represents a token from the list
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

	PowToken  // Pow = '^'
	RootToken // Root = '√'

	PiToken  // Pi number = 'π'
	NumToken // Number = n
)

// NewModToken returns a Token with ModToken kind
func NewModToken() *Token { return &Token{Kind: ModToken} }

// NewMulToken returns a Token with MulToken kind
func NewMulToken() *Token { return &Token{Kind: MulToken} }

// NewAddToken returns a Token with AddToken kind
func NewAddToken() *Token { return &Token{Kind: AddToken} }

// NewSubToken returns a Token with SubToken kind
func NewSubToken() *Token { return &Token{Kind: SubToken} }

// NewDivToken returns a Token with DivToken kind
func NewDivToken() *Token { return &Token{Kind: DivToken} }

// NewLeftToken returns a Token with LeftToken kind
func NewLeftToken() *Token { return &Token{Kind: LeftToken} }

// NewRightToken returns a Token with RightToken kind
func NewRightToken() *Token { return &Token{Kind: RightToken} }

// NewPowToken returns a Token with PowToken kind
func NewPowToken() *Token { return &Token{Kind: PowToken} }

// NewRootToken returns a Token with RootToken kind
func NewRootToken() *Token { return &Token{Kind: RootToken} }

// NewPiToken returns a Token with PiToken kind
func NewPiToken() *Token { return &Token{Kind: PiToken} }

// NewNumToken returns a Token with NumToken kind
func NewNumToken(value string) *Token {
	return &Token{Kind: NumToken, Value: value}
}
