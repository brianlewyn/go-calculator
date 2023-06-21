package analyse

import (
	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/doubly"
)

// Analyzer returns nil if the math expression has a correct sematic,
// otherwise returns an error
func Analyser(list *doubly.Doubly) error {
	nL, nR := new(int), new(int)

	err := isFirstTokenCorrect(list.Head().Token())
	if err != nil {
		return err
	}

	err = isLastTokenCorrect(list.Tail().Token())
	if err != nil {
		return err
	}

	for temp := list.Head(); temp != nil; temp = temp.Next() {
		err := canBeTogether(temp, temp.Next())
		if err != nil {
			return err
		}

		err = isNumTokenCorrect(temp.Token())
		if err != nil {
			return err
		}

		err = areCorrectParentheses(temp, list.Tail(), nL, nR)
		if err != nil {
			return err
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
	return ierr.KindStart(data.RuneMap[token.Kind()])
}

// isLastTokenCorrect returns nil if the last Token int the list is a correct Token to be last,
// otherwise returns an error
func isLastTokenCorrect(token data.Token) error {
	if data.IsLastToken(token.Kind()) {
		return nil
	}
	return ierr.KindEnd(data.RuneMap[token.Kind()])
}

// canBeTogether returns nil if there are not duplicate kinds,
// otherwise returns an error
func canBeTogether(current, next *doubly.Node) error {
	if next == nil {
		return nil
	}

	kCurr := current.Token().Kind()
	kNext := next.Token().Kind()

	if data.CanTokensBeTogether(kCurr, kNext) {
		return nil
	}

	return ierr.KindNotTogether(data.RuneMap[kCurr], data.RuneMap[kNext])
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

	unlock := true

	for i, r := range num {
		if r == data.Dot {
			if unlock = !unlock; unlock {
				return ierr.NumberMisspelled(num)
			}
		}

		if uint16(i+1) == data.DigitLimit {
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

// areCorrectParentheses returns nil if the number of parentheses is correct, otherwise returns an error
func areCorrectParentheses(current, tail *doubly.Node, nLeft, nRight *int) error {
	increaseIfIsParanthesesToken(current.Token().Kind(), nLeft, nRight)

	if *current != *tail {
		return nil
	}

	return areNLeftEqualToNRight(nLeft, nRight)
}

// increaseIfIsParanthesesToken if is a LeftToken increases nLeft or a RightToken increases nRight
func increaseIfIsParanthesesToken(kind data.TokenKind, nLeft, nRight *int) {
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
