package ierr

import (
	"errors"
	"fmt"
	"strings"
)

type kind string

// What kind of main error occurred?
const (
	Syntax     = kind("syntax error")
	Expression = kind("expression error")
)

// What kind of secundary error occurred?
const (
	IncorrectCharacter = kind("this character is incorrect")
	DigitLimit         = kind("this digit exceeds the digit limit")
	FirstChar          = kind("this character cannot be the beginning")
	LastChar           = kind("this character cannot be the end")
	NotTogether        = kind("these characters cannot be together")
)

// What error occurred?
var EmptyField = wrap(Syntax, errors.New("empty field"))

// Interface errors

type OneRune struct {
	R *rune // character
	I *int  // index
}

type TwoRune struct {
	S *rune // start
	E *rune // end
	I *int  // index
}

// The data error

func (r OneRune) Error() string { return fmt.Sprintf("%d", *r.I) }
func (r TwoRune) Error() string { return fmt.Sprintf("%d", *r.I) }

// Add context to the data error

func (r OneRune) Character() error {
	return wrapOneRune(Syntax, IncorrectCharacter, r.R, &OneRune{I: r.I})
}

func (r OneRune) Limit() error {
	return wrapOneRune(Expression, DigitLimit, r.R, &OneRune{I: r.I})
}

func (r OneRune) Start() error {
	return wrapOneRune(Expression, FirstChar, r.R, &OneRune{I: r.I})
}

func (r OneRune) Final() error {
	return wrapOneRune(Expression, LastChar, r.R, &OneRune{I: r.I})
}

func (r TwoRune) Together() error {
	return wrapTwoOneRune(Expression, NotTogether, r.S, r.E, &OneRune{I: r.I})
}

// wrap add a wrapper of type error to the already created error
func wrap(kind kind, err error) error {
	return fmt.Errorf("%s: %w", kind, err)
}

// wrapOneRune works mainly for the OneRune interface.
// Add three wrappers of type error to the already created error
func wrapOneRune(k1, k2 kind, r *rune, err error) error {
	return wrap(k1, wrap(k2, fmt.Errorf("char=%q index=%w", *r, err)))
}

// wrapTwoOneRune works mainly for the TwoOneRune interface.
// Add three wrappers of type error to the already created error
func wrapTwoOneRune(k1, k2 kind, s, e *rune, err error) error {
	return wrap(k1, wrap(k2, fmt.Errorf("start=%q end=%q index=%w", *s, *e, err)))
}

// As is similar to errors.As func of standard library
func As(err error, target kind) bool {
	if !strings.Contains(fmt.Sprint(err), string(target)) {
		return false
	}
	return true
}
