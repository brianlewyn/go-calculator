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
		gotList, err := Tokenizer("")
		assert.ErrorIsf(t, err, ierr.EmptyField, "[err != EmptyField]: %v", err)
		assert.Nil(t, gotList, "gotList != nil")
	})

	t.Run("From an expression to a empty linked list", func(t *testing.T) {
		gotList, err := Tokenizer("  ")
		assert.ErrorIsf(t, err, ierr.EmptyField, "[err != EmptyField]: %v", err)
		assert.Nil(t, gotList, "gotList != nil")
	})

	t.Run("From an expression to a linked list", func(t *testing.T) {
		gotList, err := Tokenizer("+(0 - 1 + 2 * 3 / 4 ^ 5 % 6 + √π)(-1.234)")
		assert.Nil(t, err, "error != nil")

		wantList := toList("(0-1+2*3/4^5%6+√π)*(0-1.234)")
		areEqualList(t, gotList, wantList)
	})
}

func TestToTokenizedLinkedList(t *testing.T) {
	t.Run("From an expression with some inappropriate symbols to a list", func(t *testing.T) {
		gotList, err := toTokenizedLinkedList("12345 + hola + 12345")
		errRune := new(ierr.Rune)

		assert.ErrorAsf(t, err, &errRune, "[err != Rune]: %v", err)
		assert.Nil(t, gotList, "gotList != nil")
	})

	t.Run("From a filled expression to a list", func(t *testing.T) {
		gotList, err := toTokenizedLinkedList("(0 - 1 + 2 * 3 / 4 ^ 5 % 6 + √π) - 1.234")
		assert.Nil(t, err, "error != nil")

		wantList := toList("(0-1+2*3/4^5%6+√π)-1.234")
		areEqualList(t, gotList, wantList)

		assert.Equal(t, gotList.Size(), wantList.Size(), "g.Size != w.Size")
	})
}

func TestRebuildTokenizedLinkedList(t *testing.T) {
	t.Run("From a list to a list rebuilded", func(t *testing.T) {
		gotList, err := toTokenizedLinkedList("-(+10)(-12)(*12)")
		assert.Nil(t, err, "error != nil")

		rebuildTokenizedLinkedList(gotList)
		wantList := toList("0-(10)*(0-12)*(*12)")

		areEqualList(t, gotList, wantList)
		assert.Equal(t, gotList.Size(), wantList.Size(), "g.Size != w.Size")
	})

	t.Run("From a list to a list rebuilded (complex)(+)", func(t *testing.T) {
		gotList, err := toTokenizedLinkedList("5^-2 - 5^(2^-(1/2) * 2^-√(π-8)) - 5*-√π --4")
		assert.Nil(t, err, "error != nil")

		rebuildTokenizedLinkedList(gotList)

		wantList := toList("5^(0-2)-5^(2^(0-(1/2))*2^(0-√(π-8)))-5*(0-√π)-(0-4)")
		areEqualList(t, gotList, wantList)

		assert.Equal(t, gotList.Size(), wantList.Size(), "g.Size != w.Size")
	})

	t.Run("From a list to a list rebuilded (complex)(-)", func(t *testing.T) {
		gotList, err := toTokenizedLinkedList("5^+2 + 5^(2^+(1/2) * 2^+√(π+8)) + 5*+√π ++4")
		assert.Nil(t, err, "error != nil")

		rebuildTokenizedLinkedList(gotList)

		wantList := toList("5^2+5^(2^(1/2)*2^√(π+8))+5*√π+4")
		areEqualList(t, gotList, wantList)

		assert.Equal(t, gotList.Size(), wantList.Size(), "g.Size != w.Size")
	})

	t.Run("From a list to a list rebuilded (bugs complex: single add)", func(t *testing.T) {
		gotList, err := toTokenizedLinkedList("^+")
		assert.Nil(t, err, "error != nil")

		rebuildTokenizedLinkedList(gotList)

		wantList := toList("^+")
		areEqualList(t, gotList, wantList)

		assert.Equal(t, gotList.Size(), wantList.Size(), "g.Size != w.Size")
	})

	t.Run("From a list to a list rebuilded (bugs complex: double add)", func(t *testing.T) {
		gotList, err := toTokenizedLinkedList("^++")
		assert.Nil(t, err, "error != nil")

		rebuildTokenizedLinkedList(gotList)

		wantList := toList("^++")
		areEqualList(t, gotList, wantList)

		assert.Equal(t, gotList.Size(), wantList.Size(), "g.Size != w.Size")
	})

	t.Run("From a list to a list rebuilded (bugs complex: *-√%)", func(t *testing.T) {
		gotList, err := toTokenizedLinkedList("*-√%")
		assert.Nil(t, err, "error != nil")

		rebuildTokenizedLinkedList(gotList)

		wantList := toList("*-√%")
		areEqualList(t, gotList, wantList)

		assert.Equal(t, gotList.Size(), wantList.Size(), "g.Size != w.Size")
	})
}

// areEqualList throws t.Error if assert.Equal or assert.EqualValues finds an error
func areEqualList(t *testing.T, got, want *plugin.TokenList) {
	n1, n2 := got.Head(), want.Head()

	for n1 != nil && n2 != nil {
		t1, t2 := n1.Token(), n2.Token()
		k1, k2 := t1.Kind(), t2.Kind()

		if assert.Equalf(t, k1, k2, "k1: %c != k2: %c", data.ToRune(k1), data.ToRune(k2)) {
			if k1 == data.NumToken {
				v1 := t1.(data.Number).Value()
				v2 := t2.(data.Number).Value()
				assert.IsType(t, v1, v2, "value1: %s != value2: %s", v1, v2)
			}
		}

		n1, n2 = n1.Next(), n2.Next()
	}

	if !assert.Equal(t, got, want) {
		t.Errorf("\n\n%s\n%s\n\n", got, want)
	}
}

// toList returns the expression in a raw Tokenized Linked List
func toList(expression string) *plugin.TokenList {
	num, start, lock := new(string), new(int), new(bool)
	kind, ok := data.TokenKind(0), false
	list := plugin.NewTokenList()

	for i, r := range expression {
		if kind, ok = data.TokenKindMap[r]; ok {
			list.Append(data.NewToken(kind))
			continue
		}

		if data.IsDecimal(r) {
			if isFullNumber(expression, i, start, lock, num) {
				list.Append(data.NewNumToken(*num))
			}
			continue
		}
	}

	return list
}

// go test -bench=BenchmarkTokenizer -benchmem -count=10 -benchtime=100x > bench.txt
func BenchmarkTokenizer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Tokenizer("(0.5 + 4.5 - 1) * 10 * √(6-2) / 4^2")
	}
}
