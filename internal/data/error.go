package data

// Error represents a complex error with methods
type Error interface {
	Bug() error
	Exp() string
	String() string
}

// data represents the error data to calculate an expression
type data struct {
	bug  error
	expr string
}

// NewError creates a data instance
func NewError(expr string, bug error) Error {
	return data{expr: expr, bug: bug}
}

// Exp returns the state of the basic math expression
func (d data) Exp() string {
	return d.expr
}

// Bug return the error of the basic string expression
func (d data) Bug() error {
	return d.bug
}

// String converts the error data to a string
func (d data) String() string {
	return d.bug.Error() + "\n" + d.expr
}
