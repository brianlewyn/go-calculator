package plugin

import (
	"fmt"
	"strings"

	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-linked-list/doubly"
)

// TokenNode represents the Token Node
type TokenNode struct {
	node *doubly.Node[*data.Token]
}

// Token returns the data of the node
func (n TokenNode) Token() *data.Token {
	return n.node.Data()
}

// Prev return the previous node
func (n TokenNode) Prev() *TokenNode {
	if n.node.Prev() != nil {
		return &TokenNode{node: n.node.Prev()}
	}
	return nil
}

// Next return the next node
func (n TokenNode) Next() *TokenNode {
	if n.node.Next() != nil {
		return &TokenNode{node: n.node.Next()}
	}
	return nil
}

// TokenList represents the Token List
type TokenList struct {
	list *doubly.Doubly[*data.Token]
}

// NewTokenList creates a new instance of TokenList
func NewTokenList() *TokenList {
	return &TokenList{list: &doubly.Doubly[*data.Token]{}}
}

// Size returns the total number of nodes in the list
func (l TokenList) Size() int {
	return l.list.Size()
}

// Head returns the first node in the list, but if the list is empty returns nil
func (l TokenList) Head() *TokenNode {
	return &TokenNode{node: l.list.Head()}
}

// Tail returns the last node in the list, but if the list is empty returns nil
func (l TokenList) Tail() *TokenNode {
	return &TokenNode{node: l.list.Tail()}
}

func (l *TokenList) String() string {
	if l.list.IsEmpty() {
		return "list <nil>"
	}

	var b strings.Builder
	for temp := l.list.Head(); temp != nil; temp = temp.Next() {
		fmt.Fprintf(&b, "%c", data.FromTokenKindToRune(temp.Data().Kind()))
	}

	return b.String()
}

// IsEmpty returns true if list is empty and otherwise returns false
func (l TokenList) IsEmpty() bool {
	return l.list.IsEmpty()
}

// Prepend adds a new token to the start of the list and returns nil,
// and otherwise returns an error
func (l *TokenList) Prepend(token *data.Token) error {
	return l.list.Prepend(doubly.NewNode(token))
}

// Insert inserts a new token at the given index and returns nil,
// and otherwise returns an error
func (l *TokenList) Insert(i int, token *data.Token) error {
	return l.list.Insert(i, doubly.NewNode(token))
}

// Prepend adds a new token to the end of the list and returns nil,
// and otherwise returns an error
func (l *TokenList) Append(token *data.Token) error {
	return l.list.Append(doubly.NewNode(token))
}
