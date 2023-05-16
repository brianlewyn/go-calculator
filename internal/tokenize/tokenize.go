package tokenize

import (
	"strings"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
)

// Tokenizer returns the expression in a Linked List and nil,
// otherwise returns nil and an error
func Tokenizer(expression string) (*plugin.TokenList, data.Error) {
	if len(expression) == 0 {
		return nil, data.NewError(expression, ierr.EmptyField)
	}

	list, err := initLinkedList(expression)
	if err != nil {
		return nil, data.NewError(expression, err)
	}

	err = initRebuild(list)
	if err != nil {
		return nil, data.NewError(expression, err)
	}

	return list, nil
}

// initLinkedList returns an tokenized Linked List and a possible error
func initLinkedList(expression string) (*plugin.TokenList, error) {
	list := plugin.NewTokenList()
	builder := strings.Builder{}

	for i, r := range expression {
		switch {

		// opeartors
		case data.IsMod(r):
			list.Append(data.NewModToken())
		case data.IsMul(r):
			list.Append(data.NewMulToken())
		case data.IsAdd(r):
			list.Append(data.NewAddToken())
		case data.IsSub(r):
			list.Append(data.NewSubToken())
		case data.IsDiv(r):
			list.Append(data.NewDivToken())

		// parentheses
		case data.IsLeft(r):
			list.Append(data.NewLeftToken())
		case data.IsRight(r):
			list.Append(data.NewRightToken())

		// power
		case data.IsPow(r):
			list.Append(data.NewPowToken())

		// pi or root
		case data.IsPi(r):
			list.Append(data.NewPiToken())
		case data.IsRoot(r):
			list.Append(data.NewRootToken())

		// numbers
		case data.IsFloat(r):
			builder.WriteRune(r)

			if !isNextToo(i, expression) {
				list.Append(data.NewNumToken(builder.String()))
				builder.Reset()
			}

		default:
			if !data.IsGap(r) {
				return nil, ierr.NewRune(r, i).Unknown()
			}
		}
	}

	return list, nil
}

// initRebuild returns a rebuilt tokenized Linked List and a possible error
func initRebuild(list *plugin.TokenList) error {
	if list.IsEmpty() {
		return ierr.EmptyField
	}

	for i, temp := 0, list.Head(); temp != nil; i, temp = i+1, temp.Next() {
		if canBeAddedAsterisk(temp) {
			list.Connect(temp, data.NewMulToken())

		} else if canBeAddedZero(temp) {
			list.Connect(temp, data.NewNumToken("0"))

		} else if canBeRemovedPlus(temp) {
			list.Disconnect(temp.Next())

		} else {
			addParenthesesIfPossible(i, temp, list)
		}
	}

	modifyFirstToken(list)
	return nil
}

// !Tool Functions

// modifyFirstToken
//   - removes the Head node if it is some kind of AddToken, and otherwise,
//   - adds as the Head node a NumToken if it is some kind of SubToken
func modifyFirstToken(list *plugin.TokenList) {
	if isKind(list.Head(), data.AddToken) {
		list.Disconnect(list.Head())
		return
	}

	if isKind(list.Head(), data.SubToken) {
		list.Connect(nil, data.NewNumToken("0"))
	}
}

// canBeAddedAsterisk returns true if an asterisk can be added
//
//	)( => )*(
func canBeAddedAsterisk(node *plugin.TokenNode) bool {
	if !isKind(node, data.RightToken) {
		return false
	}
	return isKind(node.Next(), data.LeftToken)
}

// canBeAddedZero returns true if an zero can be added
//
//	(- => (0-
func canBeAddedZero(node *plugin.TokenNode) bool {
	if !isKind(node, data.LeftToken) {
		return false
	}
	return isKind(node.Next(), data.SubToken)
}

// canBeRemovedPlus returns true if a AddToken can be removed
//
//	From: #+n, #+π, #+(, #+√n, #+√π, #+√(...)
//	To #n, #π, #(, #√n, #√π, #√(...)
func canBeRemovedPlus(node *plugin.TokenNode) bool {
	if !isKindFnRaw(node, data.IsSpecialToken) {
		if !isKindRaw(node, data.LeftToken) {
			return false
		}
	}

	temp := node.Next()
	if !isKind(temp, data.AddToken) {
		return false
	}

	temp = temp.Next()

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

// addParenthesesIfPossible adds parentheses as follows:
//
//	# = {%, *, +, -, /, ^, √}
//	From: #-n, #-π, #-(, #-√n, #-√π, #-√(...)
//	To #(-n), #(-π), #(-(...)), #(-√n) #(-√π) #(-√(...))
func addParenthesesIfPossible(i int, node *plugin.TokenNode, list *plugin.TokenList) {
	if !isKindFnRaw(node, data.IsSpecialToken) {
		return
	}

	temp := node.Next()

	if !isKind(temp, data.SubToken) {
		return
	}

	temp = temp.Next()

	// #-n, #-π
	if isKindFn(temp, data.IsNumPiToken) {
		addParentheses(node, 4, list)
		return
	}

	// #-(...)
	if isKind(temp, data.LeftToken) {
		addParenthesesInLoop(i, node, list)
		return
	}

	// #-√
	if isKind(temp, data.RootToken) {
		temp = temp.Next()

		// #-√n, #-√π
		if isKindFn(temp, data.IsNumPiToken) {
			addParentheses(node, 5, list)
			return
		}

		// #-√(...)
		if isKind(temp, data.LeftToken) {
			addParenthesesInLoop(i, node, list)
			return
		}
	}
}

// addParenthesesInLoop add parentheses wrapping a sign and another operation with parentheses
func addParenthesesInLoop(i int, node *plugin.TokenNode, list *plugin.TokenList) {
	var nLeft, nRight int

	for temp := node; temp != nil; i, temp = i+1, temp.Next() {
		if isKindRaw(temp, data.LeftToken) {
			nLeft++

		} else if isKindRaw(temp, data.RightToken) {
			nRight++

			if nLeft == nRight {
				list.Connect(node, data.NewLeftToken())
				list.Connect(temp, data.NewRightToken())
				break
			}
		}
	}
}

// isNextToo returns true if the first rune of the expression is float, otherwise returns false
func isNextToo(i int, expression string) bool {
	if i < len(expression)-1 {
		return data.IsFloat(rune(expression[i+1]))
	}
	return false
}

// isKind returns true if node can be token, otherwise returns false
func isKind(node *plugin.TokenNode, token data.TokenKind) bool {
	if node != nil {
		return isKindRaw(node, token)
	}
	return false
}

// isKindFn returns true if node can be token, otherwise returns false
func isKindFn(node *plugin.TokenNode, tokenFn func(token data.TokenKind) bool) bool {
	if node != nil {
		return tokenFn(node.Token().Kind())
	}
	return false
}

// isKindRaw returns true if node can be token, otherwise returns false
func isKindRaw(node *plugin.TokenNode, token data.TokenKind) bool {
	return node.Token().Kind() == token
}

// isKindFn returns true if node can be token, otherwise returns false
func isKindFnRaw(node *plugin.TokenNode, tokenFn func(token data.TokenKind) bool) bool {
	return tokenFn(node.Token().Kind())
}

// addParentheses adds a left token the of the node and a right token at position k afther node
func addParentheses(node *plugin.TokenNode, k int, list *plugin.TokenList) {
	list.Connect(node, data.NewLeftToken())
	list.ConnectFrom(node, k, data.NewRightToken())
}
