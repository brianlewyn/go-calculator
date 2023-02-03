package ierr

import "fmt"

type kind string

// !What kind of bug?
const (
	Syntax     = kind("syntax error")
	Expression = kind("expression error")
)

// !What error occurred?
const (
	IncorrectCharacter = kind("incorrect character")
	DuplicateDot       = kind("this dot is duplicated")
	DigitLimit         = kind("this digit exceeds the digit limit")
)

// !Interface errors

type Rune struct {
	R *rune // character
	I *int  // index
}

// !The data error

func (s Rune) Error() string {
	return fmt.Sprintf("%d", *s.I)
}

// !Add context to the data error

func (r Rune) Character() error {
	return wrapRune(Syntax, IncorrectCharacter, r.R, &Rune{I: r.I})
}

func (r Rune) Dot() error {
	return wrapRune(Expression, DuplicateDot, r.R, &Rune{I: r.I})
}

func (r Rune) Digit() error {
	return wrapRune(Expression, DigitLimit, r.R, &Rune{I: r.I})
}

// wrap add a wrapper of type error to the already created error
func wrap(kind kind, err error) error {
	return fmt.Errorf("%s: %w", kind, err)
}

// wrapRune works mainly for the Rune interface.
// Add three wrappers of type error to the already created error
func wrapRune(k1, k2 kind, r *rune, err error) error {
	return wrap(k1, wrap(k2, fmt.Errorf("char=%q index=%w", *r, err)))
}
