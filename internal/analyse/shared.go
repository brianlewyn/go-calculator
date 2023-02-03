package analyse

type analyse struct {
	expr *string
	err  *error
}

func New(expr *string) *analyse {
	return &analyse{expr: expr, err: nil}
}

func (a analyse) Error() error {
	return *a.err
}
