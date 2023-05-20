package analyse

import (
	"fmt"
	"strings"
	"testing"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/tokenize"
	"github.com/brianlewyn/go-linked-list/v2/doubly"
	"github.com/stretchr/testify/assert"
)

func TestAnalyser(t *testing.T) {
	test := []struct {
		name string
		list *doubly.Doubly[data.Token]
		as   ierr.KindOf
		is   error
	}{
		{
			name: "Bug: First: The first element",
			list: toList("%"),
			as:   ierr.CtxKindStart,
			// Try these: % * + - / ) ^
			// All except these: 0-9, (, ., π, √
		},
		{
			name: "Bug: Last: The last element",
			list: toList("(("),
			as:   ierr.CtxKindEnd,
			// Try these: (( √√ (√ √( 0+ 0^ 0+√ 0+(
			// All except these: 0-9, ), π
		},
		{
			name: "Bug: Number: An absurd number",
			as:   ierr.CtxNumberMisspelled,
			list: toList("1234."),
			// Try these: . 0. ...
		},
		{
			name: "Bug: Number: Almost number",
			as:   ierr.CtxNumberMisspelled,
			list: toList("1.234.5"),
			// Try these: .. .0. .0.0. ...
		},
		{
			name: "Bug: Number: Digit limit",
			list: toList(strings.Repeat("0", int(data.DigitLimit))),
			as:   ierr.CtxNumberLimit,
			// Try these: Any number with or more than 617 digits
		},
		{
			name: "Bug: Together: Elements together",
			list: toList("0^^0"),
			as:   ierr.CtxKindNotTogether,
			// # := {%, *, +, -, /, ^, √}
			// Try these: %% ** // )) ^^  %/ %) %^ () ...
			// All except these: #± ...
		},
		{
			name: "Bug: Together: Numbers together (1)",
			list: toList("π3.14"),
			as:   ierr.CtxKindNotTogether,
			// Try these: A number pi next to any number
		},
		{
			name: "Bug: Together: Numbers together (2)",
			list: toList("3.14π"),
			as:   ierr.CtxKindNotTogether,
			// Try these: Any number next to number pi
		},
		{
			name: "Bug: Parentheses: Left parentheses",
			is:   ierr.IncompleteLeft,
			list: toList("(0"),
			// Try these: (0 ((0) ...
		},
		{
			name: "Bug: Parentheses: Right parentheses",
			is:   ierr.IncompleteRight,
			list: toList("0)"),
			// Try these: 0) (0)) ...
		},
		{
			name: "NotBug: Expression",
			list: toList("(0.5 + 4.5 - 1) * 10 * √(7-2) / 4^2"),
			// Try this with a correct expression
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			bug := Analyser(tt.list)
			if bug != nil {
				t.Logf("Error:\n%s", bug)
			}

			if tt.as != "" {
				assert.Truef(ierr.As(bug, tt.as), "[error != As]: %v", bug)
				return
			}
			if tt.is != nil {
				assert.ErrorIsf(bug, tt.is, "[error != Is]: %v", bug)
				return
			}
			assert.Nilf(bug, "[error != Nil]: %v", bug)
		})
	}
}

// toList returns the expression in a raw Tokenized Linked List
func toList(expression string) *doubly.Doubly[data.Token] {
	list, err := tokenize.Tokenizer(expression)
	if err != nil {
		fmt.Printf("ERROR: %s\n\n", err)
	}
	return list
}

// go test -bench=BenchmarkAnalyser -benchmem -count=10 -benchtime=100x >> bench.txt
func BenchmarkAnalyser(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Analyser(toList("(0.5 + 4.5 - 1) * 10 * √(6-2) / 4^2"))
	}
}
