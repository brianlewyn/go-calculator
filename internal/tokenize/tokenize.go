package tokenize

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
)

// tokenize represent the tokenized linked list
type tokenize struct {
	expression *string
	lenght     *int
}

// Tokenizer tokenizes the expression in a linked list,
//
// but while creating the list, the expression is removed
func Tokenizer(expData *data.ExpData) (*plugin.TokenList, data.IErrData) {
	exp := expData.Expression()
	lenght := expData.Lenght()

	tokenizer := tokenize{
		expression: exp,
		lenght:     lenght,
	}

	list, err := tokenizer.linkedList()
	if err != nil {
		return nil, data.NewErrData(*exp, err)
	}

	list, err = tokenizer.rebuild(list)
	if err != nil {
		return nil, data.NewErrData(*exp, err)
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

		} else {
			addParenthesesIfPossible(i, temp, list)
		}
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
func canBeAddedAsterisk(node *plugin.TokenNode) bool {
	if node.Next() == nil {
		return false
	}

	if node.Token().Kind() != data.RightToken {
		return false
	}

	return node.Next().Token().Kind() == data.LeftToken
}

// canBeAddedZero returns true if an zero can be added
// (± => (0±
func canBeAddedZero(node *plugin.TokenNode) bool {
	if node.Next() == nil {
		return false
	}

	if node.Token().Kind() != data.LeftToken {
		return false
	}

	return data.IsAddSubToken(node.Next().Token().Kind())
}

// addParenthesesIfPossible adds parentheses as follows:
//
// From: ^±n, ^±π, ^±(, ^±√n, ^±√π, ^±√(...)
//
// To ^(±n), ^(±π), ^(±(...)), ^(±√n) ^(±√π) ^(±√(...))
func addParenthesesIfPossible(i int, node *plugin.TokenNode, list *plugin.TokenList) bool {
	if !isKind(node, data.PowToken) {
		return false
	}

	temp := node.Next()
	if !isKindFn(temp, data.IsAddSubToken) {
		return false
	}

	temp = temp.Next()

	// ^±n, ^±π
	if isKindFn(temp, data.IsNumPiToken) {
		addParentheses(i+1, i+4, list)
		return true
	}

	// ^±(...)
	if isKind(temp, data.LeftToken) {
		addParenthesesInLoop(i, node, list)
		return true
	}

	// ^±√
	if isKind(temp, data.RootToken) {
		temp = temp.Next()

		// ^±√n, ^±√π
		if isKindFn(temp, data.IsNumPiToken) {
			addParentheses(i+1, i+4, list)
			return true
		}

		// ^±√(...)
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
		if isKindButNodeIsNil(temp, data.LeftToken) {
			nLeft++

		} else if isKindButNodeIsNil(temp, data.RightToken) {
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
		return isKindButNodeIsNil(node, token)
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

// isKindButNodeIsNil returns true if node can be token, otherwise returns false
func isKindButNodeIsNil(node *plugin.TokenNode, token data.TokenKind) bool {
	return node.Token().Kind() == token
}

// addParentheses adds a left token at position i and a right token at position j
func addParentheses(i, j int, list *plugin.TokenList) {
	list.Insert(i, data.NewLeftToken())
	list.Insert(j, data.NewRightToken())
}
