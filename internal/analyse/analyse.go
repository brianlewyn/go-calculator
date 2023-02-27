package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
)

// analyse represents a parser for expression
type analyse struct {
	list *[]*data.Token
}

// isCorrect returns nil if the syntax is correct, otherwise returns an error
func (a analyse) isCorrect() error {
	var bug error
	var nL, nR int

	for temp := (*a.list)[0].Head(); temp.Next() != nil; temp = temp.Next() {
		bug = a.isCorrectFirst(temp)
		if bug != nil {
			return bug
		}

		bug = a.isCorrectLast(temp)
		if bug != nil {
			return bug
		}

		bug = isCorrectNumber(temp)
		if bug != nil {
			return bug
		}

		bug = canBeTogether(temp, temp.Next())
		if bug != nil {
			return bug
		}

		bug = areCorrectParentheses(&nL, &nR, temp, (*a.list)[0].Tail())
		if bug != nil {
			return bug
		}
	}

	return nil
}

// !Tool Methods

// isCorrectFirst returns nil is the number is correct, otherwise returns an error
func (a analyse) isCorrectFirst(token *data.Token) error {
	if token != (*a.list)[0].Head() {
		return nil
	}

	if !data.IsFirst(token.Kind()) {
		return ierr.StartKind
	}

	return nil
}

// isCorrectLast returns nil is the number is correct, otherwise returns an error
func (a analyse) isCorrectLast(token *data.Token) error {
	if token == (*a.list)[0].Tail() {
		return nil
	}

	if !data.IsLast(token.Kind()) {
		return ierr.EndKind
	}

	return nil
}

// !Tool Functions

// isCorrectNumber returns nil is the number is correct, otherwise returns an error
func isCorrectNumber(token *data.Token) error {
	if token.Kind() != data.NumToken {
		return nil
	}

	var number = token.Value()

	if !isAbsurdDot(number) {
		return ierr.NewNumber(*number).Misspelled()
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
				return ierr.NewNumber(*number).Misspelled()
			}
			flagDot = true
		}

		if nDigit++; nDigit == data.DigitLimit {
			return ierr.NewNumber(*number).Limit()
		}
	}

	return nil
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

// canBeTogether returns nil if there are not duplicate kinds, otherwise returns an error
func canBeTogether(token1, token2 *data.Token) error {
	if token2 == nil {
		return nil
	}

	kind1, kind2 := token1.Kind(), token2.Kind()
	beTogether := data.CanBeTogether(kind1, kind2)

	if !beTogether {
		return ierr.NewKind(format(token1, token2)).NotTogether()
	}

	return nil
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

// areCorrectParentheses returns nil if the number of parentheses is correct, otherwise returns an error
func areCorrectParentheses(nLeft, nRight *int, current, last *data.Token) error {

	if current.Kind() == data.LeftToken {
		*nLeft++
	} else if current.Kind() == data.RightToken {
		*nRight++
	}

	if *current != *last {
		return nil
	}

	if *nLeft == *nRight {
		return nil
	}

	if *nLeft > *nRight {
		return ierr.IncompleteLeft
	}

	return ierr.IncompleteRight
}
