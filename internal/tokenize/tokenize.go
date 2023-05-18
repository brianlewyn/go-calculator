package tokenize

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
)

// Tokenizer returns the expression in an Tokenized Linked List and nil,
// otherwise returns nil and an error
func Tokenizer(expression string) (*plugin.TokenList, error) {
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
func toTokenizedLinkedList(expression string) (*plugin.TokenList, error) {
	num, start, lock := new(string), new(int), new(bool)
	kind, ok := data.TokenKind(0), false
	list := plugin.NewTokenList()

	for i, r := range expression {
		if data.IsDecimal(r) {
			if isFullNumber(expression, i, start, lock, num) {
				list.Append(data.NewNumberToken(*num))
			}
			continue
		}

		if kind, ok = data.TokenKindMap[r]; ok {
			list.Append(data.NewSymbolToken(kind))
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
func rebuildTokenizedLinkedList(list *plugin.TokenList) {
	for temp := list.Head(); temp != nil; temp = temp.Next() {

		if areRightAndLeftTokenTogether(temp) {
			list.Connect(temp, data.NewSymbolToken(data.MulToken))
			continue
		}

		if areLeftAndSubTokenTogether(temp) {
			list.Connect(temp, data.NewNumberToken(data.Zero))
			continue
		}

		if canNextAddTokenBeRemoved(temp) {
			list.Disconnect(temp.Next())
			continue
		}

		addParenthesesIfPossible(temp, list)
	}

	if isKind(list.Head(), data.AddToken) {
		list.RemoveHead()
	}

	if isKind(list.Head(), data.SubToken) {
		list.Prepend(data.NewNumberToken(data.Zero))
	}
}

// !Tool Functions

// cutNumberRange cuts a string of numbers from the expression in a specified range
func cutNumberRange(expression string, start, end int) string {
	if end == len(expression) {
		return expression[start:]
	}
	return expression[start:end]
}

// isFullNumber returns true if the string number is complete
func isFullNumber(expression string, index int, start *int, lock *bool, num *string) bool {
	if !*lock {
		*start, *lock = index, true
	}

	if isNextDecimal(index+1, expression) {
		return false
	}

	*num = cutNumberRange(expression, *start, index+1)
	*lock = false
	return true
}

// isNextDecimal returns true if the rune in the given index is a Decimal, otherwise returns false
func isNextDecimal(iNext int, expression string) bool {
	if iNext < len(expression) {
		return data.IsDecimal(rune(expression[iNext]))
	}
	return false
}

// areRightAndLeftTokenTogether returns true there are a RightToken and a LeftToken together
//
//	)( => )*(
func areRightAndLeftTokenTogether(node *plugin.TokenNode) bool {
	if !isKind(node, data.RightToken) {
		return false
	}

	if node.Next() == nil {
		return false
	}

	return isKind(node.Next(), data.LeftToken)
}

// areLeftAndSubTokenTogether returns true there are a LeftToken and a SubToken together
//
//	(- => (0-
func areLeftAndSubTokenTogether(node *plugin.TokenNode) bool {
	if !isKind(node, data.LeftToken) {
		return false
	}

	if node.Next() == nil {
		return false
	}

	return isKind(node.Next(), data.SubToken)
}

// canNextAddTokenBeRemoved returns true if AddToken at the next index
// can be removed according to the following rules:
//
//	# = { %, *, +, -, /, ^, √, ( }
//
//	From: #+n, #+π, #+(, #+√n, #+√π, #+√(...)
//	To #n, #π, #(, #√n, #√π, #√(...)
func canNextAddTokenBeRemoved(node *plugin.TokenNode) bool {
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
func addParenthesesIfPossible(node *plugin.TokenNode, list *plugin.TokenList) {
	if !isKindFn(node, data.IsSpecialToken) {
		return
	}

	temp := node.Next()
	if temp == nil {
		return
	}

	if !isKind(temp, data.SubToken) {
		return
	}

	temp = temp.Next()
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

	temp = temp.Next()
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
func addParentheseInRangeAfterNode(node *plugin.TokenNode, index int, list *plugin.TokenList) {
	list.Connect(node, data.NewSymbolToken(data.LeftToken))
	list.ConnectFrom(node, index, data.NewSymbolToken(data.RightToken))
}

// wrapWithOtherParentheses add add parentheses wrapping a sign and another operation with parentheses
func wrapWithOtherParentheses(node *plugin.TokenNode, list *plugin.TokenList) {
	var nLeft, nRight int

	for temp := node; temp != nil; temp = temp.Next() {
		if isKind(temp, data.LeftToken) {
			nLeft++
			continue
		}

		if isKind(temp, data.RightToken) {
			nRight++

			if nLeft == nRight {
				list.Connect(node, data.NewSymbolToken(data.LeftToken))
				list.Connect(temp, data.NewSymbolToken(data.RightToken))
				break
			}
		}
	}
}

// isKind returns true if the node's token is equal to the given token, otherwise returns false
func isKind(node *plugin.TokenNode, token data.TokenKind) bool {
	return node.Token().Kind() == token
}

// isKindFn returns true if node's token is equal to the given token of a function, otherwise returns false
func isKindFn(node *plugin.TokenNode, tokenFn func(token data.TokenKind) bool) bool {
	return tokenFn(node.Token().Kind())
}
