package analyse

// analyse represents a parser for expression
type analyse struct {
	expr *string
}

// Analyser returns true if the basic math expression is correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error
func Analyser(expr *string) bool {
	analyser := &analyse{expr: expr}
	if !analyser.isCorrectSyntax() {
		return false
	}
	return analyser.isCorrectExpression()
}
