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

// It is the limit of digits of a data type
const (
	F64 uint16 = 309
	F32 uint8  = 39
	I64 uint8  = 18
	I32 uint8  = 10
	I16 uint8  = 5
	I08 uint8  = 3
)

// Opetator
const (
	Mul rune = '*'
	Add rune = '+'
	Sub rune = '-'
	Div rune = '/'
	Mod rune = '%'
)

// Parentheses
const (
	Left  rune = '('
	Right rune = ')'
)

// Special character
const (
	Gap  rune = ' '
	Dot  rune = '.'
	Pow  rune = '^'
	Pi   rune = 'π'
	Root rune = '√'
)

// Numbers are numbers from 0 to 9
func Numbers(r *rune) bool {
	return '0' <= *r && *r <= '9'
}

// Runes are all characters of a math expression
func Runes(r *rune) bool {
	for _, s := range [12]rune{
		Gap, Mod, Left, Right, Mul, Add,
		Sub, Dot, Div, Pow, Pi, Root,
	} {
		if *r == s {
			return true
		}
	}
	return false
}
