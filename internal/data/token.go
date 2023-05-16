package data

// TokenKind is the data type of the Token
type TokenKind uint8

// Token represents both a token Number as well as a token Symbol from the list
type Token interface {
	Kind() TokenKind
}

// Symbol represents a token symbol from the list
type Symbol struct {
	kind TokenKind
}

// Kind returns the token Symbol type
func (s Symbol) Kind() TokenKind { return s.kind }

// Number represents a number token from the list
type Number struct {
	kind  TokenKind
	value string
}

// Kind returns the token Number type
func (n Number) Kind() TokenKind { return n.kind }

// Value returns the token Number value
func (n Number) Value() string { return n.value }

// Decimal represents
type Decimal struct {
	kind  TokenKind
	value float64
}

// Kind returns the token Decimal type
func (d Decimal) Kind() TokenKind { return d.kind }

// Value returns the token Decimal value
func (d Decimal) Value() float64 { return d.value }

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

// NewToken returns a token Symbol
func NewToken(kind TokenKind) *Token {
	var token Token = Symbol{kind: kind}
	return &token
}

// NewModToken returns a Token with ModToken kind
func NewModToken() *Token { return NewToken(ModToken) }

// NewMulToken returns a Token with MulToken kind
func NewMulToken() *Token { return NewToken(MulToken) }

// NewAddToken returns a Token with AddToken kind
func NewAddToken() *Token { return NewToken(AddToken) }

// NewSubToken returns a Token with SubToken kind
func NewSubToken() *Token { return NewToken(SubToken) }

// NewDivToken returns a Token with DivToken kind
func NewDivToken() *Token { return NewToken(DivToken) }

// NewLeftToken returns a Token with LeftToken kind
func NewLeftToken() *Token { return NewToken(LeftToken) }

// NewRightToken returns a Token with RightToken kind
func NewRightToken() *Token { return NewToken(RightToken) }

// NewPowToken returns a Token with PowToken kind
func NewPowToken() *Token { return NewToken(PowToken) }

// NewRootToken returns a Token with RootToken kind
func NewRootToken() *Token { return NewToken(RootToken) }

// NewPiToken returns a Token with PiToken kind
func NewPiToken() *Token { return NewToken(PiToken) }

// NewNumToken returns a Number Token
func NewNumToken(value string) *Token {
	var token Token = Number{kind: NumToken, value: value}
	return &token
}

// NewDecToken returns a token Decimal
func NewDecToken(value float64) *Token {
	var token Token = Decimal{kind: NumToken, value: value}
	return &token
}

// !For each token group

// IsAddSubToken returns true if r is:
//
//	n, π
func IsNumPiToken(kind TokenKind) bool {
	return kind == NumToken || kind == PiToken
}

// IsFirstToken returs true if kind is:
//
//	√, (, π, n
func IsFirstToken(kind TokenKind) bool {
	switch kind {
	case RootToken:
	case LeftToken:
	case PiToken:
	case NumToken:
	default:
		return false
	}
	return true
}

// IsLastToken returns true if kind is:
//
//	), π, n
func IsLastToken(kind TokenKind) bool {
	switch kind {
	case RightToken:
	case PiToken:
	case NumToken:
	default:
		return false
	}
	return true
}

// IsOperatorToken returns true if kind is:
//
//	%, *, +, -, /
func IsOperatorToken(kind TokenKind) bool {
	switch kind {
	case ModToken:
	case MulToken:
	case AddToken:
	case SubToken:
	case DivToken:
	default:
		return false
	}
	return true
}

// IsSpecialToken returns true if kind is:
//
//	%, *, +, -, /, ^, √
func IsSpecialToken(kind TokenKind) bool {
	switch kind {
	case PowToken:
	case RootToken:
	default:
		return IsOperatorToken(kind)
	}
	return true
}

// ChangeToRune returns a TokenKind:
//
//	%, *, +, -, /, (, ), ^, √, π, n
func ChangeToRune(kind TokenKind) rune {
	switch kind {
	case ModToken:
		return Mod
	case MulToken:
		return Mul
	case AddToken:
		return Add
	case SubToken:
		return Sub
	case DivToken:
		return Div
	case LeftToken:
		return Left
	case RightToken:
		return Right
	case PowToken:
		return Pow
	case RootToken:
		return Root
	case PiToken:
		return Pi
	default: // NumToken
		return Jocker
	}
}

/*
CanTokensBeTogether returns true if k1 & k2 are:

	k1= % k2= (, n, π, √
	k1= * k2= (, n, π, √
	k1= + k2= (, n, π, √
	k1= - k2= (, n, π, √
	k1= / k2= (, n, π, √
	k1= ( k2= (, n, π, √
	k1= ^ k2= (, n, π, √
	k1= √ k2= (, n, π, √

	k1= π k2= %, *, +, -, /, ^, )
	k1= n k2= %, *, +, -, /, ^, )
	k1= ) k2= %, *, +, -, /, ^, )
*/
func CanTokensBeTogether(k1, k2 TokenKind) bool {
	switch k1 {
	case ModToken:
	case MulToken:
	case AddToken:
	case SubToken:
	case DivToken:
	case LeftToken:
	case PowToken:
	case RootToken:
	default: // Token (Pi||Num||Right)
		return isOperatorPowRight(&k2)
	}
	return isLeftNumPiRoot(&k2)
}

// isLeftNumPiRoot returns true if kind is:
//
//	(, n, π, √
func isLeftNumPiRoot(kind *TokenKind) bool {
	switch *kind {
	case LeftToken:
	case NumToken:
	case PiToken:
	case RootToken:
	default:
		return false
	}
	return true
}

// isOperatorPowRight returns true if kind is:
//
//	%, *, +, -, /, ^, )
func isOperatorPowRight(kind *TokenKind) bool {
	switch *kind {
	case PowToken:
	case RightToken:
	default:
		return IsOperatorToken(*kind)
	}
	return true
}
