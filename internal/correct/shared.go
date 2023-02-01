package correct

type analyse struct {
	expr *string
	err  error
}

func NewAnalyser(expr *string) *analyse {
	return &analyse{expr: expr, err: nil}
}

func (a analyse) Error() error {
	return a.err
}
