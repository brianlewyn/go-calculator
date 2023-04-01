package tool

import (
	"strconv"

	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
)

// Converter converts the 'Number' data.Node from TokenList to 'Float'
func Converter(list *plugin.TokenList) {
	slice := *list.Unmarshal()

	for i, token := range slice {
		if token.Kind() == data.NumToken {
			value := token.(data.Number).Value()
			float, _ := strconv.ParseFloat(value, 64)
			slice[i] = *data.NewFloatToken(float)
		}
	}

	list.Marshal(&slice)
}
