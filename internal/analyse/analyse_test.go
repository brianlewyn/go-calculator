package analyse

import (
	"strings"
	"testing"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
	"github.com/brianlewyn/go-calculator/internal/tokenize"
	"github.com/brianlewyn/go-calculator/internal/tool"
	"github.com/stretchr/testify/assert"
)

func TestAnalyser(t *testing.T) {
	assert := assert.New(t)

	type ErrKind uint8
	type Bug struct {
		kind ErrKind
		as   ierr.KindOf
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
			Bug:  Bug{kind: As, as: ierr.Kind_Start},
			list: expression("%"),
			// Try these: % * + - / ) ^
			// All except these: 0-9, (, ., π, √
		},
		{
			name: "Bug: Last: The last element",
			Bug:  Bug{kind: As, as: ierr.Kind_End},
			list: expression("(("),
			// Try these: (( √√ (√ √( 0+ 0^ 0+√ 0+(
			// All except these: 0-9, ), π
		},
		{
			name: "Bug: Number: An absurd number",
			Bug:  Bug{kind: As, as: ierr.Number_Misspelled},
			list: expression("1234."),
			// Try these: . 0. ...
		},
		{
			name: "Bug: Number: Almost number",
			Bug:  Bug{kind: As, as: ierr.Number_Misspelled},
			list: expression("1.234.5"),
			// Try these: .. .0. .0.0. ...
		},
		{
			name: "Bug: Number: Digit limit",
			Bug:  Bug{kind: As, as: ierr.Number_Limit},
			list: expression(strings.Repeat("0", int(data.DigitLimit))),
			// Try these: Any number with or more than 617 digits
		},
		{
			name: "Bug: Together: Elements together",
			Bug:  Bug{kind: As, as: ierr.Kind_Together},
			list: expression("0^^0"),
			// # := {%, *, +, -, /, ^, √}
			// Try these: %% ** // )) ^^  %/ %) %^ () ...
			// All except these: #± ...
		},
		{
			name: "Bug: Together: Numbers together (1)",
			Bug:  Bug{kind: As, as: ierr.Kind_Together},
			list: expression("π3.14"),
			// Try these: A number pi next to any number
		},
		{
			name: "Bug: Together: Numbers together (2)",
			Bug:  Bug{kind: As, as: ierr.Kind_Together},
			list: expression("3.14π"),
			// Try these: Any number next to number pi
		},
		{
			name: "Bug: Parentheses: Left parentheses",
			Bug:  Bug{kind: Is, is: ierr.IncompleteLeft},
			list: expression("(0"),
			// Try these: (0 ((0) ...
		},
		{
			name: "Bug: Parentheses: Right parentheses",
			Bug:  Bug{kind: Is, is: ierr.IncompleteRight},
			list: expression("0)"),
			// Try these: 0) (0)) ...
		},
		{
			name: "NotBug: Expression",
			Bug:  Bug{kind: Nil},
			list: expression("(0.5 + 4.5 - 1) * 10 * √(7-2) / 4^2"),
			// Try this with a correct expression
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := Analyser(tt.list)

			switch tt.kind {
			case As:
				assert.Truef(ierr.As(err.Bug(), tt.as), "Analyser() error != As: %v", err.Bug())
			case Is:
				assert.ErrorIsf(err.Bug(), tt.is, "Analyser() error != Is: %v", err.Bug())
			default:
				assert.Nilf(err, "Analyser() error != Nil: %v", err)
			}
		})
	}
}

func expression(expr string) *plugin.TokenList {
	list, _ := tokenize.Tokenizer(tool.NewInfo(&expr))
	return list
}
