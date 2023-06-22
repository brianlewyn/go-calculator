package ierr

import (
	"errors"
	"fmt"
	"strings"
)

// !KindOf represents the type of context error
type KindOf string

// !What kind of main error occurred?
const (
	Syntax = KindOf("syntax error")
	Math   = KindOf("math error")
)

// !What kind of context error occurred?
const (
	CtxRuneUnknown      = KindOf("this is an unknown rune")
	CtxNumberMisspelled = KindOf("this is a misspelled number")
	CtxNumberLimit      = KindOf("this number exceeds the digit limit")
	CtxKindNotTogether  = KindOf("these data types cannot be together")
	CtxKindStart        = KindOf("this can't be the beginning")
	CtxKindEnd          = KindOf("this can't be the end")
)

// !What error occurred?
var (
	EmptyField      = wrap(Syntax, errors.New("empty field"))
	IncompleteLeft  = wrap(Syntax, errors.New("there are incomplete left parentheses"))
	IncompleteRight = wrap(Syntax, errors.New("there are incomplete right parentheses"))
	IsNaN           = wrap(Math, errors.New("reports that the value is \"not a number\""))
	IsInf           = wrap(Math, errors.New("reports that the value is any type of infinity"))
)

// !Interface errors

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

// !Functions to create an instance with New

func NewRune(r rune, i int) *Rune {
	return &Rune{r: r, i: i}
}

func NewNumber(n string) *Number {
	return &Number{n: n}
}

func NewKind(k1, k2 rune) *Kind {
	return &Kind{k1: k1, k2: k2}
}

// !The data error

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

// !Add context to the data error

// RuneUnknown returns an error with the kind of context: CtxRuneUnknown
func RuneUnknown(r rune, i int) error {
	return doubleWrap(Syntax, CtxRuneUnknown, NewRune(r, i))
}

// NumberMisspelled returns an error with the kind of context: CtxNumberMisspelled
func NumberMisspelled(n string) error {
	return doubleWrap(Syntax, CtxNumberMisspelled, NewNumber(n))
}

// NumberLimit returns an error with the kind of context: CtxNumberLimit
func NumberLimit(n string) error {
	return doubleWrap(Syntax, CtxNumberLimit, NewNumber(n))
}

// KindNotTogether returns an error with the kind of context: CtxKindTogether
func KindNotTogether(k1, k2 rune) error {
	return doubleWrap(Syntax, CtxKindNotTogether, NewKind(k1, k2))
}

// KindStart returns an error with the kind of context: CtxKindStart
func KindStart(k rune) error {
	return doubleWrap(Syntax, CtxKindStart, NewKind(k, 0))
}

// KindEnd returns an error with the kind of context: CtxKindEnd
func KindEnd(k rune) error {
	return doubleWrap(Syntax, CtxKindEnd, NewKind(k, 0))
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
