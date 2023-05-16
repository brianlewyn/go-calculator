package math

import (
	"fmt"
	"math"
	"strconv"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/plugin"
)

// Math returns the result of calculating the expression inside the list of tokens
func Math(list *plugin.TokenList) (float64, data.Error) {
	for calculateDeepestExpression(list) {
	}

	answer := operateNodes(list)
	if math.IsNaN(answer) {
		return 0, data.NewError(fmt.Sprint(answer), ierr.AnswerIsNaN)
	}

	return answer, nil
}

// calculateDeepestExpression returns true if it can calculate the deepest expression in parentheses
func calculateDeepestExpression(list *plugin.TokenList) bool {
	right := rightTokenNode(list.Head())
	left := leftTokenNode(right)

	connection := nodeConnection(left)
	tempList := deeperList(left, right, list)
	if tempList == nil {
		return false
	}

	result := operateNodes(tempList)
	list.Connect(connection, data.NewDecToken(result))
	return true
}

// rightTokenNode gets the first Right Token in the list from left to right
func rightTokenNode(node *plugin.TokenNode) *plugin.TokenNode {
	for temp := node; temp != nil; temp = temp.Next() {
		if temp.Token().Kind() == data.RightToken {
			return temp
		}
	}
	return nil
}

// leftTokenNode gets the first Left Token in the list from right to left
func leftTokenNode(node *plugin.TokenNode) *plugin.TokenNode {
	for temp := node; temp != nil; temp = temp.Prev() {
		if temp.Token().Kind() == data.LeftToken {
			return temp
		}
	}
	return nil
}

// nodeConnection returns the converted node as node connection
func nodeConnection(node *plugin.TokenNode) *plugin.TokenNode {
	if node == nil {
		return nil
	}
	return node.Prev()
}

// deeperList returns the deepest expression in parentheses as a list
func deeperList(left, right *plugin.TokenNode, list *plugin.TokenList) *plugin.TokenList {
	tempList := plugin.NewTokenList()

	for temp := left; temp != nil; temp = temp.Next() {
		flag := temp.Token() != left.Token() && temp.Token() != right.Token()

		if flag {
			token := temp.Token()
			tempList.Append(&token)
			list.Disconnect(temp.Prev())

		} else if temp.Token() == right.Token() {
			list.Disconnect(right.Prev())
			list.Disconnect(right)
			return tempList
		}
	}

	return nil
}

// convertFloatList converts the 'Number' node of the TokenList to a 'Float'
// and also converts the PiToken with math.Pi as a 'Float'
func convertFloatList(list *plugin.TokenList) {
	for temp := list.Head(); temp != nil; temp = temp.Next() {
		token := temp.Token()

		if token.Kind() == data.PiToken {
			temp.Update(data.NewDecToken(math.Pi))

		} else if token, ok := token.(data.Number); ok {
			float, _ := strconv.ParseFloat(token.Value(), 64)
			temp.Update(data.NewDecToken(float))
		}
	}
}

func operateNodes(list *plugin.TokenList) float64 {
	convertFloatList(list)

	doPowAndRoot(list)
	doMulDivAndMod(list)
	doAddAndSub(list)

	answer := response(list)
	list.Flush()
	return answer
}

// doPowAndRoot do powers & roots
func doPowAndRoot(list *plugin.TokenList) {
	for temp := list.Head(); temp != nil; temp = temp.Next() {
		switch temp.Token().Kind() {

		case data.PowToken:
			x := temp.Prev().Token().(data.Decimal).Value()
			y := temp.Next().Token().(data.Decimal).Value()
			list.Disconnect(temp.Prev())
			list.Disconnect(temp.Next())

			temp.Update(data.NewDecToken(math.Pow(x, y)))

		case data.RootToken:
			x := temp.Next().Token().(data.Decimal).Value()
			list.Disconnect(temp.Next())
			temp.Update(data.NewDecToken(math.Sqrt(x)))
		}
	}
}

// doMulDivAndMod do multiplication, division & module
func doMulDivAndMod(list *plugin.TokenList) {
	for temp := list.Head(); temp != nil; temp = temp.Next() {
		token := temp.Token().Kind()

		if token == data.MulToken || token == data.DivToken || token == data.ModToken {
			prev, next := temp.Prev(), temp.Next()

			x := prev.Token().(data.Decimal).Value()
			y := next.Token().(data.Decimal).Value()
			list.Disconnect(prev)
			list.Disconnect(next)

			switch token {
			case data.MulToken:
				temp.Update(data.NewDecToken(x * y))
			case data.DivToken:
				temp.Update(data.NewDecToken(x / y))
			case data.ModToken:
				temp.Update(data.NewDecToken(math.Mod(x, y)))
			}
		}
	}
}

// doAddAndSub do addition & subtraction
func doAddAndSub(list *plugin.TokenList) {
	for temp := list.Head(); temp != nil; temp = temp.Next() {
		token := temp.Token().Kind()

		if token == data.AddToken || token == data.SubToken {
			prev, next := temp.Prev(), temp.Next()

			x := prev.Token().(data.Decimal).Value()
			y := next.Token().(data.Decimal).Value()
			list.Disconnect(prev)
			list.Disconnect(next)

			switch token {
			case data.AddToken:
				temp.Update(data.NewDecToken(x + y))
			case data.SubToken:
				temp.Update(data.NewDecToken(x - y))
			}
		}
	}
}

// response returns the response of calculating the list of tokens
func response(list *plugin.TokenList) float64 {
	switch {
	case list.Head() == nil:
	case list.Head().Token() == nil:
	case list.Head().Token().Kind() != data.NumToken:
	default:
		return list.Head().Token().(data.Decimal).Value()
	}
	return 0
}
