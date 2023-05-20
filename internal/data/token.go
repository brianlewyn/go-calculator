package data

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
