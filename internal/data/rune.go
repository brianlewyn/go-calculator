package data

// !Runes

const (
	Mod   rune = '%' // Module = '%'
	Mul   rune = '*' // Multiplication = '*'
	Add   rune = '+' // Addition = '+'
	Sub   rune = '-' // Subtraction = '-'
	Div   rune = '/' // Division = '/'
	Left  rune = '(' // Left Parentheses = '('
	Right rune = ')' // Right Parentheses = ')'
	Pow   rune = '^' // Power = '^'
	Root  rune = '√' // Square Root = '√'

	Pi  rune = 'π' // Pi Number = 'π'
	Dot rune = '.' // Dot = '.'
	Num rune = 'n' // Num = 'n'

	Gap rune = ' ' // Gap = ' '
)

// !For each rune

// RuneMap represent the follow symbols:
//
//	1  2  3  4  5  6  7  8  9  10  11
//	%, *, +, -, /, (, ), ^, √,  π,  n
var RuneMap = map[TokenKind]rune{
	ModToken:   Mod,
	MulToken:   Mul,
	AddToken:   Add,
	SubToken:   Sub,
	DivToken:   Div,
	LeftToken:  Left,
	RightToken: Right,
	PowToken:   Pow,
	RootToken:  Root,
	PiToken:    Pi,
	NumToken:   Num,
}

// !For each rune group

// IsNumber returns true if r is:
//
// 0-9
func IsNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

// IsDecimal retunrs true if r is:
//
// 0-9, .
func IsDecimal(r rune) bool {
	return IsNumber(r) || Dot == r
}
