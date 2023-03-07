package tokenize

import (
	"reflect"
	"testing"

	"github.com/brianlewyn/go-calculator/internal/data"
	d "github.com/brianlewyn/go-linked-list/doubly"
)

func TestTokenizer(t *testing.T) {
	type args struct {
		data *data.Data
	}
	tests := []struct {
		name    string
		args    args
		want    *d.Doubly[*data.Token]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tokenizer(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tokenizer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenizer() = %v, want %v", got, tt.want)
			}
		})
	}
}
