package correct

func (a analyse) inconsistentNumbers() {}

func (a analyse) inconsistentOperators() {}

func (a analyse) inconsistentParentheses() {}

func (a analyse) inconsistentBlankSpakces() {}

func (a analyse) inconsistentDecimalPoints() {}

func (a *analyse) IsCorrectMathExpr() bool {

	a.inconsistentNumbers()
	a.inconsistentOperators()
	a.inconsistentParentheses()
	a.inconsistentBlankSpakces()
	a.inconsistentDecimalPoints()

	return false
}
