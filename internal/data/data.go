package data

type Kind uint8

// Kind is the data type of the result
const (
	_ = Kind(iota)
	Int08
	Int16
	Int32
	Int64
	Float32
	Float64
)

// DigitLimit is the limit of digits of a float64 and dot
const DigitLimit uint16 = 617

// It is the limit of digits of a data type
const (
	F64 uint16 = 308
	F32 uint8  = 38
	I64 uint8  = 18
	I32 uint8  = 10
	I16 uint8  = 5
	I08 uint8  = 3
)

const (
	// Opetator

	Mod rune = '%' // Mod = '%'
	Mul rune = '*' // Mul = '*'
	Add rune = '+' // Add = '+'
	Sub rune = '-' // Sub = '-'
	Div rune = '/' // Div = '/'
)

const (
	// Parentheses

	Left  rune = '(' // Left  = '('
	Right rune = ')' // Right = ')'
)

const (
	// Special character

	Gap  rune = ' ' // Gap  = '\0'
	Dot  rune = '.' // Dot  = '.'
	Pow  rune = '^' // Pow  = '^'
	Pi   rune = 'π' // Pi   = 'π'
	Root rune = '√' // Root = '√'
)

// IsNumber is number from 0 to 9
func IsNumber(r *rune) bool {
	return '0' <= *r && *r <= '9'
}

// IsRune is symbol or operator of a math expression
func IsRune(r *rune) bool {
	for _, s := range [11]rune{
		Mod, Left, Right, Mul, Add,
		Sub, Dot, Div, Pow, Pi, Root,
	} {
		if *r == s {
			return true
		}
	}
	return false
}

// IsFirstChar is the possible first character of a math expression
func IsFirstChar(r *rune) bool {
	switch {
	case IsNumber(r):
	case *r == Left:
	case *r == Add:
	case *r == Sub:
	case *r == Dot:
	case *r == Pi:
	case *r == Root:
	default:
		return false
	}
	return true
}

// IsLastChar is the possible last character of a mathematical expression
func IsLastChar(r *rune) bool {
	return IsNumber(r) || *r == Right
}
