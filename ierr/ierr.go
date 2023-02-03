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

func (e Rune) Error() string {
	return fmt.Sprintf("char=%q index=%d", *e.R, *e.I)
}

// !Add context to the data error

func (r Rune) Character() error {
	return wrap(Syntax,
		wrap(IncorrectCharacter,
			&Rune{R: r.R, I: r.I},
		),
	)
}

func (r Rune) Dot() error {
	return wrap(Expression,
		wrap(DuplicateDot,
			&Rune{R: r.R, I: r.I},
		),
	)
}

func (r Rune) Digit() error {
	return wrap(Expression,
		wrap(DigitLimit,
			&Rune{R: r.R, I: r.I},
		),
	)
}

// Add a wrapper of type error to the already created error
func wrap(kind kind, err error) error {
	return fmt.Errorf("%s: %w", kind, err)
}
