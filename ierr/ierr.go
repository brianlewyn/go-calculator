package ierr

import (
	"fmt"
	"strings"
)

type kind string

// !What kind of bug?
const (
	Syntax     = kind("syntax error")
	Expression = kind("expression error")
)

// !What error occurred?
const (
	IncorrectCharacter = kind("this character is incorrect")
	DigitLimit         = kind("this digit exceeds the digit limit")
	FirstChar          = kind("this character cannot be the beginning")
	LastChar           = kind("this character cannot be the end")
	NotTogether        = kind("these characters cannot be together")
)

// !Interface errors

type Rune struct {
	R *rune // character
	I *int  // index
}

type TwoRune struct {
	S *rune // start
	E *rune // end
	I *int  // index
}

// !The data error

func (r Rune) Error() string {
	return fmt.Sprintf("%d", *r.I)
}

func (t TwoRune) Error() string {
	return fmt.Sprintf("%d", *t.I)
}

// !Add context to the data error

func (r Rune) Character() error {
	return wrapRune(Syntax, IncorrectCharacter, r.R, &Rune{I: r.I})
}

func (r Rune) Limit() error {
	return wrapRune(Expression, DigitLimit, r.R, &Rune{I: r.I})
}

func (r Rune) Start() error {
	return wrapRune(Expression, FirstChar, r.R, &Rune{I: r.I})
}

func (r Rune) Final() error {
	return wrapRune(Expression, LastChar, r.R, &Rune{I: r.I})
}

func (t TwoRune) Together() error {
	return wrapTwoRune(Expression, NotTogether, t.S, t.E, &Rune{I: t.I})
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

// wrapTwoRune works mainly for the TwoRune interface.
// Add three wrappers of type error to the already created error
func wrapTwoRune(k1, k2 kind, s, e *rune, err error) error {
	return wrap(k1, wrap(k2, fmt.Errorf("start=%q end=%q index=%w", *s, *e, err)))
}

// As is similar to errors.As func of standard library
func As(err error, target kind) bool {
	if !strings.Contains(fmt.Sprint(err), string(target)) {
		return false
	}
	return true
}
