package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
)

// analyse represents a parser for expression
type analyse struct {
	list *plugin.TokenList
}

// Analyser returns nil if the basic math expression is correct, otherwise returns an error
func Analyser(list *plugin.TokenList) error {
	return analyse{list: list}.isCorrect()
}

// isCorrect returns nil if the syntax is correct, otherwise returns an error
func (a analyse) isCorrect() error {
	var bug error
	var nL, nR int

	for temp := a.list.Head(); temp != nil; temp = temp.Next() {
		bug = a.isCorrectFirst(temp.Token())
		if bug != nil {
			return bug
		}

		bug = a.isCorrectLast(temp.Token())
		if bug != nil {
			return bug
		}

		bug = isCorrectNumber(temp.Token())
		if bug != nil {
			return bug
		}

		bug = canBeTogether(temp, temp.Next())
		if bug != nil {
			return bug
		}

		bug = areCorrectParentheses(&nL, &nR, temp, a.list.Tail())
		if bug != nil {
			return bug
		}
	}

	return nil
}

// !Tool Methods

// isCorrectFirst returns nil is the number is correct, otherwise returns an error
func (a analyse) isCorrectFirst(token *data.Token) error {
	if *token != *a.list.Head().Token() {
		return nil
	}

	if !data.IsFirstToken(token.Kind()) {
		return ierr.StartKind
	}

	return nil
}

// isCorrectLast returns nil is the number is correct, otherwise returns an error
func (a analyse) isCorrectLast(token *data.Token) error {
	if *token != *a.list.Tail().Token() {
		return nil
	}

	if !data.IsLastToken(token.Kind()) {
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

	if isAbsurdDot(number) {
		return ierr.NewNumber(*number).Misspelled()
	}

	var flagDot bool
	var nDigit uint16

	for _, r := range *number {
		if data.IsDot(&r) {
			if flagDot {
				return ierr.NewNumber(*number).Misspelled()
			}
			flagDot = true
		}

		if nDigit++; nDigit >= data.DigitLimit {
			return ierr.NewNumber(*number).Limit()
		}
	}

	return nil
}

// isAbsurdDot returns true if there is a dot in absurd position, otherwise returns false
func isAbsurdDot(num *string) bool {
	return data.Dot == rune((*num)[len(*num)-1])
}

// canBeTogether returns nil if there are not duplicate kinds, otherwise returns an error
func canBeTogether(curr, next *plugin.TokenNode) error {
	if next == nil {
		return nil
	}

	token1, token2 := curr.Token(), next.Token()
	kind1, kind2 := token1.Kind(), token2.Kind()

	beTogether := data.CanTokensBeTogether(kind1, kind2)

	if !beTogether {
		return ierr.NewKind(format(token1, token2)).NotTogether()
	}

	return nil
}

// format returns the value of the kind
func format(token1, token2 *data.Token) (string, string) {
	var fmt1, fmt2 string

	rune1 := data.FromTokenKindToRune(token1.Kind())
	rune2 := data.FromTokenKindToRune(token2.Kind())

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
func areCorrectParentheses(nLeft, nRight *int, curr, last *plugin.TokenNode) error {
	kind := curr.Token().Kind()

	if kind == data.LeftToken {
		*nLeft++
	} else if kind == data.RightToken {
		*nRight++
	}

	if *curr != *last {
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
