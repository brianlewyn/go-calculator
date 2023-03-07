package tokenize

import (
	"testing"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
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

		wantList := plugin.NewTokenList()
		// (0-1+2*3/4^5%6+√π)
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("0"))
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewNumToken("1"))
		wantList.Append(data.NewAddToken())
		wantList.Append(data.NewNumToken("2"))
		wantList.Append(data.NewMulToken())
		wantList.Append(data.NewNumToken("3"))
		wantList.Append(data.NewDivToken())
		wantList.Append(data.NewNumToken("4"))
		wantList.Append(data.NewPowToken())
		wantList.Append(data.NewNumToken("5"))
		wantList.Append(data.NewModToken())
		wantList.Append(data.NewNumToken("6"))
		wantList.Append(data.NewAddToken())
		wantList.Append(data.NewRootToken())
		wantList.Append(data.NewPiToken())
		wantList.Append(data.NewRightToken())
		// *(-1.234)
		wantList.Append(data.NewMulToken())
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("0"))
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewNumToken("1.234"))
		wantList.Append(data.NewRightToken())

		areEqualList(t, gotList, wantList)
	})
}
