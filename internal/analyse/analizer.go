package analyse

import "github.com/brianlewyn/go-calculator/internal/plugin"

// Analyser returns nil if the basic math expression is correct, otherwise returns an error
func Analyser(list *plugin.TokenList) error {
	return analyse{list: list}.isCorrect()
}
