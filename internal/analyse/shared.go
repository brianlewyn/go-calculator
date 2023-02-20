package analyse

// analyse represents the expression parser
type analyse struct {
	expr *string
}

// New creates an analyse instance
func New(expr *string) *analyse {
	return &analyse{expr: expr}
}
