package tool

import (
	"strconv"

	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
)

// Converter converts the 'Number' data.Node from TokenList to 'Float'
func Converter(list *plugin.TokenList) {
	for temp := list.Head(); temp != nil; temp = temp.Next() {
		token := temp.Token()

		if token.Kind() == data.NumToken {
			value := temp.Token().(data.Number).Value()
			float, _ := strconv.ParseFloat(value, 64)
			temp.Update(data.NewFloatToken(float))
		}
	}
}
