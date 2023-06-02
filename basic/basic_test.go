package basic

import (
	"math"
	"testing"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		expr string
		want float64
		as   ierr.KindOf
		is   error
	}{
		// TODO: Add test cases
		{
			name: "Root Square: NaN",
			expr: "√-2",
			want: math.Sqrt(-2),
			is:   ierr.ResultIsNaN,
		},
		{
			name: "Power: NaN",
			expr: "(-2)^(1/2)",
			want: math.Pow(-2, 1.0/2),
			is:   ierr.ResultIsNaN,
		},
		{
			name: "Tokenizer: Bug",
			expr: "#π3.14",
			as:   ierr.CtxRuneUnknown,
		},
		{
			name: "Analyser: Bug",
			expr: "π3.14",
			as:   ierr.CtxKindNotTogether,
		},
		{
			name: "Root Square",
			expr: "√2",
			want: math.Sqrt(2),
		},
		{
			name: "Double Root Square",
			expr: "√√2",
			want: math.Sqrt(math.Sqrt(2)),
		},
		{
			name: "Double Root Square with wrap",
			expr: "√√(2+4)",
			want: math.Sqrt(math.Sqrt(2 + 4)),
		},
		{
			name: "Triple Root Square",
			expr: "√√√2",
			want: math.Sqrt(math.Sqrt(math.Sqrt(2))),
		},
		{
			name: "Triple Root Square with wrap",
			expr: "√√√(2+4)",
			want: math.Sqrt(math.Sqrt(math.Sqrt(2 + 4))),
		},
		{
			name: "Power",
			expr: "2^2",
			want: math.Pow(2, 2),
		},
		{
			name: "Power: as in NaN but without parentheses",
			expr: "-2^(1/2)",
			want: -math.Pow(2, 1.0/2),
		},
		{
			name: "Pi number & Multiplication",
			expr: "π * 2",
			want: math.Pi * 2,
		},
		{
			name: "Multiplication & Division",
			expr: "30 / 3 * 5",
			want: 50,
		},
		{
			name: "Multiplication, Division & Parentheses",
			expr: "30 / (3 * 5)",
			want: 2,
		},
		{
			name: "Solve a complex expression",
			expr: "(0.5 + 4.5 - 1) * 10 * √(6-2) / 4^2",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, bug := Calculate(tt.expr)
			if bug != nil {
				t.Logf("Error:\n%s", bug)
			}

			switch {
			case tt.as != "":
				assert.Truef(t, ierr.As(bug, tt.as), "Bug != %v", tt.as)
			case tt.is != nil:
				assert.ErrorIs(t, bug, tt.is, "Bug != nil")
			default:
				assert.Equalf(t, got, tt.want, "got: %v, want: %v", got, tt.want)
			}
		})
	}
}

// go test -bench=BenchmarkCalculate -benchmem -count=10 -benchtime=100x >> bench.txt
func BenchmarkCalculate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate("(0.5 + 4.5 - 1) * 10 * √(6-2) / 4^2")
	}
}
