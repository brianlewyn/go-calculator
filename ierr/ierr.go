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

// What kind of context error occurred?
const (
	IncorrectCharacter = kind("this character is incorrect")
	DigitLimit         = kind("this digit exceeds the digit limit")
	FirstChar          = kind("this character cannot be the beginning")
	LastChar           = kind("this character cannot be the end")
	NotTogether        = kind("these characters cannot be together")
)

// What error occurred?
var (
	EmptyField            = wrap(Syntax, errors.New("empty field"))
	IncompleteParentheses = wrap(Expression, errors.New("there are incomplete parentheses"))
)

// Interface errors

type OneRune struct {
	R rune // character
	I int  // index
}

type TwoRune struct {
	S rune // start
	E rune // end
	I int  // index
}

type ThreeRune struct {
	B rune // before
	M rune // middle
	A rune // after
	I int  // index
}

// The data error

func (r OneRune) Error() string   { return fmt.Sprintf("%d", r.I) }
func (r TwoRune) Error() string   { return fmt.Sprintf("%d", r.I) }
func (r ThreeRune) Error() string { return fmt.Sprintf("%d", r.I) }

// Add context to the data error

// Character returns an error with the kind of context: IncorrectCharacter
func (r OneRune) Character() error {
	return wOneRune(Syntax, IncorrectCharacter, r.R, &OneRune{I: r.I})
}

// Limit returns an error with the kind of context: DigitLimit
func (r OneRune) Limit() error {
	return wOneRune(Expression, DigitLimit, r.R, &OneRune{I: r.I})
}

// Start returns an error with the kind of context: FirstChar
func (r OneRune) Start() error {
	return wOneRune(Expression, FirstChar, r.R, &OneRune{I: r.I})
}

// Final returns an error with the kind of context: LastChar
func (r OneRune) Final() error {
	return wOneRune(Expression, LastChar, r.R, &OneRune{I: r.I})
}

// Together returns an error with the kind of context: NotTogether
func (r TwoRune) Together() error {
	return wTwoRune(Expression, NotTogether, r.S, r.E, &TwoRune{I: r.I})
}

// Together returns an error with the kind of context: NotTogether
func (r ThreeRune) Together() error {
	return wThreeRune(Expression, NotTogether, r.B, r.M, r.A, &ThreeRune{I: r.I})
}

// !Tool Functions

// wrap add a wrapper of type error to the already created error
func wrap(kind kind, err error) error {
	return fmt.Errorf("%s: %w", kind, err)
}

// Add three wrappers of type error to the already created error

// wOneRune works mainly for the OneRune interface
func wOneRune(k1, k2 kind, r rune, err error) error {
	return wrap(k1, wrap(k2,
		fmt.Errorf("char=%q index=%w", r, err),
	))
}

// wTwoRune works mainly for the TwoOneRune interface
func wTwoRune(k1, k2 kind, s, e rune, err error) error {
	return wrap(k1, wrap(k2,
		fmt.Errorf("start=%q end=%q index=%w", s, e, err),
	))
}

// wThreeRune works mainly for the TwoOneRune interface
func wThreeRune(k1, k2 kind, b, m, a rune, err error) error {
	return wrap(k1, wrap(k2,
		fmt.Errorf("before=%q middle=%q after=%q index=%w", b, m, a, err),
	))
}

// As is similar to errors.As func of standard library
func As(err error, target kind) bool {
	return strings.Contains(fmt.Sprint(err), string(target))
}
