package tokenize

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-linked-list/v2/doubly"
)

var (
	left  = data.NewSymbolToken(data.LeftToken)
	right = data.NewSymbolToken(data.RightToken)
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
	k, list := 0, doubly.New[data.Token]()

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
			symbol := data.NewSymbolToken(data.MulToken)
			list.Connect(temp, doubly.NewNode(symbol))
			continue
		}

		if areLeftAndSubTokenTogether(temp) {
			zero := data.NewNumberToken("0")
			list.Connect(temp, doubly.NewNode(zero))
			continue
		}

		if canRemoveNextAddToken(temp) {
			list.Disconnect(temp.NNext())
			continue
		}

		if !canWrapNextSubToken(temp, list) {
			canWrapNextRootToken(temp, list)
		}
	}

	if isKind(list.NHead(), data.AddToken) {
		if list.Size() != 1 {
			list.RemoveHead()
		}
	}

	if isKind(list.NHead(), data.SubToken) {
		list.DPrepend(data.NewNumberToken("0"))
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
	if !isKind(node, data.RightToken) || node.NNext() == nil {
		return false
	}
	return isKind(node.NNext(), data.LeftToken)
}

// areLeftAndSubTokenTogether returns true there are a LeftToken and a SubToken together
//
//	(- => (0-
func areLeftAndSubTokenTogether(node *doubly.Node[data.Token]) bool {
	if !isKind(node, data.LeftToken) || node.NNext() == nil {
		return false
	}
	return isKind(node.NNext(), data.SubToken)
}

// canRemoveNextAddToken returns true if AddToken at the next index
// can be removed according to the following rules:
//
//	# = { %, *, +, -, /, ^, √, ( }
//
//	From: #+n, #+π, #+(, #+√n, #+√π, #+√(...)
//	To: #n, #π, #(, #√n, #√π, #√(...)
func canRemoveNextAddToken(node *doubly.Node[data.Token]) bool {
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

	// #+n, #+π
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

// canWrapNextSubToken adds parentheses according to the following rules:
//
//	# = {%, *, +, -, /, ^, √}
//
//	From: #-n, #-π, #-(, #-√n, #-√π, #-√(...)
//	To: #(-n), #(-π), #(-(...)), #(-√n) #(-√π) #(-√(...))
func canWrapNextSubToken(node *doubly.Node[data.Token], list *doubly.Doubly[data.Token]) bool {
	if !isKindFn(node, data.IsSpecialToken) {
		return false
	}

	temp := node.NNext()
	if temp == nil {
		return false
	}

	if !isKind(temp, data.SubToken) {
		return false
	}

	temp = temp.NNext()
	if temp == nil {
		return false
	}

	// #-n, #-π
	if isKindFn(temp, data.IsNumPiToken) {
		addParentheseInRangeAfterNode(node, 4, list)
		return true
	}

	// #-(...)
	if isKind(temp, data.LeftToken) {
		wrapWithOtherParentheses(node, list)
		return true
	}

	// #-√
	if !isKind(temp, data.RootToken) {
		return true
	}

	temp = temp.NNext()
	if temp == nil {
		return false
	}

	// #-√n, #-√π
	if isKindFn(temp, data.IsNumPiToken) {
		addParentheseInRangeAfterNode(node, 5, list)
		return true
	}

	// #-√(...)
	if isKind(temp, data.LeftToken) {
		wrapWithOtherParentheses(node, list)
		return true
	}

	return false
}

// canWrapNextRootToken adds parentheses according to the following rules:
//
//	Can: √√, √√√, √√√√, ...
//
//	From: √√n, √√π, √√(...)
//	To: √(√n), √(√π), √(√(...))
func canWrapNextRootToken(node *doubly.Node[data.Token], list *doubly.Doubly[data.Token]) {
	if !isKind(node, data.RootToken) {
		return
	}

	temp := node.NNext()
	if temp == nil {
		return
	}

	for space := 4; true; space++ {
		if !isKind(temp, data.RootToken) {
			return
		}

		temp = temp.NNext()
		if temp == nil {
			return
		}

		// √√n, √√√n, ...;  √√π, √√√π, ...
		if isKindFn(temp, data.IsNumPiToken) {
			addParentheseInRangeAfterNode(node, space, list)
			return
		}

		// √√(...), √√√(...), ...
		if isKind(temp, data.LeftToken) {
			wrapWithOtherParentheses(node, list)
			return
		}
	}
}

// addParentheseInRangeAfterNode adds a LeftToken at the next index of the node and
// a RightToken at given index after the node
func addParentheseInRangeAfterNode(node *doubly.Node[data.Token], index int, list *doubly.Doubly[data.Token]) {
	list.Connect(node, doubly.NewNode(left))
	list.ConnectFrom(node, index, doubly.NewNode(right))
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
				list.Connect(node, doubly.NewNode(left))
				list.Connect(temp, doubly.NewNode(right))
				break
			}
		}
	}
}

// isKind returns true if the node's kind is equal to the given kind, otherwise returns false
func isKind(node *doubly.Node[data.Token], token data.TokenKind) bool {
	return node.Data().Kind() == token
}

// isKindFn returns true if node's kind is equal to the given kind of a function, otherwise returns false
func isKindFn(node *doubly.Node[data.Token], tokenFn func(token data.TokenKind) bool) bool {
	return tokenFn(node.Data().Kind())
}
