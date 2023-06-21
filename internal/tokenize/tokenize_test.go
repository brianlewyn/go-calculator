package tokenize

import (
	"fmt"
	"strings"
	"testing"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/doubly"
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
func areEqualList(t *testing.T, g, w *doubly.Doubly) {
	for n1, n2 := g.Head(), w.Head(); n1 != nil && n2 != nil; n1, n2 = n1.Next(), n2.Next() {
		k1, k2 := n1.Token().Kind(), n2.Token().Kind()
		assert.Equal(t, k1, k2, "\n\nk1 != k2\n\n")
	}

	if !assert.Equal(t, g, w) {
		t.Errorf("\n\n%s\n%s\n\n", toString(g), toString(w))
	}
}

// toList returns the expression in a raw Tokenized Linked List
func toList(expression string) *doubly.Doubly {
	k, list := 0, doubly.New()

	for i, r := range expression {
		if data.IsDecimal(r) {
			if i >= k {
				num := getFullNumber(expression[i:])
				list.PushBack(data.NewNumberToken(num))
				k = i + len(num)
			}
			continue
		}

		if kind, ok := data.TokenKindMap[r]; ok {
			list.PushBack(data.NewSymbolToken(kind))
		}
	}

	return list
}

// toString converts a Doubly Linked List to string
func toString(list *doubly.Doubly) string {
	if list.IsEmpty() {
		return "list <nil>"
	}

	var b strings.Builder
	for temp := list.Head(); temp != nil; temp = temp.Next() {
		fmt.Fprintf(&b, "%c", data.RuneMap[temp.Token().Kind()])
	}

	return b.String()
}

// go test -bench=BenchmarkTokenizer -benchmem -count=10 -benchtime=100x >> bench.txt
func BenchmarkTokenizer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Tokenizer("(0.5 + 4.5 - 1) * 10 * √(6-2) / 4^2")
	}
}
