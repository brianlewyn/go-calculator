package plugin

import (
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-linked-list/v2/doubly"
)

// TokenNode represents the Token Node
type TokenNode struct {
	node *doubly.Node[data.Token]
}

// Token returns a data.Token
func (n TokenNode) Token() data.Token {
	return n.node.Data()
}

// Update updates a data.Token
func (n TokenNode) Update(data data.Token) {
	n.node.Update(data)
}

// Prev returns the previous TokenNode
func (n TokenNode) Prev() *TokenNode {
	if n.node.NPrev() == nil {
		return nil
	}
	return &TokenNode{node: n.node.NPrev()}
}

// Next returns the next TokenNode
func (n TokenNode) Next() *TokenNode {
	if n.node.NNext() == nil {
		return nil
	}
	return &TokenNode{node: n.node.NNext()}
}

// TPrev returns the previous data.Token
func (n TokenNode) TPrev() data.Token {
	if n.node.NPrev() == nil {
		return nil
	}
	return n.node.DPrev()
}

// TNext returns the next data.Token
func (n TokenNode) TNext() data.Token {
	if n.node.NNext() == nil {
		return nil
	}
	return n.node.DNext()
}
