package ierr

import "fmt"

type kind string

// !What kind of bug?
const (
	Syntax     = kind("syntax error")
	Expression = kind("math expression error")
)

// !What error occurred?
const (
	IncorrectCharacter = kind("incorrect character")
	DuplicateDots      = kind("duplicate dots")
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

func (r Rune) Dots() error {
	return wrap(Expression,
		wrap(DuplicateDots,
			&Rune{R: r.R, I: r.I},
		),
	)
}

// Add a wrapper of type error to the already created error
func wrap(kind kind, err error) error {
	return fmt.Errorf("%s: %w", kind, err)
}
