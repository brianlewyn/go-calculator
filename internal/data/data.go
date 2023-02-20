package data

// DigitLimit is the limit of digits of a float64 and dot
const DigitLimit uint16 = 617

const ( // Assistant Characthers
	Jocker rune   = '#'
	Empty  string = ""
	Zero   string = "0"
)

var (
	Error  error   // Error is an error global
	Answer float64 // Answer is the calculator answer
	Lenght int     // Lenght is the lenght of the string expression
)
