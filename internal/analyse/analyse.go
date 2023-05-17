package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
)

// Analyzer returns nil if the math expression has a correct sematic,
// otherwise returns an error
func Analyser(list *plugin.TokenList) error {
	nL, nR := new(int), new(int)
	var bug error

	bug = isFirstTokenCorrect(list.Head().Token())
	if bug != nil {
		return bug
	}

	bug = isLastTokenCorrect(list.Tail().Token())
	if bug != nil {
		return bug
	}

	for temp := list.Head(); temp != nil; temp = temp.Next() {
		bug = canBeTogether(temp, temp.Next())
		if bug != nil {
			return bug
		}

		bug = isNumTokenCorrect(temp.Token())
		if bug != nil {
			return bug
		}

		bug = areCorrectParentheses(temp, list.Tail(), nL, nR)
		if bug != nil {
			return bug
		}
	}

	return nil
}

// !Tool Functions

// isFirstTokenCorrect returns nil if the first Token in the list is a correct Token to be first
// otherwise returns an error
func isFirstTokenCorrect(token data.Token) error {
	if data.IsFirstToken(token.Kind()) {
		return nil
	}
	return ierr.KindStart(data.ToRune(token.Kind()))
}

// isLastTokenCorrect returns nil if the last Token int the list is a correct Token to be last,
// otherwise returns an error
func isLastTokenCorrect(token data.Token) error {
	if data.IsLastToken(token.Kind()) {
		return nil
	}
	return ierr.KindEnd(data.ToRune(token.Kind()))
}

// isNumTokenCorrect returns nil is the number is correct, otherwise returns an error
func isNumTokenCorrect(token data.Token) error {
	if token.Kind() != data.NumToken {
		return nil
	}

	num := token.(data.Number).Value()
	if isAbsurdDot(num) {
		return ierr.NumberMisspelled(num)
	}

	lock, limit := new(bool), new(uint16)

	for _, r := range num {
		if r == data.Dot {
			if hasDotBeenRepeated(lock) {
				return ierr.NumberMisspelled(num)
			}
		}

		if hasLimitBeenExceeded(limit) {
			return ierr.NumberLimit(num)
		}
	}

	return nil
}

// isAbsurdDot returns true if there is a dot in absurd position,
// otherwise returns false
func isAbsurdDot(num string) bool {
	return data.Dot == rune(num[len(num)-1])
}

// hasDotBeenRepeated returns true if the dot has been repeated,
// otherwise returns false
func hasDotBeenRepeated(lock *bool) bool {
	if !*lock {
		*lock = true
		return false
	}
	return true
}

// hasLimitBeenExceeded returns true if the digit limit has been exceeded
// otherwise returns false
func hasLimitBeenExceeded(limit *uint16) bool {
	*limit++
	return *limit >= data.DigitLimit
}

// canBeTogether returns nil if there are not duplicate kinds,
// otherwise returns an error
func canBeTogether(current, next *plugin.TokenNode) error {
	if next == nil {
		return nil
	}

	kc1 := current.Token().Kind()
	kn2 := next.Token().Kind()

	beTogether := data.CanTokensBeTogether(kc1, kn2)
	if beTogether {
		return nil
	}

	return ierr.KindNotTogether(data.ToRune(kc1), data.ToRune(kn2))
}

// areCorrectParentheses returns nil if the number of parentheses is correct, otherwise returns an error
func areCorrectParentheses(current, tail *plugin.TokenNode, nLeft, nRight *int) error {
	isParanthesesTokenIncrease(current.Token().Kind(), nLeft, nRight)
	if *current != *tail {
		return nil
	}
	return areNLeftEqualToNRight(nLeft, nRight)
}

// isParanthesesTokenIncrease if is a LeftToken increases nLeft or a RightToken increases nRight
func isParanthesesTokenIncrease(kind data.TokenKind, nLeft, nRight *int) {
	if kind == data.LeftToken {
		*nLeft++
		return
	}
	if kind == data.RightToken {
		*nRight++
	}
}

// areNLeftEqualToNRight returns nil if both are equals,
// but if nLeft is greater then nRight return an error,
// otherwise return an error for right
func areNLeftEqualToNRight(nLeft, nRight *int) error {
	if *nLeft == *nRight {
		return nil
	}
	if *nLeft > *nRight {
		return ierr.IncompleteLeft
	}
	return ierr.IncompleteRight
}
