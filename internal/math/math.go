package math

import (
	"math"
	"strconv"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-linked-list/v2/doubly"
)

// Math returns the result of calculating the expression inside the list of tokens
func Math(list *doubly.Doubly[data.Token]) (float64, error) {
	for calculateDeepestExpression(list) {
	}

	answer := operateNodes(list)
	if math.IsNaN(answer) {
		return 0, ierr.ResultIsNaN
	}

	return answer, nil
}

// calculateDeepestExpression returns true if it can calculate the deepest expression in parentheses
func calculateDeepestExpression(list *doubly.Doubly[data.Token]) bool {
	right := rightTokenNode(list.NHead())
	left := leftTokenNode(right)

	connection := nodeConnection(left)
	tempList := deeperList(left, right, list)
	if tempList == nil {
		return false
	}

	result := operateNodes(tempList)

	if connection == nil {
		list.DPrepend(data.NewDecimalToken(result))
	} else {
		list.Connect(connection, doubly.NewNode(data.NewDecimalToken(result)))
	}
	return true
}

// rightTokenNode gets the first Right Token in the list from left to right
func rightTokenNode(node *doubly.Node[data.Token]) *doubly.Node[data.Token] {
	for temp := node; temp != nil; temp = temp.NNext() {
		if temp.Data().Kind() == data.RightToken {
			return temp
		}
	}
	return nil
}

// leftTokenNode gets the first Left Token in the list from right to left
func leftTokenNode(node *doubly.Node[data.Token]) *doubly.Node[data.Token] {
	for temp := node; temp != nil; temp = temp.NPrev() {
		if temp.Data().Kind() == data.LeftToken {
			return temp
		}
	}
	return nil
}

// nodeConnection returns the converted node as node connection
func nodeConnection(node *doubly.Node[data.Token]) *doubly.Node[data.Token] {
	if node == nil {
		return nil
	}
	return node.NPrev()
}

// deeperList returns the deepest expression in parentheses as a list
func deeperList(left, right *doubly.Node[data.Token], list *doubly.Doubly[data.Token]) *doubly.Doubly[data.Token] {
	tempList := doubly.NewDoubly[data.Token]()

	for temp := left; temp != nil; temp = temp.NNext() {
		flag := temp.Data() != left.Data() && temp.Data() != right.Data()

		if flag {
			tempList.DAppend(temp.Data())
			list.Disconnect(temp.NPrev())

		} else if temp.Data() == right.Data() {
			list.Disconnect(right.NPrev())
			list.Disconnect(right)
			return tempList
		}
	}

	return nil
}

// convertFloatList converts the 'Number' node of the TokenList to a 'Float'
// and also converts the PiToken with math.Pi as a 'Float'
func convertFloatList(list *doubly.Doubly[data.Token]) {
	for temp := list.NHead(); temp != nil; temp = temp.NNext() {
		token := temp.Data()

		if token.Kind() == data.PiToken {
			temp.Update(data.NewDecimalToken(math.Pi))

		} else if token, ok := token.(data.Number); ok {
			float, _ := strconv.ParseFloat(token.Value(), 64)
			temp.Update(data.NewDecimalToken(float))
		}
	}
}

func operateNodes(list *doubly.Doubly[data.Token]) float64 {
	convertFloatList(list)

	doPowAndRoot(list)
	doMulDivAndMod(list)
	doAddAndSub(list)

	answer := response(list)
	list.Flush(false)
	return answer
}

// doPowAndRoot do powers & roots
func doPowAndRoot(list *doubly.Doubly[data.Token]) {
	for temp := list.NHead(); temp != nil; temp = temp.NNext() {
		switch temp.Data().Kind() {

		case data.PowToken:
			x := temp.DPrev().(data.Decimal).Value()
			y := temp.DNext().(data.Decimal).Value()
			list.Disconnect(temp.NPrev())
			list.Disconnect(temp.NNext())

			temp.Update(data.NewDecimalToken(math.Pow(x, y)))

		case data.RootToken:
			x := temp.DNext().(data.Decimal).Value()
			list.Disconnect(temp.NNext())
			temp.Update(data.NewDecimalToken(math.Sqrt(x)))
		}
	}
}

// doMulDivAndMod do multiplication, division & module
func doMulDivAndMod(list *doubly.Doubly[data.Token]) {
	for temp := list.NHead(); temp != nil; temp = temp.NNext() {
		token := temp.Data().Kind()

		if token == data.MulToken || token == data.DivToken || token == data.ModToken {
			prev, next := temp.NPrev(), temp.NNext()

			x := prev.Data().(data.Decimal).Value()
			y := next.Data().(data.Decimal).Value()
			list.Disconnect(prev)
			list.Disconnect(next)

			switch token {
			case data.MulToken:
				temp.Update(data.NewDecimalToken(x * y))
			case data.DivToken:
				temp.Update(data.NewDecimalToken(x / y))
			case data.ModToken:
				temp.Update(data.NewDecimalToken(math.Mod(x, y)))
			}
		}
	}
}

// doAddAndSub do addition & subtraction
func doAddAndSub(list *doubly.Doubly[data.Token]) {
	for temp := list.NHead(); temp != nil; temp = temp.NNext() {
		token := temp.Data().Kind()

		if token == data.AddToken || token == data.SubToken {
			prev, next := temp.NPrev(), temp.NNext()

			x := prev.Data().(data.Decimal).Value()
			y := next.Data().(data.Decimal).Value()
			list.Disconnect(prev)
			list.Disconnect(next)

			switch token {
			case data.AddToken:
				temp.Update(data.NewDecimalToken(x + y))
			case data.SubToken:
				temp.Update(data.NewDecimalToken(x - y))
			}
		}
	}
}

// response returns the response of calculating the list of tokens
func response(list *doubly.Doubly[data.Token]) float64 {
	switch {
	case list.NHead() == nil:
	case list.DHead() == nil:
	case list.DHead().Kind() != data.NumToken:
	default:
		return list.DHead().(data.Decimal).Value()
	}
	return 0
}
