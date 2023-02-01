package data

type Kind int8

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

// Tnn is the limit of digits of a data type
const (
	T64 int8 = 18
	T32 int8 = 10
	T16 int8 = 5
	T08 int8 = 3
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

// Runes are all characters used
var Runes = [22]rune{
	Gap, Mod, Left, Right, // 4
	'0', '1', '2', '3', '4', // 5
	'5', '6', '7', '8', '9', // 5
	Mul, Add, Sub, Dot, Div, // 5
	Pow, Pi, Root, //3
}
