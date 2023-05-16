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

// Size returns the total number of nodes in the list
func (l TokenList) Size() int {
	return l.list.Size()
}

// Head returns the first node in the list, but if the list is empty returns nil
func (l TokenList) Head() *TokenNode {
	return &TokenNode{node: l.list.NHead()}
}

// Tail returns the last node in the list, but if the list is empty returns nil
func (l TokenList) Tail() *TokenNode {
	return &TokenNode{node: l.list.NTail()}
}

func (l *TokenList) String() string {
	if l.list.IsEmpty() {
		return "list <nil>"
	}

	var b strings.Builder
	for temp := l.list.NHead(); temp != nil; temp = temp.NNext() {
		fmt.Fprintf(&b, "%c", data.ChangeToRune(temp.Data().Kind()))
	}

	return b.String()
}

// IsEmpty returns true if list is empty and otherwise returns false
func (l TokenList) IsEmpty() bool {
	return l.list.IsEmpty()
}

// Apppend adds a new token to the end of the list
func (l *TokenList) Append(token *data.Token) {
	l.list.DAppend(*token)
}

// Flush delete list
func (l *TokenList) Flush() {
	l.list.Flush(false)
}

// Connect connets one node to another
func (l *TokenList) Connect(from *TokenNode, token *data.Token) {
	if from == nil {
		l.list.Connect(nil, doubly.NewNode(*token))
	} else {
		l.list.Connect(from.node, doubly.NewNode(*token))
	}
}

// Disconnect disconnets one node of list
func (l *TokenList) Disconnect(node *TokenNode) {
	if node != nil {
		l.list.Disconnect(node.node)
	}
}

// ConnectForwardFrom connects a new node at the given position 'kIndex' forward from the 'from' node
func (l *TokenList) ConnectFrom(from *TokenNode, kIndex int, token *data.Token) {
	if from != nil {
		l.list.ConnectForwardFrom(from.node, kIndex, doubly.NewNode(*token))
	}
}
