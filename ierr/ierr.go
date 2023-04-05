package ierr

import (
	"errors"
	"fmt"
	"strings"
)

// KindOf represents the type of context error
type KindOf string

// What kind of main error occurred?
const (
	Syntax = KindOf("syntax error")
	Math   = KindOf("math error")
)

// What kind of context error occurred?
const (
	Rune_Unknown      = KindOf("this is an unknown rune")
	Number_Misspelled = KindOf("this is a misspelled number")
	Number_Limit      = KindOf("this number exceeds the digit limit")
	Kind_Together     = KindOf("these data types cannot be together")
	Kind_Start        = KindOf("this can't be the beginning")
	Kind_End          = KindOf("this can't be the end")
)

// What error occurred?
var (
	EmptyField      = wrap(Syntax, errors.New("empty field"))
	IncompleteLeft  = wrap(Syntax, errors.New("there are incomplete left parentheses"))
	IncompleteRight = wrap(Syntax, errors.New("there are incomplete right parentheses"))
	AnswerIsNaN     = wrap(Math, errors.New("reports that the value is \"not a number\""))
)

// Interface errors

type Rune struct {
	r rune
	i int
}

type Number struct {
	n string
}

type Kind struct {
	k1, k2 rune
}

// Functions to create an instance with New

func NewRune(r rune, i int) *Rune {
	return &Rune{r: r, i: i}
}

func NewNumber(n string) *Number {
	return &Number{n: n}
}

func NewKind(k1, k2 rune) *Kind {
	return &Kind{k1: k1, k2: k2}
}

// The data error

func (r Rune) Error() string {
	return fmt.Sprintf("'%c' in index: %d", r.r, r.i)
}

func (n Number) Error() string {
	return n.n
}

func (k Kind) Error() string {
	if k.k2 == 0 {
		return fmt.Sprintf("%c", k.k1)
	}
	return fmt.Sprintf("%c:%c", k.k1, k.k2)
}

// Add context to the data error

// Unknown returns an error with the kind of context: Rune_Unknown
func (r Rune) Unknown() error {
	return doubleWrap(Syntax, Rune_Unknown, NewRune(r.r, r.i))
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

// NotTogether returns an error with the kind of context: Kind_Start
func (k Kind) Start() error {
	return doubleWrap(Syntax, Kind_Start, NewKind(k.k1, k.k2))
}

// NotTogether returns an error with the kind of context: Kind_End
func (k Kind) End() error {
	return doubleWrap(Syntax, Kind_End, NewKind(k.k1, k.k2))
}

// !Tool Functions

// wrap adds a wrapper of type error to the already created error
func wrap(kind KindOf, err error) error {
	return fmt.Errorf("%s: %w", kind, err)
}

// doubleWrap adds double wrapper of type error to the already created error
func doubleWrap(k1, k2 KindOf, err error) error {
	return wrap(k1, wrap(k2, err))
}

// As is similar to errors.As func of standard library
func As(err error, target KindOf) bool {
	return strings.Contains(fmt.Sprint(err), string(target))
}
