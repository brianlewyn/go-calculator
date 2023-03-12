package analyse

import (
	"strings"
	"testing"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
	"github.com/brianlewyn/go-calculator/internal/tokenize"
	"github.com/stretchr/testify/assert"
)

func TestAnalyser(t *testing.T) {
	assert := assert.New(t)

	type ErrKind uint8
	type Bug struct {
		kind ErrKind
		as   interface{}
		is   error
	}

	const (
		_ = ErrKind(iota)
		As
		Is
		Nil
	)

	test := []struct {
		Bug
		name string
		list *plugin.TokenList
	}{
		{
			name: "Bug: First: The first element",
			Bug:  Bug{kind: Is, is: ierr.StartKind},
			list: Expr("%"),
			// Try these: % * + - / ) ^
			// All except these: 0-9, (, ., π, √
		},
		{
			name: "Bug: Last: The last element",
			Bug:  Bug{kind: Is, is: ierr.EndKind},
			list: Expr("(("),
			// Try these: (( √√ (√ √( 0+ 0^ 0+√ 0+(
			// All except these: 0-9, ), π
		},
		{
			name: "Bug: Number: An absurd number",
			Bug:  Bug{kind: As, as: new(ierr.Number)},
			list: Expr("1234."),
			// Try these: . 0. ...
		},
		{
			name: "Bug: Number: Almost number",
			Bug:  Bug{kind: As, as: new(ierr.Number)},
			list: Expr("1.234.5"),
			// Try these: .. .0. .0.0. ...
		},
		{
			name: "Bug: Number: Digit limit",
			Bug:  Bug{kind: As, as: new(ierr.Number)},
			list: Expr(strings.Repeat("0", int(data.DigitLimit))),
			// Try these: Any number with or more than 617 digits
		},
		{
			name: "Bug: Together: Elements together",
			Bug:  Bug{kind: As, as: new(ierr.Kind)},
			list: Expr("0++0)"),
			// Try these: %% ** ++ -- // )) ^^ %+ %- %/ %) %^ () ...
			// All except these: +0+, +.0, +π+, +√, 0^+0 ...
		},
		{
			name: "Bug: Together: Numbers together (1)",
			Bug:  Bug{kind: As, as: new(ierr.Kind)},
			list: Expr("π3.14"),
			// Try these: A number pi next to any number
		},
		{
			name: "Bug: Together: Numbers together (2)",
			Bug:  Bug{kind: As, as: new(ierr.Kind)},
			list: Expr("3.14π"),
			// Try these: Any number next to number pi
		},
		{
			name: "Bug: Parentheses: Left parentheses",
			Bug:  Bug{kind: Is, is: ierr.IncompleteLeft},
			list: Expr("(0"),
			// Try these: (0 ((0) ...
		},
		{
			name: "Bug: Parentheses: Right parentheses",
			Bug:  Bug{kind: Is, is: ierr.IncompleteRight},
			list: Expr("0)"),
			// Try these: 0) (0)) ...
		},
		{
			name: "NotBug: Expression",
			Bug:  Bug{kind: Nil},
			list: Expr("(0.5 + 4.5 - 1) * 10 * √(7-2) / 4^2"),
			// Try this with a correct expression
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := Analyser(tt.list)

			switch tt.kind {
			case As:
				assert.ErrorAsf(err, &tt.as, "Analyser() error != As: %v", err)
			case Is:
				assert.ErrorIsf(err, tt.is, "Analyser() error != Is: %v", err)
			default:
				assert.Nilf(err, "Analyser() error != Nil: %v", err)
			}
		})
	}
}

func Expr(expr string) *plugin.TokenList {
	list, _ := tokenize.Tokenizer(data.New(&expr))
	return list
}
