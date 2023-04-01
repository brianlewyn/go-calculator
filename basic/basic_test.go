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
			got, err := Calculate(tt.expr)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
			if err != nil {
				t.Logf("MyError:\n%s", err)
			}
		})
	}
}
