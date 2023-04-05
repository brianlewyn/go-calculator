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
		bug  ierr.KindOf
		nan  bool
	}{
		// TODO: Add test cases
		{
			name: "Root Square: NaN",
			expr: "√-2",
			want: math.Sqrt(-2),
			nan:  true,
		},
		{
			name: "Power: NaN",
			expr: "(-2)^(1/2)",
			want: math.Pow(-2, 1.0/2),
			nan:  true,
		},
		{
			name: "Tokenizer: Bug",
			expr: "#π3.14",
			bug:  ierr.Rune_Unknown,
		},
		{
			name: "Analyser: Bug",
			expr: "π3.14",
			bug:  ierr.Kind_Together,
		},
		{
			name: "Pi number & Multiplication",
			expr: "π * 2",
			want: math.Pi * 2,
		},
		{
			name: "Root Square",
			expr: "√2",
			want: math.Sqrt(2),
		},
		{
			name: "Power",
			expr: "2^2",
			want: math.Pow(2, 2),
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
			name: "first test",
			expr: "(0.5 + 4.5 - 1) * 10 * √(6-2) / 4^2",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.expr)

			if tt.bug != "" {
				assert.Truef(t, ierr.As(err.Bug(), tt.bug), "Bug != %v", tt.bug)
				if err != nil {
					t.Logf("Error:\n%s", err)
				}

			} else if tt.nan {
				assert.True(t, math.IsNaN(got), "got != NaN")
				if err != nil {
					t.Logf("Error:\n%s", err)
				}

			} else {
				assert.Equalf(t, got, tt.want, "got: %v, want: %v", got, tt.want)
				if err != nil {
					t.Logf("Error:\n%s", err)
				}
			}
		})
	}
}
