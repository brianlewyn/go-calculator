package analyse

import "github.com/brianlewyn/go-calculator/internal/data"

// Analyser returns nil if the basic math expression is correct, otherwise returns an error
func Analyser(list *[]*data.Token) error {
	return analyse{list: list}.isCorrect()
}
