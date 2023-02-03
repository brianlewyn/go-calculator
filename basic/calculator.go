package basic

func Calculate(expr string) (float64, error) {
	calculator := New(expr)

	if !calculator.Calculate() {
		return 0, calculator.Error()
	}

	return calculator.Answer(), nil
}
