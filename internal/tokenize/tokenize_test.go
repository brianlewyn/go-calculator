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
		assert.Equal(gotList.Size(), 0, "gotList.Size() != 0")
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

		wantList := plugin.NewTokenList()
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
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewNumToken("1.234"))

		areEqualList(t, gotList, wantList)

		assert.Equal(gotList.Size(), wantList.Size(), "gotList.Size() != wantList()")
		assert.Equal(expression, "", "expression != \"\"")
	})
}

func Test_tokenize_rebuild(t *testing.T) {
	assert := assert.New(t)

	t.Run("From an empty list to a list rebuilded", func(t *testing.T) {
		expression := "   "
		lenght := len(expression)

		tokenizer := tokenize{
			expression: &expression,
			lenght:     &lenght,
		}

		gotList, err := tokenizer.linkedList()
		assert.NoError(err, "tokenizer.linkedList() error != nil")

		gotList, err = tokenizer.rebuild(gotList)
		assert.ErrorIsf(err, ierr.EmptyField, "ierr.EmptyField != %v", err)
		assert.Nil(gotList, "gotList != nil")
	})

	t.Run("From a list to a list rebuilded", func(t *testing.T) {
		expression := "(+10)(-12)(*12)"
		lenght := len(expression)

		tokenizer := tokenize{
			expression: &expression,
			lenght:     &lenght,
		}

		gotList, err := tokenizer.linkedList()
		assert.NoError(err, "tokenizer.linkedList() error != nil")

		gotList, err = tokenizer.rebuild(gotList)
		assert.NoError(err, "tokenizer.rebuild(gotList) error != nil")

		wantList := plugin.NewTokenList()
		// (0+10)
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("0"))
		wantList.Append(data.NewAddToken())
		wantList.Append(data.NewNumToken("10"))
		wantList.Append(data.NewRightToken())
		// *(0-12)
		wantList.Append(data.NewMulToken())
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("0"))
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewNumToken("12"))
		wantList.Append(data.NewRightToken())
		// *(*12)
		wantList.Append(data.NewMulToken())
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewMulToken())
		wantList.Append(data.NewNumToken("12"))
		wantList.Append(data.NewRightToken())

		areEqualList(t, gotList, wantList)
		assert.Equal(gotList.Size(), wantList.Size(), "gotList.Size() != wantList()")
	})
}

func areEqualList(t *testing.T, got, want *plugin.TokenList) {
	node1, node2 := got.Head(), want.Head()

	for node1 != nil && node2 != nil {
		token1, token2 := node1.Data(), node2.Data()

		if assert.EqualValues(t, token1.Kind(), token2.Kind(), "kind1 != kind2") {
			if token1.Kind() == data.NumToken {
				assert.EqualValues(t, *token1.Value(), *token2.Value(), "value1 != value2")
			}
		}

		node1, node2 = node1.Next(), node2.Next()
	}
}
