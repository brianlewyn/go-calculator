package tokenize

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/doubly"
)

var (
	left  = data.NewSymbolToken(data.LeftToken)
	right = data.NewSymbolToken(data.RightToken)
)

// Tokenizer returns the expression in an Tokenized Linked List and nil,
// otherwise returns nil and an error
func Tokenizer(expression string) (*doubly.Doubly, error) {
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
func toTokenizedLinkedList(expression string) (*doubly.Doubly, error) {
	k, list := 0, doubly.New()

	for i, r := range expression {
		if data.IsDecimal(r) {
			if i >= k {
				num := getFullNumber(expression[i:])
				list.PushBack(data.NewNumberToken(num))
				k = i + len(num)
			}
			continue
		}

		if kind, ok := data.TokenKindMap[r]; ok {
			list.PushBack(data.NewSymbolToken(kind))
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
func rebuildTokenizedLinkedList(list *doubly.Doubly) {
	for temp := list.Head(); temp != nil; temp = temp.Next() {

		if areRightAndLeftTokenTogether(temp) {
			symbol := data.NewSymbolToken(data.MulToken)
			list.ConnectAfterNode(temp, doubly.NewNode(symbol))
			continue
		}

		if areLeftAndSubTokenTogether(temp) {
			zero := data.NewNumberToken("0")
			list.ConnectAfterNode(temp, doubly.NewNode(zero))
			continue
		}

		if canRemoveNextAddToken(temp) {
			list.RemoveNode(temp.Next())
			continue
		}

		if !canWrapNextSubToken(temp, list) {
			canWrapNextRootToken(temp, list)
		}
	}

	if isKind(list.Head(), data.AddToken) {
		if list.Size() != 1 {
			list.RemoveHead()
		}
	}

	if isKind(list.Head(), data.SubToken) {
		list.PushFront(data.NewNumberToken("0"))
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
func areRightAndLeftTokenTogether(node *doubly.Node) bool {
	if !isKind(node, data.RightToken) || node.Next() == nil {
		return false
	}
	return isKind(node.Next(), data.LeftToken)
}

// areLeftAndSubTokenTogether returns true there are a LeftToken and a SubToken together
//
//	(- => (0-
func areLeftAndSubTokenTogether(node *doubly.Node) bool {
	if !isKind(node, data.LeftToken) || node.Next() == nil {
		return false
	}
	return isKind(node.Next(), data.SubToken)
}

// canRemoveNextAddToken returns true if AddToken at the next index
// can be removed according to the following rules:
//
//	# = { %, *, +, -, /, ^, √, ( }
//
//	From: #+n, #+π, #+(, #+√n, #+√π, #+√(...)
//	To: #n, #π, #(, #√n, #√π, #√(...)
func canRemoveNextAddToken(node *doubly.Node) bool {
	if !isKindFn(node, data.IsSpecialToken) {
		if !isKind(node, data.LeftToken) {
			return false
		}
	}

	temp := node.Next()
	if temp == nil {
		return false
	}

	if !isKind(temp, data.AddToken) {
		return false
	}

	temp = temp.Next()
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
func canWrapNextSubToken(node *doubly.Node, list *doubly.Doubly) bool {
	if !isKindFn(node, data.IsSpecialToken) {
		return false
	}

	temp := node.Next()
	if temp == nil {
		return false
	}

	if !isKind(temp, data.SubToken) {
		return false
	}

	temp = temp.Next()
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

	temp = temp.Next()
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
func canWrapNextRootToken(node *doubly.Node, list *doubly.Doubly) {
	if !isKind(node, data.RootToken) {
		return
	}

	temp := node.Next()
	if temp == nil {
		return
	}

	for space := 4; true; space++ {
		if !isKind(temp, data.RootToken) {
			return
		}

		temp = temp.Next()
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
func addParentheseInRangeAfterNode(node *doubly.Node, index int, list *doubly.Doubly) {
	list.ConnectAfterNode(node, doubly.NewNode(left))
	list.PushNodeForwardFrom(node, index, doubly.NewNode(right))
}

// wrapWithOtherParentheses add add parentheses wrapping a sign and another operation with parentheses
func wrapWithOtherParentheses(node *doubly.Node, list *doubly.Doubly) {
	var nLeft, nRight int

	for temp := node; temp != nil; temp = temp.Next() {
		if isKind(temp, data.LeftToken) {
			nLeft++
			continue
		}

		if isKind(temp, data.RightToken) {
			nRight++

			if nLeft == nRight {
				list.ConnectAfterNode(node, doubly.NewNode(left))
				list.ConnectAfterNode(temp, doubly.NewNode(right))
				break
			}
		}
	}
}

// isKind returns true if the node's kind is equal to the given kind, otherwise returns false
func isKind(node *doubly.Node, token data.TokenKind) bool {
	return node.Token().Kind() == token
}

// isKindFn returns true if node's kind is equal to the given kind of a function, otherwise returns false
func isKindFn(node *doubly.Node, tokenFn func(token data.TokenKind) bool) bool {
	return tokenFn(node.Token().Kind())
}
