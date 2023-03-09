package plugin

import (
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-linked-list/doubly"
)

// TokenNode represents the Token Node
type TokenNode struct {
	*doubly.Node[*data.Token]
}

// NewTokenNode creates a new instance of TokenNode
func NewTokenNode(node *doubly.Node[*data.Token]) *TokenNode {
	return &TokenNode{Node: node}
}

// TokenList represents the Token List
type TokenList struct {
	*doubly.Doubly[*data.Token]
}

// NewTokenList creates a new instance of TokenList
func NewTokenList() *TokenList {
	return &TokenList{&doubly.Doubly[*data.Token]{}}
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
