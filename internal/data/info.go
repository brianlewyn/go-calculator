package data

// Info represents the data to calculate an expression
type Info struct {
	expression *string
	lenght     int
}

// NewInfo abstracts the information of an expression
func NewInfo(expr *string) *Info {
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
