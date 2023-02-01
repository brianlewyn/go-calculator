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

const Numbers string = "0123456789"

// Opetator
const (
	Mul       string = "*"
	Add       string = "+"
	Sub       string = "-"
	Div       string = "/"
	Mod       string = "%"
	Opetators string = Mul + Add + Sub + Div + Mod
)

// Parentheses
const (
	Left        string = "("
	Right       string = ")"
	Parentheses string = Left + Right
)

// Special character
const (
	Dot      string = "."
	Pow      string = "^"
	Gap      string = " "
	Pi       string = "π"
	Root     string = "√"
	Specials string = Dot + Pow + Gap + Pi + Root
)

// Runes are all characters used
const Runes string = Numbers + Opetators + Parentheses + Specials
