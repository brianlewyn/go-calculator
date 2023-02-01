package correct

func (b base) inconsistentNumbers() {}

func (b base) inconsistentOperators() {}

func (b base) inconsistentParentheses() {}

func (b base) inconsistentBlankSpakces() {}

func (b base) inconsistentDecimalPoints() {}

func (b base) isCorrectMathExpr() error {

	b.inconsistentNumbers()
	b.inconsistentOperators()
	b.inconsistentParentheses()
	b.inconsistentBlankSpakces()
	b.inconsistentDecimalPoints()

	return nil
}
