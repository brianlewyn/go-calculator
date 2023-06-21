package math

import (
	"math"
	"strconv"

	"github.com/brianlewyn/go-calculator/ierr"
	"github.com/brianlewyn/go-calculator/internal/data"
	"github.com/brianlewyn/go-calculator/internal/doubly"
)

// Math returns the result of calculating the expression inside the list of tokens
func Math(list *doubly.Doubly) (float64, error) {
	alwaysWrapInParentheses(list)

	for {
		left, right := obtainDeeperParentheses(list.Head())
		if left != nil && right != nil {
			fromNumberToDecimalFrom(left, right)
			calculateExpression(list, left, right)
			deleteParentheses(list, left, right)
			continue
		}
		break
	}

	return result(list.Head())
}

// alwaysWrapInParentheses always wrap tokenized list in parentheses
func alwaysWrapInParentheses(list *doubly.Doubly) {
	list.PushFront(data.NewSymbolToken(data.LeftToken))
	list.PushBack(data.NewSymbolToken(data.RightToken))
}

// obtainDeeperParentheses gets the deepest LeftToken and RightToken nodes of the Tokenize Linked List
func obtainDeeperParentheses(head *doubly.Node) (left, right *doubly.Node) {
	right = searchFirstRightTokenNodeWith(head)
	return searchFirstLeftTokenNodeWith(right), right
}

// searchFirstRightTokenNodeWith gets the first RightToken in the list from left to right from 'head' node
func searchFirstRightTokenNodeWith(head *doubly.Node) *doubly.Node {
	for temp := head; temp != nil; temp = temp.Next() {
		if isKind(temp, data.RightToken) {
			return temp
		}
	}
	return nil
}

// searchFirstLeftTokenNodeWith gets the first LeftToken in the list from right to left from 'right' node
func searchFirstLeftTokenNodeWith(right *doubly.Node) *doubly.Node {
	for temp := right; temp != nil; temp = temp.Prev() {
		if isKind(temp, data.LeftToken) {
			return temp
		}
	}
	return nil
}

// fromNumberToDecimalFrom converts the following:
//   - the 'Number' nodes of the Tokenized Linked List to a 'Decimal'
//   - the PiToken with math.Pi as a 'Float'
func fromNumberToDecimalFrom(left, right *doubly.Node) {
	for temp := left.Next(); temp != right; temp = temp.Next() {

		if isKind(temp, data.PiToken) {
			temp.Update(data.NewDecimalToken(math.Pi))
			continue
		}

		if token, ok := temp.Token().(data.Number); ok {
			decimal, _ := strconv.ParseFloat(token.Value(), 64)
			temp.Update(data.NewDecimalToken(decimal))
		}
	}
}

// doPowAndRoot do powers & roots
func doPowAndRoot(list *doubly.Doubly, left, right *doubly.Node) {
	for temp := left.Next(); temp != right; temp = temp.Next() {

		if isKind(temp, data.PowToken) {
			x, y := toDecimal(temp.Prev()), toDecimal(temp.Next())
			temp.Update(data.NewDecimalToken(math.Pow(x, y)))
			removeNodeEnds(list, temp)
			continue
		}

		if isKind(temp, data.RootToken) {
			sqrt := math.Sqrt(toDecimal(temp.Next()))
			temp.Update(data.NewDecimalToken(sqrt))
			list.RemoveNode(temp.Next())
		}
	}
}

// doMulDivAndMod do multiplication, division & module
func doMulDivAndMod(list *doubly.Doubly, left, right *doubly.Node) {
	for temp := left.Next(); temp != right; temp = temp.Next() {

		if isKind(temp, data.MulToken) {
			x, y := toDecimal(temp.Prev()), toDecimal(temp.Next())
			temp.Update(data.NewDecimalToken(x * y))
			removeNodeEnds(list, temp)
			continue
		}

		if isKind(temp, data.DivToken) {
			x, y := toDecimal(temp.Prev()), toDecimal(temp.Next())
			temp.Update(data.NewDecimalToken(x / y))
			removeNodeEnds(list, temp)
			continue
		}

		if isKind(temp, data.ModToken) {
			x, y := toDecimal(temp.Prev()), toDecimal(temp.Next())
			temp.Update(data.NewDecimalToken(math.Mod(x, y)))
			removeNodeEnds(list, temp)
		}
	}
}

// doAddAndSub do addition & subtraction
func doAddAndSub(list *doubly.Doubly, left, right *doubly.Node) {
	for temp := left.Next(); temp != right; temp = temp.Next() {

		if isKind(temp, data.AddToken) {
			x, y := toDecimal(temp.Prev()), toDecimal(temp.Next())
			temp.Update(data.NewDecimalToken(x + y))
			removeNodeEnds(list, temp)
			continue
		}

		if isKind(temp, data.SubToken) {
			x, y := toDecimal(temp.Prev()), toDecimal(temp.Next())
			temp.Update(data.NewDecimalToken(x - y))
			removeNodeEnds(list, temp)
		}
	}
}

// isKind returns true if the node's kind is equal to the given kind, otherwise returns false
func isKind(node *doubly.Node, kind data.TokenKind) bool {
	return node.Token().Kind() == kind
}

// toDecimal returns node's value as type float64
func toDecimal(node *doubly.Node) float64 {
	return node.Token().(data.Decimal).Value()
}

// removeNodeEnds RemoveNodes end nodes to current node
func removeNodeEnds(list *doubly.Doubly, node *doubly.Node) {
	list.RemoveNode(node.Prev())
	list.RemoveNode(node.Next())
}

// calculateExpression calculate an expression from a given range of nodes
func calculateExpression(list *doubly.Doubly, left, right *doubly.Node) {
	doPowAndRoot(list, left, right)
	doMulDivAndMod(list, left, right)
	doAddAndSub(list, left, right)
}

// deleteParentheses deletes the end nodes of the current node
func deleteParentheses(list *doubly.Doubly, left, right *doubly.Node) {
	list.RemoveNode(left)
	list.RemoveNode(right)
}

func result(head *doubly.Node) (float64, error) {
	if res64 := toDecimal(head); !math.IsNaN(res64) {
		return res64, nil
	}
	return 0, ierr.ResultIsNaN
}
