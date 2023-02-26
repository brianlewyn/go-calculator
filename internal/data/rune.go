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

const ( // Special Pi
	Pi      rune = 'π' // Pi Number = 'π'
	PiFirst rune = 207 // Firt sub rune: Pi Number
	PiLast  rune = 128 // Last sub rune: Pi Number
)

const ( // Special Root
	Root       rune = '√' // Square Root = '√'
	RootFirst  rune = 226 // Firt sub rune: Square Root
	RootSecond rune = 136 // Second sub rune: Square Root
	RootLast   rune = 154 // Last sub rune: Square Root
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

// IsSpecial returns true if r is a initial of π or √
func IsSpecial(r *rune) bool {
	return *r == PiFirst || *r == RootFirst
}

// IsPi returns true if r1 & r2 is a initial of π
func IsPi(r1, r2 *rune) bool {
	return *r1 == PiFirst && *r2 == PiLast
}

// IsRoot returns true if r1, r2 & r3 is a initial of √
func IsRoot(r1, r2, r3 *rune) bool {
	return *r1 == RootFirst && *r2 == RootSecond && *r3 == RootLast
}

// !For each rune group

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
