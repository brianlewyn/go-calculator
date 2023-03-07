package tokenize

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	d "github.com/brianlewyn/go-linked-list/doubly"
)

// tokenize represent the tokenized linked list
type tokenize struct {
	expression *string
	lenght     *int
}

// linkedList returns an tokenized linked list and a possible error
func (t *tokenize) linkedList() (*d.Doubly[*data.Token], error) {
	list := d.NewDoubly[*data.Token]()
	var value string

	if *t.lenght == 0 {
		return list, ierr.EmptyField
	}

	for *t.lenght > 0 {
		r := rune((*t.expression)[0])

		switch {

		// opeartors
		case data.IsMod(&r):
			list.Append(d.NewNode(data.NewModToken()))
		case data.IsMul(&r):
			list.Append(d.NewNode(data.NewMulToken()))
		case data.IsAdd(&r):
			list.Append(d.NewNode(data.NewAddToken()))
		case data.IsSub(&r):
			list.Append(d.NewNode(data.NewSubToken()))
		case data.IsDiv(&r):
			list.Append(d.NewNode(data.NewDivToken()))

		// parentheses
		case data.IsLeft(&r):
			list.Append(d.NewNode(data.NewLeftToken()))
		case data.IsRight(&r):
			list.Append(d.NewNode(data.NewRightToken()))

		// power
		case data.IsPow(&r):
			list.Append(d.NewNode(data.NewPowToken()))

		// special: pi or root
		case data.IsSpecial(&r):

			if t.isPi(&r) {
				list.Append(d.NewNode(data.NewPiToken()))
				continue

			} else if t.isRoot(&r) {
				list.Append(d.NewNode(data.NewRootToken()))
				continue
			}

		// numbers
		case data.IsFloat(&r):
			value += string(r)

			if !t.isFloat() {
				list.Append(d.NewNode(data.NewNumToken(value)))
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

	return list, nil
}

// rebuild returns a rebuilt tokenized linked list and a possible error
func (t tokenize) rebuild(list *d.Doubly[*data.Token]) (*d.Doubly[*data.Token], error) {
	if list.IsEmpty() {
		return nil, ierr.EmptyField
	}

	for i, temp := 1, list.Head(); temp.Next() != nil; i, temp = i+1, temp.Next() {
		if canBeAddedAsterisk(temp) {
			list.Insert(i, d.NewNode(data.NewMulToken()))

		} else if canBeAddedZero(temp) {
			list.Insert(i, d.NewNode(data.NewNumToken("0")))
		}
	}

	return list, nil
}

// !Tool Methods

// isPi return true if r is pi number, otherwise returns false
func (t *tokenize) isPi(r *rune) bool {
	if *t.lenght < 2 {
		return false
		// ^ I didn't find a value to test it, but I leave this line just in case
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
		// ^ I didn't find a value to test it, but I leave this line just in case
	}

	second := rune((*t.expression)[1])
	last := rune((*t.expression)[2])

	if data.IsRoot(r, &second, &last) {
		*t.expression = (*t.expression)[3:]
		*t.lenght -= 3
		return true
	}

	return false
	// ^ I didn't find a value to test it, but I leave this line just in case
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
func canBeAddedAsterisk(node *d.Node[*data.Token]) bool {
	if node.Data().Kind() == data.RightToken {
		return node.Next().Data().Kind() == data.LeftToken
	}
	return false
}

// canBeAddedZero returns true if an zero can be added
func canBeAddedZero(node *d.Node[*data.Token]) bool {
	if node.Data().Kind() == data.LeftToken {
		switch node.Next().Data().Kind() {
		case data.AddToken:
		case data.SubToken:
		default:
			return false
		}
		return true
	}
	return false
}
