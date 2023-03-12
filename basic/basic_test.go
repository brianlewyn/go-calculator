package basic

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		expr string
		want float64
	}{
		// TODO: Add test cases
		{
			name: "first test",
			expr: "(0.5 + 4.5 - 1) * 10 * √(7-2) / 4^^2",
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, expr, err := Calculate(tt.expr)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
			t.Logf("\n%v", err)
			t.Logf("\n%v", expr)
		})
	}
}
