package data

// Data represents the data to calculate an expression
type Data struct {
	expression *string
	lenght     int
}

// New creates an Data instance
func New(expr *string) *Data {
	return &Data{
		expression: expr,
		lenght:     len(*expr),
	}
}

// Expression returns the basic math expression
func (d Data) Expression() *string {
	return d.expression
}

// Lenght return the lenght of the basic string expression
func (d Data) Lenght() *int {
	return &d.lenght
}
