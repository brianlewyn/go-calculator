package plugin

import (
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-linked-list/doubly"
)

// TokenNode represents the Token Node
type TokenNode struct {
	*doubly.Node[*data.Token]
}

// TokenList represents the Token List
type TokenList struct {
	*doubly.Doubly[*data.Token]
}

func NewTokenList() *TokenList {
	return &TokenList{&doubly.Doubly[*data.Token]{}}
}

// Head returns the first node in the list, but if the list is empty returns nil
func (l TokenList) Head() *TokenNode {
	return &TokenNode{l.Doubly.Head()}
}

// Tail returns the last node in the list, but if the list is empty returns nil
func (l TokenList) Tail() *TokenNode {
	return &TokenNode{l.Doubly.Tail()}
}

// Prepend adds a new token to the end of the list and returns nil,
// and otherwise returns an error
func (l *TokenList) Append(token *data.Token) error {
	return l.Doubly.Append(doubly.NewNode(token))
}

// Prepend adds a new token to the start of the list and returns nil,
// and otherwise returns an error
func (l *TokenList) Prepend(token *data.Token) error {
	return l.Doubly.Prepend(doubly.NewNode(token))
}

// Insert inserts a new token at the given index and returns nil,
// and otherwise returns an error
func (l *TokenList) Insert(i int, token *data.Token) error {
	return l.Doubly.Insert(i, doubly.NewNode(token))
}
