package tokenize

import (
	"testing"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	t.Run("From an empty expression to a linked list", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := Tokenizer("")
		assert.ErrorIsf(err.Bug(), ierr.EmptyField, "ierr.EmptyField != %v", err.Bug())
		assert.Nil(gotList, "gotList != nil")
	})

	t.Run("From an expression to a empty linked list", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := Tokenizer("  ")
		assert.ErrorIsf(err.Bug(), ierr.EmptyField, "ierr.EmptyField != %v", err.Bug())
		assert.Nil(gotList, "gotList != nil")
	})

	t.Run("From an expression to a linked list", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := Tokenizer("+(0 - 1 + 2 * 3 / 4 ^ 5 % 6 + √π)(-1.234)")
		assert.Nil(err, "Tokenizer(data) error != nil")

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

func TestInitLinkedList(t *testing.T) {
	t.Run("From an expression with some inappropriate symbols to a list", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := initLinkedList("12345 + hola + 12345")

		errRune := new(ierr.Rune)
		assert.ErrorAsf(err, &errRune, "&ierr.Rune != %v", err)
		assert.Nil(gotList, "gotList != nil")
	})

	t.Run("From a filled expression to a list", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := initLinkedList("(0 - 1 + 2 * 3 / 4 ^ 5 % 6 + √π) - 1.234")
		assert.NoError(err, "initLinkedList(expression) error != nil")

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
	})
}

