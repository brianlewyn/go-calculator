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
	alwaysWrapInParentheses(list)

	for {
		left, right := obtainDeeperParentheses(list.NHead())
		if left != nil && right != nil {
			fromNumberToDecimalFrom(left, right)
			calculateExpression(list, left, right)
			deleteParentheses(list, left, right)
			continue
		}
		break
	}

	return result(list.NHead())
}

// alwaysWrapInParentheses always wrap tokenized list in parentheses
func alwaysWrapInParentheses(list *doubly.Doubly[data.Token]) {
	list.DPrepend(data.NewSymbolToken(data.LeftToken))
	list.DAppend(data.NewSymbolToken(data.RightToken))
}

// obtainDeeperParentheses gets the deepest LeftToken and RightToken nodes of the Tokenize Linked List
func obtainDeeperParentheses(head *doubly.Node[data.Token]) (left, right *doubly.Node[data.Token]) {
	right = searchFirstRightTokenNodeWith(head)
	return searchFirstLeftTokenNodeWith(right), right
}

// searchFirstRightTokenNodeWith gets the first RightToken in the list from left to right from 'head' node
func searchFirstRightTokenNodeWith(head *doubly.Node[data.Token]) *doubly.Node[data.Token] {
	for temp := head; temp != nil; temp = temp.NNext() {
		if isKind(temp, data.RightToken) {
			return temp
		}
	}
	return nil
}

// searchFirstLeftTokenNodeWith gets the first LeftToken in the list from right to left from 'right' node
func searchFirstLeftTokenNodeWith(right *doubly.Node[data.Token]) *doubly.Node[data.Token] {
	for temp := right; temp != nil; temp = temp.NPrev() {
		if isKind(temp, data.LeftToken) {
			return temp
		}
	}
	return nil
}

// fromNumberToDecimalFrom converts the following:
//   - the 'Number' nodes of the Tokenized Linked List to a 'Decimal'
//   - the PiToken with math.Pi as a 'Float'
func fromNumberToDecimalFrom(left, right *doubly.Node[data.Token]) {
	for temp := left.NNext(); temp != right; temp = temp.NNext() {

		if isKind(temp, data.PiToken) {
			temp.Update(data.NewDecimalToken(math.Pi))
			continue
		}

		if token, ok := temp.Data().(data.Number); ok {
			decimal, _ := strconv.ParseFloat(token.Value(), 64)
			temp.Update(data.NewDecimalToken(decimal))
		}

	}
}

// doPowAndRoot do powers & roots
func doPowAndRoot(list *doubly.Doubly[data.Token], left, right *doubly.Node[data.Token]) {
	for temp := left.NNext(); temp != right; temp = temp.NNext() {

		if isKind(temp, data.PowToken) {
			x, y := toDecimal(temp.NPrev()), toDecimal(temp.NNext())
			temp.Update(data.NewDecimalToken(math.Pow(x, y)))
			disconnectEnds(list, temp)
			continue
		}

		if isKind(temp, data.RootToken) {
			sqrt := math.Sqrt(toDecimal(temp.NNext()))
			temp.Update(data.NewDecimalToken(sqrt))
			list.Disconnect(temp.NNext())
		}

	}
}

// doMulDivAndMod do multiplication, division & module
func doMulDivAndMod(list *doubly.Doubly[data.Token], left, right *doubly.Node[data.Token]) {
	for temp := left.NNext(); temp != right; temp = temp.NNext() {

		if isKind(temp, data.MulToken) {
			x, y := toDecimal(temp.NPrev()), toDecimal(temp.NNext())
			temp.Update(data.NewDecimalToken(x * y))
			disconnectEnds(list, temp)
			continue
		}

		if isKind(temp, data.DivToken) {
			x, y := toDecimal(temp.NPrev()), toDecimal(temp.NNext())
			temp.Update(data.NewDecimalToken(x / y))
			disconnectEnds(list, temp)
			continue
		}

		if isKind(temp, data.ModToken) {
			x, y := toDecimal(temp.NPrev()), toDecimal(temp.NNext())
			temp.Update(data.NewDecimalToken(math.Mod(x, y)))
			disconnectEnds(list, temp)
		}

	}
}

// doAddAndSub do addition & subtraction
func doAddAndSub(list *doubly.Doubly[data.Token], left, right *doubly.Node[data.Token]) {
	for temp := left.NNext(); temp != right; temp = temp.NNext() {

		if isKind(temp, data.AddToken) {
			x, y := toDecimal(temp.NPrev()), toDecimal(temp.NNext())
			temp.Update(data.NewDecimalToken(x + y))
			disconnectEnds(list, temp)
			continue
		}

		if isKind(temp, data.SubToken) {
			x, y := toDecimal(temp.NPrev()), toDecimal(temp.NNext())
			temp.Update(data.NewDecimalToken(x - y))
			disconnectEnds(list, temp)
		}

	}
}

// isKind returns true if the node's kind is equal to the given kind, otherwise returns false
func isKind(node *doubly.Node[data.Token], kind data.TokenKind) bool {
	return node.Data().Kind() == kind
}

// toDecimal returns node's value as type float64
func toDecimal(node *doubly.Node[data.Token]) float64 {
	return node.Data().(data.Decimal).Value()
}

// disconnectEnds disconnects end nodes to current node
func disconnectEnds(list *doubly.Doubly[data.Token], node *doubly.Node[data.Token]) {
	list.Disconnect(node.NPrev())
	list.Disconnect(node.NNext())
}

// calculateExpression calculate an expression from a given range of nodes
func calculateExpression(list *doubly.Doubly[data.Token], left, right *doubly.Node[data.Token]) {
	doPowAndRoot(list, left, right)
	doMulDivAndMod(list, left, right)
	doAddAndSub(list, left, right)
}

// deleteParentheses deletes the end nodes of the current node
func deleteParentheses(list *doubly.Doubly[data.Token], left, right *doubly.Node[data.Token]) {
	list.Disconnect(left)
	list.Disconnect(right)
}

func result(head *doubly.Node[data.Token]) (float64, error) {
	if res64 := toDecimal(head); !math.IsNaN(res64) {
		return res64, nil
	}
	return 0, ierr.ResultIsNaN
}
