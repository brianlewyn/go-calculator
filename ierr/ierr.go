package ierr

import (
	"fmt"
)

type kind string

// ! What kind of bug?
const (
	syntax = kind("syntax error")
	// math   = kind("math error")
)

// ! What error occurred?

// var IncorrectSyntax = wrap(syntax, errors.New("empty list"))

type Syntax struct {
	S *rune // Symbol or character
}

// ! The data error

func (e Syntax) Error() string {
	return string(*e.S)
}

// ! Add context to the data error

func (e Syntax) Wrap() error {
	return wrap(syntax, wrap(
		kind("incorrect syntax"),
		&Syntax{S: e.S}),
	)
}

// Add a wrapper of type error to the already created error
func wrap(kind kind, err error) error {
	return fmt.Errorf("%s: %w", kind, err)
}
