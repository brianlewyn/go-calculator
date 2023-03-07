package tokenize

import (
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
)

// Tokenizer tokenizes the expression in a linked list,
//
// but while creating the list, the expression is removed
func Tokenizer(data *data.Data) (*plugin.TokenList, error) {
	tokenizer := tokenize{
		expression: data.Expression(),
		lenght:     data.Lenght(),
	}

	list, err := tokenizer.linkedList()
	if err != nil {
		return nil, err
	}

	list, err = tokenizer.rebuild(list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
