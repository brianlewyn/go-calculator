package plugin

import (
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-linked-list/v2/doubly"
)

// TokenNode represents the Token Node
type TokenNode struct {
	node *doubly.Node[data.Token]
}

// Token returns the data of the node
func (n TokenNode) Token() data.Token {
	return n.node.Data()
}

// Update updates a data.Token
func (n TokenNode) Update(data *data.Token) {
	n.node.Update(*data)
}

// Prev return the previous node
func (n TokenNode) Prev() *TokenNode {
	if prev := n.node.NPrev(); prev != nil {
		return &TokenNode{node: prev}
	}
	return nil
}

// Next return the next node
func (n TokenNode) Next() *TokenNode {
	if next := n.node.NNext(); next != nil {
		return &TokenNode{node: next}
	}
	return nil
}