func TestInitRebuild(t *testing.T) {
	t.Run("From an empty list to a list rebuilded", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := initLinkedList("   ")
		assert.NoError(err, "initLinkedList(expression) error != nil")

		err = initRebuild(gotList)
		assert.ErrorIsf(err, ierr.EmptyField, "ierr.EmptyField != %v", err)
	})

	t.Run("From a list to a list rebuilded", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := initLinkedList("-(+10)(-12)(*12)")
		assert.NoError(err, "initLinkedList(expression) error != nil")

		err = initRebuild(gotList)
		assert.NoError(err, "tokenizer.rebuild(gotList) error != nil")

		wantList := plugin.NewTokenList()
		// 0-(10)
		wantList.Append(data.NewNumToken("0"))
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewLeftToken())
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

	t.Run("From a list to a list rebuilded (complex)(+)", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := initLinkedList("5^-2 - 5^(2^-(1/2) * 2^-√(π-8)) - 5*-√π --4")
		assert.NoError(err, "initLinkedList(expression) error != nil")

		err = initRebuild(gotList)
		assert.NoError(err, "tokenizer.rebuild(gotList) error != nil")

		wantList := plugin.NewTokenList()
		// 5^(0-2)
		wantList.Append(data.NewNumToken("5"))
		wantList.Append(data.NewPowToken())
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("0"))
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewNumToken("2"))
		wantList.Append(data.NewRightToken())
		// -5^
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewNumToken("5"))
		wantList.Append(data.NewPowToken())
		// (2^(0-(1/2))*2^
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("2"))
		wantList.Append(data.NewPowToken())
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("0"))
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("1"))
		wantList.Append(data.NewDivToken())
		wantList.Append(data.NewNumToken("2"))
		wantList.Append(data.NewRightToken())
		wantList.Append(data.NewRightToken())
		wantList.Append(data.NewMulToken())
		wantList.Append(data.NewNumToken("2"))
		wantList.Append(data.NewPowToken())
		// (0-√(π-8)))
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("0"))
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewRootToken())
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewPiToken())
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewNumToken("8"))
		wantList.Append(data.NewRightToken())
		wantList.Append(data.NewRightToken())
		wantList.Append(data.NewRightToken())
		// -5*-√π
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewNumToken("5"))
		wantList.Append(data.NewMulToken())
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("0"))
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewRootToken())
		wantList.Append(data.NewPiToken())
		wantList.Append(data.NewRightToken())
		// --4
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("0"))
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewNumToken("4"))
		wantList.Append(data.NewRightToken())

		areEqualList(t, gotList, wantList)
		assert.Equal(gotList.Size(), wantList.Size(), "gotList.Size() != wantList.Size()")
	})

	t.Run("From a list to a list rebuilded (complex)(-)", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := initLinkedList("5^+2 + 5^(2^+(1/2) * 2^+√(π+8)) + 5*+√π ++4")
		assert.NoError(err, "initLinkedList(expression) error != nil")

		err = initRebuild(gotList)
		assert.NoError(err, "tokenizer.rebuild(gotList) error != nil")

		wantList := plugin.NewTokenList()
		// 5^2
		wantList.Append(data.NewNumToken("5"))
		wantList.Append(data.NewPowToken())
		wantList.Append(data.NewNumToken("2"))
		// +5^
		wantList.Append(data.NewAddToken())
		wantList.Append(data.NewNumToken("5"))
		wantList.Append(data.NewPowToken())
		// (2^(1/2)*2^
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("2"))
		wantList.Append(data.NewPowToken())
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewNumToken("1"))
		wantList.Append(data.NewDivToken())
		wantList.Append(data.NewNumToken("2"))
		wantList.Append(data.NewRightToken())
		wantList.Append(data.NewMulToken())
		wantList.Append(data.NewNumToken("2"))
		wantList.Append(data.NewPowToken())
		// √(π+8))
		wantList.Append(data.NewRootToken())
		wantList.Append(data.NewLeftToken())
		wantList.Append(data.NewPiToken())
		wantList.Append(data.NewAddToken())
		wantList.Append(data.NewNumToken("8"))
		wantList.Append(data.NewRightToken())
		wantList.Append(data.NewRightToken())
		// +5*√π
		wantList.Append(data.NewAddToken())
		wantList.Append(data.NewNumToken("5"))
		wantList.Append(data.NewMulToken())
		wantList.Append(data.NewRootToken())
		wantList.Append(data.NewPiToken())
		// +4
		wantList.Append(data.NewAddToken())
		wantList.Append(data.NewNumToken("4"))

		areEqualList(t, gotList, wantList)
		assert.Equal(gotList.Size(), wantList.Size(), "gotList.Size() != wantList.Size()")
	})

	t.Run("From a list to a list rebuilded (bugs complex: single add)", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := initLinkedList("^+")
		assert.NoError(err, "initLinkedList(expression) error != nil")

		err = initRebuild(gotList)
		assert.NoError(err, "tokenizer.rebuild(gotList) error != nil")

		wantList := plugin.NewTokenList()
		// ^+
		wantList.Append(data.NewPowToken())
		wantList.Append(data.NewAddToken())

		areEqualList(t, gotList, wantList)
		assert.Equal(gotList.Size(), wantList.Size(), "gotList.Size() != wantList.Size()")
	})

	t.Run("From a list to a list rebuilded (bugs complex: double add)", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := initLinkedList("^++")
		assert.NoError(err, "initLinkedList(expression) error != nil")

		err = initRebuild(gotList)
		assert.NoError(err, "tokenizer.rebuild(gotList) error != nil")

		wantList := plugin.NewTokenList()
		// ^++
		wantList.Append(data.NewPowToken())
		wantList.Append(data.NewAddToken())
		wantList.Append(data.NewAddToken())

		areEqualList(t, gotList, wantList)
		assert.Equal(gotList.Size(), wantList.Size(), "gotList.Size() != wantList.Size()")
	})

	t.Run("From a list to a list rebuilded (bugs complex: √{%,*,+,-,/,^})", func(t *testing.T) {
		assert := assert.New(t)

		gotList, err := initLinkedList("*-√%")
		assert.NoError(err, "initLinkedList(expression) error != nil")

		err = initRebuild(gotList)
		assert.NoError(err, "tokenizer.rebuild(gotList) error != nil")

		wantList := plugin.NewTokenList()
		// *-√%
		wantList.Append(data.NewMulToken())
		wantList.Append(data.NewSubToken())
		wantList.Append(data.NewRootToken())
		wantList.Append(data.NewModToken())

		areEqualList(t, gotList, wantList)
		assert.Equal(gotList.Size(), wantList.Size(), "gotList.Size() != wantList.Size()")
	})
}

func areEqualList(t *testing.T, got, want *plugin.TokenList) {
	node1, node2 := got.Head(), want.Head()

	for node1 != nil && node2 != nil {
		token1, token2 := node1.Token(), node2.Token()

		if assert.Equalf(t, token1.Kind(), token2.Kind(), "kind1: %c != kind2: %c",
			data.ChangeToRune(token1.Kind()), data.ChangeToRune(token2.Kind())) {
			if token1.Kind() == data.NumToken {
				value1 := token1.(data.Number).Value()
				value2 := token2.(data.Number).Value()
				assert.EqualValues(t, value1, value2, "value1 != value2")
			}
		}

		node1, node2 = node1.Next(), node2.Next()
	}
}

// go test -bench=BenchmarkTokenizer -benchmem -count=10 -benchtime=100x > bench.txt
func BenchmarkTokenizer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Tokenizer("(0.5 + 4.5 - 1) * 10 * √(6-2) / 4^2")
	}
}
