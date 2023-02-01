package correct

func (b base) numbers() {}

func (b base) operators() {}

func (b base) parentheses() {}

func (b base) blankSpakces() {}

func (b base) decimalPoints() {}

func (b base) isCorrectSyntax() error {

	b.numbers()
	b.operators()
	b.parentheses()
	b.blankSpakces()
	b.decimalPoints()

	return nil
}
