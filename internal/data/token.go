package data

// TokenKind is the data type of the Token
type TokenKind uint8

// Token represents both a token Symbol and a token Number from the list
type Token interface {
	Kind() TokenKind
}

// Symbol represents a token symbol from the list
type Symbol struct {
	kind TokenKind
}

// Number represents a number token from the list
type Number struct {
	kind  TokenKind
	value string
}

// Decimal represents a decimal number token from the list
type Decimal struct {
	kind  TokenKind
	value float64
}

// NewSymbolToken returns a token Symbol
func NewSymbolToken(kind TokenKind) Token {
	return Symbol{kind: kind}
}

// NewNumberToken returns a token Number
func NewNumberToken(value string) Token {
	return Number{kind: NumToken, value: value}
}

// NewDecimalToken returns a token Decimal
func NewDecimalToken(value float64) Token {
	return Decimal{kind: NumToken, value: value}
}

// Kind returns the token Symbol type
func (s Symbol) Kind() TokenKind { return s.kind }

// Kind returns the token Number type
func (n Number) Kind() TokenKind { return n.kind }

// Kind returns the token Decimal type
func (d Decimal) Kind() TokenKind { return d.kind }

// Value returns the token Number value
func (n Number) Value() string { return n.value }

// Value returns the token Decimal value
func (d Decimal) Value() float64 { return d.value }

const (
	_ = TokenKind(iota)

	ModToken   // Mod = '%'
	MulToken   // Mul = '*'
	AddToken   // Add = '+'
	SubToken   // Sub = '-'
	DivToken   // Div = '/'
	LeftToken  // Left = '('
	RightToken // Right = ')'
	PowToken   // Pow = '^'
	RootToken  // Root = '√'
	PiToken    // Pi number = 'π'

	NumToken // Number = n
)

// !For each TokenKind

// TokenKindMap represent the follow kinds:
//
//	%, *, +, -, /, (, ), ^, √   π
//	1  2  3  4  5  6  7  8  9  10
var TokenKindMap = map[rune]TokenKind{
	Mod:   ModToken,
	Mul:   MulToken,
	Add:   AddToken,
	Sub:   SubToken,
	Div:   DivToken,
	Left:  LeftToken,
	Right: RightToken,
	Pow:   PowToken,
	Root:  RootToken,
	Pi:    PiToken,
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
		return isOperatorPowRight(k2)
	}
	return isLeftNumPiRoot(k2)
}

// isOperatorPowRight returns true if kind is:
//
//	%, *, +, -, /, ^, )
func isOperatorPowRight(kind TokenKind) bool {
	switch kind {
	case PowToken:
	case RightToken:
	default:
		return IsOperatorToken(kind)
	}
	return true
}

// isLeftNumPiRoot returns true if kind is:
//
//	(, n, π, √
func isLeftNumPiRoot(kind TokenKind) bool {
	switch kind {
	case LeftToken:
	case NumToken:
	case PiToken:
	case RootToken:
	default:
		return false
	}
	return true
}
