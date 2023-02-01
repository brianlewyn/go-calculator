package correct

type base struct {
	mathExpr *string
}

func IsCorrect(mathExpr *string) error {
	expr := base{mathExpr: mathExpr}

	if err := expr.isCorrectSyntax(); err != nil {
		return err
	}

	if err := expr.isCorrectMathExpr(); err != nil {
		return err
	}

	return nil
}
