package data

import "fmt"

// ExpData represents the data to calculate an expression
type ExpData struct {
	expression *string
	lenght     int
}

// New creates a ExpData instance
func NewExpData(exp *string) *ExpData {
	return &ExpData{
		expression: exp,
		lenght:     len(*exp),
	}
}

// Expression returns the basic math expression
func (e ExpData) Expression() *string {
	return e.expression
}

// Lenght return the lenght of the basic string expression
func (e ExpData) Lenght() *int {
	return &e.lenght
}

// IErrData represents the error data as a complex error
type IErrData interface {
	Bug() error
	Exp() string
	String() string
}

// ErrData represents the error data to calculate an expression
type ErrData struct {
	bug error
	exp string
}

// New creates a ExpData instance
func NewErrData(exp string, bug error) *ErrData {
	return &ErrData{exp: exp, bug: bug}
}

// Exp returns the state of the basic math expression
func (e ErrData) Exp() string {
	return e.exp
}

// Bug return the error of the basic string expression
func (e ErrData) Bug() error {
	return e.bug
}

// String converts the error data to a string
func (e ErrData) String() string {
	return fmt.Sprintf("%s\n%s", e.bug, e.exp)
}
