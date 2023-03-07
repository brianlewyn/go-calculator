package tokenize

import (
	"testing"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	d "github.com/brianlewyn/go-linked-list/doubly"
	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	assert := assert.New(t)

	t.Run("From an empty expression to a linked list ", func(t *testing.T) {
		expression := ""
		gotList, err := Tokenizer(data.New(&expression))
		assert.ErrorIsf(err, ierr.EmptyField, "ierr.EmptyField != %v", err)
		assert.Nil(gotList, "gotList != nil")
	})

	t.Run("From an expression to a empty linked list ", func(t *testing.T) {
		expression := "  "
		gotList, err := Tokenizer(data.New(&expression))
		assert.ErrorIsf(err, ierr.EmptyField, "ierr.EmptyField != %v", err)
		assert.Nil(gotList, "gotList != nil")
	})

	t.Run("From an expression to a linked list", func(t *testing.T) {
		expression := "(0 - 1 + 2 * 3 / 4 ^ 5 % 6 + √π)(-1.234)"

		gotList, err := Tokenizer(data.New(&expression))
		assert.NoError(err, "Tokenizer(data) error != nil")

		wantList := d.NewDoubly[*data.Token]()

		// (0-1+2*3/4^5%6+√π)
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

		// *(-1.234)
		wantList.Append(d.NewNode(data.NewMulToken()))
		wantList.Append(d.NewNode(data.NewLeftToken()))
		wantList.Append(d.NewNode(data.NewNumToken("0")))
		wantList.Append(d.NewNode(data.NewSubToken()))
		wantList.Append(d.NewNode(data.NewNumToken("1.234")))
		wantList.Append(d.NewNode(data.NewRightToken()))

		areEqualList(t, gotList, wantList)
	})
}
