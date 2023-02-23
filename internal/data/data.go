package data

// DigitLimit is the limit of digits of a float64 and dot
const DigitLimit uint16 = 617

const ( // Assistant Characthers
	Jocker rune   = '#'
	Empty  string = ""
)

var (
	Lenght     int     // Lenght is the lenght of the string expression
	Expression *string // Expression is basic math expression
	Answer     float64 // Answer is the calculator answer
	Error      error   // Error is a possible error
)
