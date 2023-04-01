package data

import "fmt"

// Info represents the data to calculate an expression
type Info struct {
	expression *string
	lenght     int
}

// Abstraction gets the information of an expression
func Abstraction(expr *string) *Info {
	return &Info{
		expression: expr,
		lenght:     len(*expr),
	}
}

// Expression returns the basic math expression
func (e Info) Expression() *string {
	return e.expression
}

// Lenght return the lenght of the basic string expression
func (e Info) Lenght() *int {
	return &e.lenght
}

// Error represents a complex error with methods
type Error interface {
	Bug() error
	Exp() string
	String() string
}

// ErrData represents the error data to calculate an expression
type ErrData struct {
	bug  error
	expr string
}

// New creates a Info instance
func NewErrData(expr string, bug error) *ErrData {
	return &ErrData{expr: expr, bug: bug}
}

// Exp returns the state of the basic math expression
func (e ErrData) Exp() string {
	return e.expr
}

// Bug return the error of the basic string expression
func (e ErrData) Bug() error {
	return e.bug
}

// String converts the error data to a string
func (e ErrData) String() string {
	return fmt.Sprintf("%s\n%s", e.bug, e.expr)
}
