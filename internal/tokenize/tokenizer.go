package tokenize

import "github.com/brianlewyn/go-calculator/internal/data"

// Tokenizer tokenizes the expression in a linked list,
//
// but while creating the list, the expression is removed
func Tokenizer(data *data.Data) *[]*data.Token {
	tokenizer := tokenize{
		expression: data.Expression(),
		lenght:     data.Lenght(),
	}
	return tokenizer.rebuild(tokenizer.linkedList())
}
