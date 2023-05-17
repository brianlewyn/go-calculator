package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
)

// Analyser returns nil if the math expression is correct,
// otherwise returns an error
func Analyser(list *plugin.TokenList) error {
	var bug error
	var nL, nR int

	for temp := list.Head(); temp != nil; temp = temp.Next() {
		if bug = isCorrectFirst(temp.Token(), list); bug != nil {
			return bug
		}

		if bug = isCorrectLast(temp.Token(), list); bug != nil {
			return bug
		}

		if bug = isCorrectNumber(temp.Token()); bug != nil {
			return bug
		}

		if bug = canBeTogether(temp, temp.Next()); bug != nil {
			return bug
		}

		bug = areCorrectParentheses(&nL, &nR, temp, list.Tail())
		if bug != nil {
			return bug
		}
	}

	return nil
}

// !Tool Methods

// isCorrectFirst returns nil is the number is correct, otherwise returns an error
func isCorrectFirst(token data.Token, list *plugin.TokenList) error {
	if token != list.Head().Token() {
		return nil
	}

	if !data.IsFirstToken(token.Kind()) {
		kind := data.ToRune(token.Kind())
		return ierr.KindStart(kind)
	}

	return nil
}

// isCorrectLast returns nil is the number is correct, otherwise returns an error
func isCorrectLast(token data.Token, list *plugin.TokenList) error {
	if token != list.Tail().Token() {
		return nil
	}

	if !data.IsLastToken(token.Kind()) {
		kind := data.ToRune(token.Kind())
		return ierr.KindEnd(kind)
	}

	return nil
}

// !Tool Functions

// isCorrectNumber returns nil is the number is correct, otherwise returns an error
func isCorrectNumber(token data.Token) error {
	if token.Kind() != data.NumToken {
		return nil
	}

	num := token.(data.Number).Value()
	if isAbsurdDot(num) {
		return ierr.NumberMisspelled(num)
	}

	var flagDot bool
	var nDigit uint16

	for _, r := range num {
		if r == data.Dot {
			if flagDot {
				return ierr.NumberMisspelled(num)
			}
			flagDot = true
		}

		if nDigit++; nDigit >= data.DigitLimit {
			return ierr.NumberLimit(num)
		}
	}

	return nil
}

// isAbsurdDot returns true if there is a dot in absurd position, otherwise returns false
func isAbsurdDot(num string) bool {
	return data.Dot == rune(num[len(num)-1])
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
		return ierr.KindNotTogether(data.ToRune(kind1), data.ToRune(kind2))
	}

	return nil
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
