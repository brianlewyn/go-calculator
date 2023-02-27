package tokenize

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// tokenize represent the tokenized linked list
type tokenize struct {
	expression *string
	lenght     *int
}

// linkedList returns an tokenized linked list and a possible error
func (t *tokenize) linkedList() (*[]*data.Token, error) {
	var list [](*data.Token)
	var value string

	for *t.lenght > 0 {
		r := rune((*t.expression)[0])

		switch {

		// opeartors
		case data.IsMod(&r):
			list[0] = data.NewModToken()
		case data.IsMul(&r):
			list[0] = data.NewMulToken()
		case data.IsAdd(&r):
			list[0] = data.NewAddToken()
		case data.IsSub(&r):
			list[0] = data.NewSubToken()
		case data.IsDiv(&r):
			list[0] = data.NewDivToken()

		// parentheses
		case data.IsLeft(&r):
			list[0] = data.NewLeftToken()
		case data.IsRight(&r):
			list[0] = data.NewRightToken()

		// power
		case data.IsPow(&r):
			list[0] = data.NewPowToken()

		// special: pi or root
		case data.IsSpecial(&r):

			if t.isPi(&r) {
				list[0] = data.NewPiToken()
				continue

			} else if t.isRoot(&r) {
				list[0] = data.NewRootToken()
				continue
			}

		// numbers
		case data.IsFloat(&r):
			value += string(r)

			if !t.isFloat() {
				list[0] = data.NewNumToken(value)
				value = data.Empty
			}

		default:
			if !data.IsGap(&r) {
				return nil, ierr.NewRune(r).Unknown()
			}
		}

		*t.expression = (*t.expression)[1:]
		*t.lenght--
	}

	*t.expression = data.Empty
	return &list, nil
}

// rebuild returns a rebuilt tokenized linked list and a possible error
func (t tokenize) rebuild(list *[]*data.Token) (*[]*data.Token, error) {
	if (*list)[0].IsEmpty() {
		return nil, ierr.EmptyField
	}

	for i, temp := -1, (*list)[0].Head(); temp.Next() != nil; temp = temp.Next() {
		if i++; canBeAddedAsterisk(temp) {
			(*list)[0].Insert(i+1, data.NewMulToken())
		} else if canBeAddedZero(temp) {
			(*list)[0].Insert(i+1, data.NewNumToken("0"))
		}
	}

	return list, nil
}

// !Tool Methods

// isPi return true if r is pi number, otherwise returns false
func (t *tokenize) isPi(r *rune) bool {
	if *t.lenght < 2 {
		return false
	}

	last := rune((*t.expression)[1])

	if data.IsPi(r, &last) {
		*t.expression = (*t.expression)[2:]
		*t.lenght -= 2
		return true
	}

	return false
}

// isRoot return true if r is square root, otherwise returns false
func (t *tokenize) isRoot(r *rune) bool {
	if *t.lenght < 3 {
		return false
	}

	second := rune((*t.expression)[1])
	last := rune((*t.expression)[2])

	if data.IsRoot(r, &second, &last) {
		*t.expression = (*t.expression)[3:]
		*t.lenght -= 3
		return true
	}

	return false
}

// isFloat return true if the first rune of the expression is float, otherwise returns false
func (t tokenize) isFloat() bool {
	var after rune

	if *t.lenght >= 2 {
		after = rune((*t.expression)[1])
	} else {
		after = data.Jocker
	}

	return data.IsFloat(&after)
}

// !Tool Functions

// canBeAddedAsterisk returns true if an asterisk can be added
func canBeAddedAsterisk(token *data.Token) bool {
	if token.Kind() == data.RightToken {
		return token.Next().Kind() == data.LeftToken
	}
	return false
}

// canBeAddedZero returns true if an zero can be added
func canBeAddedZero(token *data.Token) bool {
	if token.Kind() == data.LeftToken {
		switch token.Next().Kind() {
		case data.AddToken:
		case data.SubToken:
		default:
			return false
		}
		return true
	}
	return false
}
