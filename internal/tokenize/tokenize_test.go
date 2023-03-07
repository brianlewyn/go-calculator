package tokenize

import (
	"reflect"
	"testing"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	d "github.com/brianlewyn/go-linked-list/doubly"
	"github.com/stretchr/testify/assert"
)

func Test_tokenize_linkedList(t *testing.T) {
	assert := assert.New(t)

	t.Run("From an empty expression to a list", func(t *testing.T) {
		expression := ""
		lenght := len(expression)

		tokenizer := tokenize{
			expression: &expression,
			lenght:     &lenght,
		}

		gotList, err := tokenizer.linkedList()
		assert.ErrorIsf(err, ierr.EmptyField, "ierr.EmptyField != %v", err)

		wantList := d.NewDoubly[*data.Token]()

		assert.Equal(gotList.Size(), wantList.Size(), "gotList.Size() != wantList()")
	})

	t.Run("From an expression with some inappropriate symbols to a list", func(t *testing.T) {
		expression := "12345 + hola + 12345"
		lenght := len(expression)

		tokenizer := tokenize{
			expression: &expression,
			lenght:     &lenght,
		}

		gotList, err := tokenizer.linkedList()

		errRune := new(ierr.Rune)
		assert.ErrorAsf(err, &errRune, "&ierr.Rune != %v", err)
		assert.Nil(gotList, "gotList != nil")
	})

	t.Run("From a filled expression to a list", func(t *testing.T) {
		expression := "(0 - 1 + 2 * 3 / 4 ^ 5 % 6 + √π) - 1.234"
		lenght := len(expression)

		tokenizer := tokenize{
			expression: &expression,
			lenght:     &lenght,
		}

		gotList, err := tokenizer.linkedList()
		assert.NoError(err, "tokenizer.linkedList() error != nil")

		wantList := d.NewDoubly[*data.Token]()
		wantList.Append(d.NewNode(data.NewLeftToken()))
		wantList.Append(d.NewNode(data.NewNumToken("0")))
		wantList.Append(d.NewNode(data.NewSubToken()))
		wantList.Append(d.NewNode(data.NewNumToken("1")))
		wantList.Append(d.NewNode(data.NewAddToken()))
		wantList.Append(d.NewNode(data.NewNumToken("2")))
		wantList.Append(d.NewNode(data.NewMulToken()))
		wantList.Append(d.NewNode(data.NewNumToken("3")))
		wantList.Append(d.NewNode(data.NewDivToken()))
		wantList.Append(d.NewNode(data.NewNumToken("4")))
		wantList.Append(d.NewNode(data.NewPowToken()))
		wantList.Append(d.NewNode(data.NewNumToken("5")))
		wantList.Append(d.NewNode(data.NewModToken()))
		wantList.Append(d.NewNode(data.NewNumToken("6")))
		wantList.Append(d.NewNode(data.NewAddToken()))
		wantList.Append(d.NewNode(data.NewRootToken()))
		wantList.Append(d.NewNode(data.NewPiToken()))
		wantList.Append(d.NewNode(data.NewRightToken()))
		wantList.Append(d.NewNode(data.NewSubToken()))
		wantList.Append(d.NewNode(data.NewNumToken("1.234")))

		node1 := gotList.Head()
		node2 := wantList.Head()

		for i := 0; node1 != nil && node2 != nil; i++ {
			token1 := node1.Data()
			token2 := node2.Data()

			if assert.EqualValues(token1.Kind(), token2.Kind(), "kind1 != kind2") {
				if token1.Kind() == data.NumToken {
					assert.EqualValues(*token1.Value(), *token2.Value(), "value1 != value2")
				}
			}

			node1 = node1.Next()
			node2 = node2.Next()
		}

		assert.Equal(gotList.Size(), wantList.Size(), "gotList.Size() != wantList()")
		assert.Equal(expression, "", "expression != \"\"")
	})
}

func Test_tokenize_rebuild(t *testing.T) {
	type fields struct {
		expression *string
		lenght     *int
	}
	type args struct {
		list *d.Doubly[*data.Token]
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *d.Doubly[*data.Token]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tokenize{
				expression: tt.fields.expression,
				lenght:     tt.fields.lenght,
			}
			got, err := tr.rebuild(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenize.rebuild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenize.rebuild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canBeAddedAsterisk(t *testing.T) {
	type args struct {
		node *d.Node[*data.Token]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeAddedAsterisk(tt.args.node); got != tt.want {
				t.Errorf("canBeAddedAsterisk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canBeAddedZero(t *testing.T) {
	type args struct {
		node *d.Node[*data.Token]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeAddedZero(tt.args.node); got != tt.want {
				t.Errorf("canBeAddedZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
