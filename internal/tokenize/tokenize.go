package tokenize

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
	"github.com/brianlewyn/go-calculator/internal/tool"
)

// tokenize represent the tokenized linked list
type tokenize struct {
	expression *string
	lenght     *int
}

// Tokenizer returns the expression in a linked list and nil,
// otherwise returns nil and an error
func Tokenizer(info *tool.Info) (*plugin.TokenList, data.Error) {
	expr := info.Expression()
	lenght := info.Lenght()

	tokenizer := tokenize{
		expression: expr,
		lenght:     lenght,
	}

	list, err := tokenizer.linkedList()
	if err != nil {
		return nil, data.NewError(*expr, err)
	}

	list, err = tokenizer.rebuild(list)
	if err != nil {
		return nil, data.NewError(*expr, err)
	}

	return list, nil
}

// linkedList returns an tokenized linked list and a possible error
func (t *tokenize) linkedList() (*plugin.TokenList, error) {
	list := plugin.NewTokenList()
	var value string

	if *t.lenght == 0 {
		return list, ierr.EmptyField
	}

	for i, r := range *t.expression {
		switch {

		// opeartors
		case data.IsMod(&r):
			list.Append(data.NewModToken())
		case data.IsMul(&r):
			list.Append(data.NewMulToken())
		case data.IsAdd(&r):
			list.Append(data.NewAddToken())
		case data.IsSub(&r):
			list.Append(data.NewSubToken())
		case data.IsDiv(&r):
			list.Append(data.NewDivToken())

		// parentheses
		case data.IsLeft(&r):
			list.Append(data.NewLeftToken())
		case data.IsRight(&r):
			list.Append(data.NewRightToken())

		// power
		case data.IsPow(&r):
			list.Append(data.NewPowToken())

		// pi or root
		case data.IsPi(&r):
			list.Append(data.NewPiToken())
		case data.IsRoot(&r):
			list.Append(data.NewRootToken())

		// numbers
		case data.IsFloat(&r):
			value += string(r)

			if !t.isFloat(i) {
				list.Append(data.NewNumToken(value))
				value = ""
			}

		default:
			if !data.IsGap(&r) {
				return nil, ierr.NewRune(r, i).Unknown()
			}
		}
	}

	*t.expression = ""
	*t.lenght = 0

	return list, nil
}

// rebuild returns a rebuilt tokenized linked list and a possible error
func (t tokenize) rebuild(list *plugin.TokenList) (*plugin.TokenList, error) {
	if list.IsEmpty() {
		return nil, ierr.EmptyField
	}

	for i, temp := 0, list.Head(); temp != nil; i, temp = i+1, temp.Next() {
		if canBeAddedAsterisk(temp) {
			list.Insert(i+1, data.NewMulToken())

		} else if canBeAddedZero(temp) {
			list.Insert(i+1, data.NewNumToken("0"))

		} else if canBeRemovedPlus(temp) {
			list.Remove(i + 1)

		} else {
			addParenthesesIfPossible(i, temp, list)
		}
	}

	if isKindPlus(list.Head()) {
		list.Remove(0)
	}

	return list, nil
}

// !Tool Methods

// isFloat return true if the first rune of the expression is float, otherwise returns false
func (t tokenize) isFloat(i int) bool {
	var after rune

	if *t.lenght < 1 || *t.lenght-1 <= i {
		after = data.Jocker
	} else {
		after = rune((*t.expression)[i+1])
	}

	return data.IsFloat(&after)
}

// !Tool Functions

// canBeAddedAsterisk returns true if an asterisk can be added
// )( => )*(
func canBeAddedAsterisk(node *plugin.TokenNode) bool {
	if !isKind(node, data.RightToken) {
		return false
	}
	return isKind(node.Next(), data.LeftToken)
}

// canBeAddedZero returns true if an zero can be added
// (-	=> (0-
func canBeAddedZero(node *plugin.TokenNode) bool {
	if !isKind(node, data.LeftToken) {
		return false
	}
	return isKind(node.Next(), data.SubToken)
}

// isKindPlus returns true if a AddToken can be removed
func isKindPlus(node *plugin.TokenNode) bool {
	return isKind(node, data.AddToken)
}

// canBeRemovedPlus returns true if a AddToken can be removed
//
// From: #+n, #+π, #+(, #+√n, #+√π, #+√(...)
//
// To #n, #π, #(, #√n, #√π, #√(...)
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
// # := {%, *, +, -, /, ^, √}
//
// From: #-n, #-π, #-(, #-√n, #-√π, #-√(...)
//
// To #(-n), #(-π), #(-(...)), #(-√n) #(-√π) #(-√(...))
func addParenthesesIfPossible(i int, node *plugin.TokenNode, list *plugin.TokenList) bool {
	if !isKindFnRaw(node, data.IsSpecialToken) {
		return false
	}

	temp := node.Next()
	if !isKind(temp, data.SubToken) {
		return false
	}

	temp = temp.Next()

	// #-n, #-π
	if isKindFn(temp, data.IsNumPiToken) {
		addParentheses(i+1, i+4, list)
		return true
	}

	// #-(...)
	if isKind(temp, data.LeftToken) {
		addParenthesesInLoop(i, node, list)
		return true
	}

	// #-√
	if isKind(temp, data.RootToken) {
		temp = temp.Next()

		// #-√n, #-√π
		if isKindFn(temp, data.IsNumPiToken) {
			addParentheses(i+1, i+5, list)
			return true
		}

		// #-√(...)
		if isKind(temp, data.LeftToken) {
			addParenthesesInLoop(i, node, list)
			return true
		}
	}

	return false
}

// addParenthesesInLoop add parentheses wrapping a sign and another operation with parentheses
func addParenthesesInLoop(i int, node *plugin.TokenNode, list *plugin.TokenList) {
	var nLeft, nRight int
	var flag bool
	j := i

	for temp := node; temp != nil; j, temp = j+1, temp.Next() {
		if isKindRaw(temp, data.LeftToken) {
			nLeft++

		} else if isKindRaw(temp, data.RightToken) {
			nRight++

			if nLeft == nRight {
				flag = true
				break
			}
		}
	}

	if flag {
		addParentheses(i+1, j+1, list)
	}
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

// addParentheses adds a left token at position i and a right token at position j
func addParentheses(i, j int, list *plugin.TokenList) {
	list.Insert(i, data.NewLeftToken())
	list.Insert(j, data.NewRightToken())
}
