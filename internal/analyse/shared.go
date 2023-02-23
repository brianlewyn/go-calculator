package analyse

import "github.com/brianlewyn/go-calculator/internal/data"

var eError error

// analyse represents a parser for expression
type analyse struct {
	expression *string
	lenght     *int
}

// Analyser returns true if the basic math expression is correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error
func Analyser(data *data.Data) bool {
	analyser := &analyse{
		expression: data.Expression(),
		lenght:     data.Lenght(),
	}

	if !analyser.isCorrectSyntax() {
		return false
	}

	return analyser.isCorrectExpression()
}

// Error returns a possible error
func Error() error {
	return eError
}
