package ierr

import (
	"errors"
	"fmt"
	"strings"
)

type kind string

// What kind of main error occurred?
const Syntax = kind("syntax error")

// What kind of context error occurred?
const (
	Rune_Unknown      = kind("this is an unknown rune")
	Number_Misspelled = kind("this is a misspelled number")
	Number_Limit      = kind("this number exceeds the digit limit")
	Kind_Together     = kind("these data types cannot be together")
)

// What error occurred?
var (
	EmptyField      = wrap(Syntax, errors.New("empty field"))
	StartKind       = wrap(Syntax, errors.New("the start of the expression can't be the beginning"))
	EndKind         = wrap(Syntax, errors.New("the end of the expression can't be the end"))
	IncompleteLeft  = wrap(Syntax, errors.New("there are incomplete left parentheses"))
	IncompleteRight = wrap(Syntax, errors.New("there are incomplete right parentheses"))
)

// Interface errors

type Rune struct{ r rune }
type Number struct{ n string }
type Kind struct{ k1, k2 string }

// Functions to create an instance with New

func NewRune(r rune) *Rune {
	return &Rune{r: r}
}
func NewNumber(n string) *Number {
	return &Number{n: n}
}
func NewKind(k1, k2 string) *Kind {
	return &Kind{k1: k1, k2: k2}
}

// The data error

func (r Rune) Error() string {
	return string(r.r)
}
func (n Number) Error() string {
	return n.n
}
func (k Kind) Error() string {
	return fmt.Sprintf("%s:%s", k.k1, k.k2)
}

// Add context to the data error

// Unknown returns an error with the kind of context: Rune_Unknown
func (r Rune) Unknown() error {
	return doubleWrap(Syntax, Rune_Unknown, NewRune(r.r))
}

// Misspelled returns an error with the kind of context: Number_Misspelled
func (n Number) Misspelled() error {
	return doubleWrap(Syntax, Number_Misspelled, NewNumber(n.n))
}

// Limit returns an error with the kind of context: Number_Limit
func (n Number) Limit() error {
	return doubleWrap(Syntax, Number_Limit, NewNumber(n.n))
}

// NotTogether returns an error with the kind of context: Kind_Together
func (k Kind) NotTogether() error {
	return doubleWrap(Syntax, Kind_Together, NewKind(k.k1, k.k2))
}

// !Tool Functions

// wrap adds a wrapper of type error to the already created error
func wrap(kind kind, err error) error {
	return fmt.Errorf("%s: %w", kind, err)
}

// doubleWrap adds double wrapper of type error to the already created error
func doubleWrap(k1, k2 kind, err error) error {
	return wrap(k1, wrap(k2, err))
}

// As is similar to errors.As func of standard library
func As(err error, target kind) bool {
	return strings.Contains(fmt.Sprint(err), string(target))
}
