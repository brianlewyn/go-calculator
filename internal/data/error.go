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
func (e data) Exp() string {
	return e.expr
}

// Bug return the error of the basic string expression
func (e data) Bug() error {
	return e.bug
}

// String converts the error data to a string
func (e data) String() string {
	return e.bug.Error() + "\n" + e.expr
}
