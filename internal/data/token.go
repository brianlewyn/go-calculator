package data

// TokenKind is the data type of the Token
type TokenKind uint8

// Token represents a token from the list
type Token interface {
	Kind() TokenKind
}

type Kind struct {
	kind TokenKind
}

// Kind returns the kind of Kind Token
func (k Kind) Kind() TokenKind {
	return k.kind
}

type Number struct {
	kind  TokenKind
	value string
}

// Kind returns the kind of Number Token
func (n Number) Kind() TokenKind {
	return n.kind
}

// Value returns the value of the Number Token
func (n Number) Value() string {
	return n.value
}

type Float struct {
	kind  TokenKind
	value float64
}

// Kind returns the kind of Float Token
func (f Float) Kind() TokenKind {
	return f.kind
}

// Value returns the value of the Float Token
func (f Float) Value() float64 {
	return f.value
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

// NewToken returns a Kind Token
func NewToken(kind TokenKind) *Token {
	var token Token = Kind{kind: kind}
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

// NewFloatToken returns a Float Token
func NewFloatToken(value float64) *Token {
	var token Token = Float{kind: NumToken, value: value}
	return &token
}

// !For each token group

// IsAddSubToken returns true if r is:
//
// +, -
func IsAddSubToken(kind TokenKind) bool {
	if kind != AddToken {
		return kind == SubToken
	}
	return true
}

// IsAddSubToken returns true if r is:
//
// n, π
func IsNumPiToken(kind TokenKind) bool {
	if kind != NumToken {
		return kind == PiToken
	}
	return true
}

// IsFirstToken returs true if kind is:
//
// √, (, π, n
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
// ), π, n
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
// %, *, +, -, /
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
// %, *, +, -, /, ^, √
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
// %, *, +, -, /, (, ), ^, √, π, n=#
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
// (, n, π, √
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
// %, *, +, -, /, ^, )
func isOperatorPowRight(kind *TokenKind) bool {
	switch *kind {
	case PowToken:
	case RightToken:
	default:
		return IsOperatorToken(*kind)
	}
	return true
}
