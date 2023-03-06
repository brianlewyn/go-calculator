package analyse

import (
	"github.com/brianlewyn/go-calculator/internal/data"
	d "github.com/brianlewyn/go-linked-list/doubly"
)

// Analyser returns nil if the basic math expression is correct, otherwise returns an error
func Analyser(list *d.Doubly[*data.Token]) error {
	return analyse{list: list}.isCorrect()
}
