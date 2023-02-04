package data

// !Runes

const ( // Operators
	Mod rune = '%' // Module = '%'
	Mul rune = '*' // Multiplication = '*'
	Add rune = '+' // Addition = '+'
	Sub rune = '-' // Subtraction = '-'
	Div rune = '/' // Division = '/'
)

const ( // Parenthese
	Left  rune = '(' // Left Parentheses = '('
	Right rune = ')' // Right Parentheses = ')'
)

const ( // Special Characthers
	Gap  rune = ' ' // Gap = '\0'
	Dot  rune = '.' // Dot = '.'
	Pow  rune = '^' // Power = '^'
	Pi   rune = 'π' // Pi Number = 'π'
	Root rune = '√' // Square Root = '√'
)

// !For each rune

// IsMod returns true if r is a mudule operator
func IsMod(r *rune) bool { return *r == Mod }

// IsMul returns true if r is a multiplication operator
func IsMul(r *rune) bool { return *r == Mul }

// IsAdd returns true if r is a addition operator
func IsAdd(r *rune) bool { return *r == Add }

// IsSub returns true if r is a subtraction operator
func IsSub(r *rune) bool { return *r == Sub }

// IsDiv returns true if r is a division operator
func IsDiv(r *rune) bool { return *r == Div }

// IsLeft returns true if r is a left parentheses operator
func IsLeft(r *rune) bool { return *r == Left }

// IsRight returns true if r is a right parentheses operator
func IsRight(r *rune) bool { return *r == Right }

// IsGap returns true if r is a gap
func IsGap(r *rune) bool { return *r == Gap }

// IsDot returns true if r is a dot
func IsDot(r *rune) bool { return *r == Dot }

// IsPow returns true if r is a power
func IsPow(r *rune) bool { return *r == Pow }

// IsPi returns true if r is a Pi number
func IsPi(r *rune) bool { return *r == Pi }

// IsRoot returns true if r is a square root
func IsRoot(r *rune) bool { return *r == Root }

// IsNumber returns true if r is a number from 0 to 9
func IsNumber(r *rune) bool {
	return '0' <= *r && *r <= '9'
}

// !For each rune group

// IsOperator is operator: %, *, +, -, /
func IsOperator(r *rune) bool {
	switch *r {
	case Mod:
	case Mul:
	case Add:
	case Sub:
	default:
		return IsDiv(r)
	}
	return true
}

// IsRune is symbol or operator of a math expression
func IsRune(r *rune) bool {
	switch *r {
	case Left:
	case Right:
	case Dot:
	case Pow:
	case Pi:
	case Root:
	default:
		return IsOperator(r)
	}
	return true
}

// IsFirstChar is the possible first character of a math expression
func IsFirstChar(r *rune) bool {
	switch *r {
	case Left:
	case Add:
	case Sub:
	case Dot:
	case Pi:
	case Root:
	default:
		return IsNumber(r)
	}
	return true
}

// IsLastChar is the possible last character of a mathematical expression
func IsLastChar(r *rune) bool {
	if !IsRight(r) {
		return IsNumber(r)
	}
	return true
}