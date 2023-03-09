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

const ( // Others
	Gap rune = ' ' // Gap = '\0'
	Dot rune = '.' // Dot = '.'
	Pow rune = '^' // Power = '^'
)

const ( // Pi
	Pi       rune  = 'π' // Pi Number = 'π'
	PrefixPi rune  = 207 // The prefix of Pi
	PiLenght uint8 = 2   // Pi lenght
)

const ( // Root
	Root       rune  = '√' // Square Root = '√'
	PrefixRoot rune  = 226 // The prefix of Root
	RootLenght uint8 = 3   // Root lenght
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

// IsPrefixPi returns true if r is a Pi number
func IsPrefixPi(r *rune) bool { return *r == PrefixPi }

// IsRoot returns true if r is a square root
func IsRoot(r *rune) bool { return *r == Root }

// IsPrefixRoot returns true if r is a square root
func IsPrefixRoot(r *rune) bool { return *r == PrefixRoot }

// !For each rune group

// IsPrefix returns true if r is a prefix:
//
// π, √
func IsPrefix(r *rune) bool {
	return *r == PrefixPi || *r == PrefixRoot
}

// IsNumber returns true if r is:
//
// 0-9
func IsNumber(r *rune) bool {
	return '0' <= *r && *r <= '9'
}

// IsFloat retunrs true if r is:
//
// 0-9, .
func IsFloat(r *rune) bool {
	if !IsNumber(r) {
		return IsDot(r)
	}
	return true
}

// IsMoreLess returns true if r is:
//
// +, -
func IsMoreLess(r *rune) bool {
	if !IsAdd(r) {
		return IsSub(r)
	}
	return true
}
