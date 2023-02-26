package analyse

import "github.com/brianlewyn/go-calculator/internal/data"

// Analyser returns true if the basic math expression is correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error
func Analyser(list *[]*data.Token) bool {
	analyser := &analyse{list: list}
	return analyser.isCorrect()
}
