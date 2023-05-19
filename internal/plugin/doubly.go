package plugin

import (
	"fmt"
	"strings"

	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-linked-list/v2/doubly"
)

// TokenList represents the Token List
type TokenList struct {
	list *doubly.Doubly[data.Token]
}

// NewTokenList creates a new instance of TokenList
func NewTokenList() *TokenList {
	return &TokenList{list: &doubly.Doubly[data.Token]{}}
}

// Size returns the size list
func (l TokenList) Size() int {
	return l.list.Size()
}

// Head returns the first TokenNode in the list
func (l TokenList) Head() *TokenNode {
	return &TokenNode{node: l.list.NHead()}
}

// Tail returns the last TokenNode in the list
func (l TokenList) Tail() *TokenNode {
	return &TokenNode{node: l.list.NTail()}
}

// String converts a TokenList to string
func (l *TokenList) String() string {
	var b strings.Builder

	for temp := l.list.NHead(); temp != nil; temp = temp.NNext() {
		fmt.Fprintf(&b, "%c", data.RuneMap[temp.Data().Kind()])
	}

	return b.String()
}

// IsEmpty returns true if list is empty, otherwise returns false
func (l TokenList) IsEmpty() bool {
	return l.list.IsEmpty()
}

// Apppend adds a new token to the end of the list
func (l *TokenList) Append(token data.Token) {
	l.list.DAppend(token)
}

// Flush deletes all list items
func (l *TokenList) Flush() {
	l.list.Flush(false)
}

// Connect connets a TokenNode to a new TokenNode with the data.Token
func (l *TokenList) Connect(from *TokenNode, token data.Token) {
	l.list.Connect(from.node, doubly.NewNode(token))
}

// Disconnect disconnets a TokenNode from the list
func (l *TokenList) Disconnect(node *TokenNode) {
	l.list.Disconnect(node.node)
}

// ConnectFrom connects a new TokenNode at 'kIndex' forward from the 'from' TokenNode
func (l *TokenList) ConnectFrom(from *TokenNode, kIndex int, token data.Token) {
	l.list.ConnectForwardFrom(from.node, kIndex, doubly.NewNode(token))
}

// RemoveHead removes a head TokenNode from the list
func (l *TokenList) RemoveHead() {
	l.list.NPullHead()
}

// Prepend adds a new data.Token to the beginning of the list
func (l *TokenList) Prepend(token data.Token) {
	l.list.DPrepend(token)
}
