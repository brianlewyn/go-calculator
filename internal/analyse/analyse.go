package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// analyse represents a parser for expression
type analyse struct {
	list *[]*data.Token
}

// isCorrect returns true if the syntax is correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error
func (a analyse) isCorrect() bool {
	if thereIsAnyError() {
		return false
	}

	var nL, nR int
	for temp := (*a.list)[0].Head(); temp.Next() != nil; temp = temp.Next() {
		if !a.isCorrectFirst(temp) {
			return false
		}
		if !a.isCorrectLast(temp) {
			return false
		}
		if !isCorrectNumber(temp) {
			return false
		}
		if !canBeTogether(temp, temp.Next()) {
			return false
		}
		if !areCorrectParentheses(&nL, &nR, temp, (*a.list)[0].Tail()) {
			return false
		}
	}

	return true
}

// !Tool Methods

// isCorrectFirst returns true is the number is correct, otherwise returns false
//
// But if there is any error then the error is stored in data.Error
func (a analyse) isCorrectFirst(token *data.Token) bool {
	if token != (*a.list)[0].Head() {
		return true // go out
	}

	if !data.IsFirst(token.Kind()) {
		data.Error = ierr.StartKind
		return false
	}

	return true
}

// isCorrectLast returns true is the number is correct, otherwise returns false
//
// But if there is any error then the error is stored in data.Error
func (a analyse) isCorrectLast(token *data.Token) bool {
	if token == (*a.list)[0].Tail() {
		return true // go out
	}

	if !data.IsLast(token.Kind()) {
		data.Error = ierr.EndKind
		return false
	}

	return true
}

// !Tool Functions

// thereIsAnyError returns true if data.Error isn't nil, otherwise returns false
func thereIsAnyError() bool {
	return data.Error != nil
}

// isCorrectNumber returns true is the number is correct, otherwise returns false
//
// But if there is any error then the error is stored in data.Error
func isCorrectNumber(token *data.Token) bool {
	if token.Kind() != data.NumToken {
		return true // go out
	}

	var number = token.Value()

	if !isAbsurdDot(number) {
		data.Error = ierr.NewNumber(*number).Misspelled()
		return false
	}

	var flagDot bool
	var nDigit uint16

	for _, r := range *number {
		if !data.IsFloat(&r) {
			flagDot = false
			nDigit = 0
			continue
		}

		if data.IsDot(&r) {
			if flagDot {
				data.Error = ierr.NewNumber(*number).Misspelled()
				return false
			}
			flagDot = true
		}

		if nDigit++; nDigit == data.DigitLimit {
			data.Error = ierr.NewNumber(*number).Limit()
			return false
		}
	}

	return true
}

// isAbsurdDot returns true if there is a dot in absurd position, otherwise returns false
func isAbsurdDot(num *string) bool {
	n := len(*num)

	if data.Dot == rune((*num)[n-1]) {
		return true
	}

	if n > 3 {
		return false
	}

	switch *num {
	case ".":
	case "0.":
	case ".0":
	case "0.0":
	default:
		return false
	}

	return true
}

// canBeTogether returns true if there are duplicate kinds
//
// But if there is any error then the error is stored in data.Error
func canBeTogether(token1, token2 *data.Token) bool {
	if token2 == nil {
		return true // go out
	}

	kind1, kind2 := token1.Kind(), token2.Kind()
	beTogether := data.CanBeTogether(kind1, kind2)

	if !beTogether {
		data.Error = ierr.NewKind(format(token1, token2)).NotTogether()
	}

	return beTogether
}

// format returns the value of the kind
func format(token1, token2 *data.Token) (string, string) {
	var fmt1, fmt2 string

	rune1 := data.Kind(token1.Kind())
	rune2 := data.Kind(token2.Kind())

	if rune1 == data.Jocker {
		fmt1 = *token1.Value()
	} else {
		fmt1 = string(rune1)
	}

	if rune2 == data.Jocker {
		fmt2 = *token2.Value()
	} else {
		fmt2 = string(rune2)
	}

	return fmt1, fmt2
}

// areCorrectParentheses returns true if the number of parentheses is correct, otherwise returns false.
//
// But if there is any error then the error is stored in data.Error
func areCorrectParentheses(nLeft, nRight *int, current, last *data.Token) bool {

	if current.Kind() == data.LeftToken {
		*nLeft++
	} else if current.Kind() == data.RightToken {
		*nRight++
	}

	if *current != *last {
		return true
	}

	if *nLeft == *nRight {
		return true
	}

	if *nLeft > *nRight {
		data.Error = ierr.IncompleteLeft
		return false
	}

	data.Error = ierr.IncompleteRight
	return false
}
