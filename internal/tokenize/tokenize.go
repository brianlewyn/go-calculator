package tokenize

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-linked-list/v2/doubly"
)

// Tokenizer returns the expression in an Tokenized Linked List and nil,
// otherwise returns nil and an error
func Tokenizer(expression string) (*doubly.Doubly[data.Token], error) {
	if expression == "" {
		return nil, ierr.EmptyField
	}

	list, err := toTokenizedLinkedList(expression)
	if err != nil {
		return nil, err
	}

	rebuildTokenizedLinkedList(list)
	return list, nil
}

// toTokenizedLinkedList returns the expression in a raw Tokenized Linked List and nil,
// otherwise returns nil and an error
func toTokenizedLinkedList(expression string) (*doubly.Doubly[data.Token], error) {
	k, list := 0, doubly.NewDoubly[data.Token]()

	for i, r := range expression {
		if data.IsDecimal(r) {
			if i >= k {
				num := getFullNumber(expression[i:])
				list.DAppend(data.NewNumberToken(num))
				k = i + len(num)
			}
			continue
		}

		if kind, ok := data.TokenKindMap[r]; ok {
			list.DAppend(data.NewSymbolToken(kind))
			continue
		}

		if r != data.Gap {
			return nil, ierr.RuneUnknown(r, i)
		}
	}

	if list.IsEmpty() {
		return nil, ierr.EmptyField
	}

	return list, nil
}

// rebuildTokenizedLinkedList returns a rebuilt Tokenized Linked List
func rebuildTokenizedLinkedList(list *doubly.Doubly[data.Token]) {
	for temp := list.NHead(); temp != nil; temp = temp.NNext() {

		if areRightAndLeftTokenTogether(temp) {
			list.Connect(temp, doubly.NewNode(data.NewSymbolToken(data.MulToken)))
			continue
		}

		if areLeftAndSubTokenTogether(temp) {
			list.Connect(temp, doubly.NewNode(data.NewNumberToken(data.Zero)))
			continue
		}

		if canNextAddTokenBeRemoved(temp) {
			list.Disconnect(temp.NNext())
			continue
		}

		addParenthesesIfPossible(temp, list)
	}

	if isKind(list.NHead(), data.AddToken) {
		list.NPullHead()
	}

	if isKind(list.NHead(), data.SubToken) {
		list.DPrepend(data.NewNumberToken(data.Zero))
	}
}

// !Tool Functions

// getFullNumber returns a full number
func getFullNumber(expression string) string {
	for i, r := range expression {
		if data.IsDecimal(r) {
			continue
		}
		return expression[0:i]
	}
	return expression[0:]
}

// areRightAndLeftTokenTogether returns true there are a RightToken and a LeftToken together
//
//	)( => )*(
func areRightAndLeftTokenTogether(node *doubly.Node[data.Token]) bool {
	if !isKind(node, data.RightToken) {
		return false
	}

	if node.NNext() == nil {
		return false
	}

	return isKind(node.NNext(), data.LeftToken)
}

// areLeftAndSubTokenTogether returns true there are a LeftToken and a SubToken together
//
//	(- => (0-
func areLeftAndSubTokenTogether(node *doubly.Node[data.Token]) bool {
	if !isKind(node, data.LeftToken) {
		return false
	}

	if node.NNext() == nil {
		return false
	}

	return isKind(node.NNext(), data.SubToken)
}

// canNextAddTokenBeRemoved returns true if AddToken at the next index
// can be removed according to the following rules:
//
//	# = { %, *, +, -, /, ^, √, ( }
//
//	From: #+n, #+π, #+(, #+√n, #+√π, #+√(...)
//	To #n, #π, #(, #√n, #√π, #√(...)
func canNextAddTokenBeRemoved(node *doubly.Node[data.Token]) bool {
	if !isKindFn(node, data.IsSpecialToken) {
		if !isKind(node, data.LeftToken) {
			return false
		}
	}

	temp := node.NNext()
	if temp == nil {
		return false
	}

	if !isKind(temp, data.AddToken) {
		return false
	}

	temp = temp.NNext()
	if temp == nil {
		return false
	}

	// #+n of a function, #+π
	if isKindFn(temp, data.IsNumPiToken) {
		return true
	}

	// #+(...)
	if isKind(temp, data.LeftToken) {
		return true
	}

	// #+√
	if isKind(temp, data.RootToken) {
		return true
	}

	return false
}

// addParenthesesIfPossible adds parentheses according to the following rules:
//
//	# = {%, *, +, -, /, ^, √}
//
//	From: #-n, #-π, #-(, #-√n, #-√π, #-√(...)
//	To #(-n), #(-π), #(-(...)), #(-√n) #(-√π) #(-√(...))
func addParenthesesIfPossible(node *doubly.Node[data.Token], list *doubly.Doubly[data.Token]) {
	if !isKindFn(node, data.IsSpecialToken) {
		return
	}

	temp := node.NNext()
	if temp == nil {
		return
	}

	if !isKind(temp, data.SubToken) {
		return
	}

	temp = temp.NNext()
	if temp == nil {
		return
	}

	// #-n of a function, #-π
	if isKindFn(temp, data.IsNumPiToken) {
		addParentheseInRangeAfterNode(node, 4, list)
		return
	}

	// #-(...)
	if isKind(temp, data.LeftToken) {
		wrapWithOtherParentheses(node, list)
		return
	}

	// #-√
	if !isKind(temp, data.RootToken) {
		return
	}

	temp = temp.NNext()
	if temp == nil {
		return
	}

	// #-√n, #-√π
	if isKindFn(temp, data.IsNumPiToken) {
		addParentheseInRangeAfterNode(node, 5, list)
		return
	}

	// #-√(...)
	if isKind(temp, data.LeftToken) {
		wrapWithOtherParentheses(node, list)
		return
	}
}

// addParentheseInRangeAfterNode adds a LeftToken at the next index of the node and
// a RightToken at given index after the node
func addParentheseInRangeAfterNode(node *doubly.Node[data.Token], index int, list *doubly.Doubly[data.Token]) {
	list.Connect(node, doubly.NewNode(data.NewSymbolToken(data.LeftToken)))
	list.ConnectFrom(node, index, doubly.NewNode(data.NewSymbolToken(data.RightToken)))
}

// wrapWithOtherParentheses add add parentheses wrapping a sign and another operation with parentheses
func wrapWithOtherParentheses(node *doubly.Node[data.Token], list *doubly.Doubly[data.Token]) {
	var nLeft, nRight int

	for temp := node; temp != nil; temp = temp.NNext() {
		if isKind(temp, data.LeftToken) {
			nLeft++
			continue
		}

		if isKind(temp, data.RightToken) {
			nRight++

			if nLeft == nRight {
				list.Connect(node, doubly.NewNode(data.NewSymbolToken(data.LeftToken)))
				list.Connect(temp, doubly.NewNode(data.NewSymbolToken(data.RightToken)))
				break
			}
		}
	}
}

// isKind returns true if the node's token is equal to the given token, otherwise returns false
func isKind(node *doubly.Node[data.Token], token data.TokenKind) bool {
	return node.Data().Kind() == token
}

// isKindFn returns true if node's token is equal to the given token of a function, otherwise returns false
func isKindFn(node *doubly.Node[data.Token], tokenFn func(token data.TokenKind) bool) bool {
	return tokenFn(node.Data().Kind())
}
