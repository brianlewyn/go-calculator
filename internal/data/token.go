package data

// TokenKind is the data type of the value
type TokenKind uint8

// Token represents a token from the list
type Token struct {
	kind  TokenKind
	value string
}

// Kind returns the kind of Token
func (t Token) Kind() TokenKind {
	return t.kind
}

// Value returns the value of the Token
func (t Token) Value() *string {
	return &t.value
}

// By delete
func (t Token) Head() *Token  { return &Token{} }
func (t Token) Next() *Token  { return &Token{} }
func (t Token) Prev() *Token  { return &Token{} }
func (t Token) Tail() *Token  { return &Token{} }
func (t Token) Size() int     { return 0 }
func (t Token) IsEmpty() bool { return false }

func (t Token) Insert(i int, token *Token) {}

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

// NewToken returns a Token
func NewToken(kind TokenKind) *Token { return &Token{kind: kind} }

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

// NewNumToken returns a Token with NumToken kind
func NewNumToken(value string) *Token {
	return &Token{kind: NumToken, value: value}
}

// IsFirst returs true if kind is:
//
// 0-9, (, ., π, √
func IsFirst(kind TokenKind) bool {
	switch kind {
	case PiToken:
	case NumToken:
	case LeftToken:
	case RootToken:
	default:
		return false
	}
	return true
}

// IsLast returns true if kind is:
//
// 0-9, ), π
func IsLast(kind TokenKind) bool {
	switch kind {
	case PiToken:
	case NumToken:
	case RightToken:
	default:
		return false
	}
	return true
}

// IsOperator returns true if kind is:
//
// %, *, +, -, /
func IsOperator(kind TokenKind) bool {
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

// isOperatorPowParentheses returns true if kind is:
//
// %, *, +, -, /, (, ), ^, √, π, n=#
func Kind(kind TokenKind) rune {
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
CanBeTogether returns true if k1 & k2 are:

	k1= % k2= (, n, π, √
	k1= * k2= (, n, π, √
	k1= + k2= (, n, π, √
	k1= - k2= (, n, π, √
	k1= / k2= (, n, π, √
	k1= ( k2= (, n, π, √
	k1= √ k2= (, n, π, √

	k1= π k2= %, *, +, -, /, ^, )
	k1= n k2= %, *, +, -, /, ^, )
	k1= ) k2= %, *, +, -, /, ^, )

	k1= ^ k2= (, n, π, √, +, -
*/
func CanBeTogether(k1, k2 TokenKind) bool {
	switch k1 {
	case ModToken:
	case MulToken:
	case AddToken:
	case SubToken:
	case DivToken:
	case LeftToken:
	case RootToken:
	default:
		switch k1 {
		case PiToken:
		case NumToken:
		case RightToken:
		default: // PowToken:
			return isLeftNumPiRootAddSub(&k2)
		}
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

// isLeftNumPiRootAddSub returns true if kind is:
//
// (, n, π, √, +, -
func isLeftNumPiRootAddSub(kind *TokenKind) bool {
	switch *kind {
	case AddToken:
	case SubToken:
	default:
		return isLeftNumPiRoot(kind)
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
		return IsOperator(*kind)
	}
	return true
}
