package datastructure

import (
	"strconv"
)

// expression = ((2 + 5) * (6 / 2) + 3)
// supported values: positive integer, (, ), +, -, *, /, SPACE
func EvaluateExpression(expression string) (float64, error) {
	operatorStack := NewStack()
	valueStack := NewStack()

	nextStartFrom := 0
	var nextItem item

loopThroughExpression:
	for ; ; {
		nextItem, nextStartFrom = readNextItem(expression, nextStartFrom)
		switch nextItem.t {
		case itemTypeLeftBracket:
		case itemTypeNumber:
			valueStack.Push(nextItem.value)
		case itemTypePlus:
			operatorStack.Push(itemTypePlus)
		case itemTypeMinus:
			operatorStack.Push(itemTypeMinus)
		case itemTypeDivide:
			operatorStack.Push(itemTypeDivide)
		case itemTypeMultiply:
			operatorStack.Push(itemTypeMultiply)
		case itemTypeRightBracket:
			if err := calculateOnce(operatorStack, valueStack); err != nil {
				return 0, err
			}
		default:
			break loopThroughExpression
		}
	}

	for ; !operatorStack.IsEmpty(); {
		if err := calculateOnce(operatorStack, valueStack); err != nil {
			return 0, err
		}
	}

	result, err := valueStack.Pop()
	if err != nil {
		return 0, err
	}
	return result.(float64), nil
}

func calculateOnce(operatorStack, valueStack Stack) error {
	op, err := operatorStack.Pop()
	if err != nil {
		return err
	}
	rightValue, err := valueStack.Pop()
	if err != nil {
		return err
	}
	leftValue, err := valueStack.Pop()
	if err != nil {
		return err
	}
	rightFloatValue := rightValue.(float64)
	leftFloatValue := leftValue.(float64)
	var calculatedValue float64
	switch op.(itemType) {
	case itemTypePlus:
		calculatedValue = leftFloatValue + rightFloatValue
	case itemTypeMinus:
		calculatedValue = leftFloatValue - rightFloatValue
	case itemTypeMultiply:
		calculatedValue = leftFloatValue * rightFloatValue
	case itemTypeDivide:
		calculatedValue = leftFloatValue / rightFloatValue
	}
	valueStack.Push(calculatedValue)
	return nil
}

type itemType int

const (
	itemTypeLeftBracket itemType = iota + 1
	itemTypeNumber
	itemTypeRightBracket
	itemTypePlus
	itemTypeMinus
	itemTypeDivide
	itemTypeMultiply
	itemTypeEof
)

type item struct {
	value float64
	t     itemType
}

func readNextItem(expression string, startFrom int) (v item, nextStartFrom int) {
	if startFrom >= len(expression) {
		return item{t: itemTypeEof}, -1
	}

	currentCharacter := expression[startFrom]
	switch currentCharacter {
	case '(':
		return item{t: itemTypeLeftBracket}, startFrom + 1
	case ')':
		return item{t: itemTypeRightBracket}, startFrom + 1
	case '+':
		return item{t: itemTypePlus}, startFrom + 1
	case '-':
		return item{t: itemTypeMinus}, startFrom + 1
	case '*':
		return item{t: itemTypeMultiply}, startFrom + 1
	case '/':
		return item{t: itemTypeDivide}, startFrom + 1
	case ' ':
		return readNextItem(expression, startFrom+1)
	default:
		if expression[startFrom] != '.' && (expression[startFrom] < '0' || expression[startFrom] > '9') {
			panic("not supported character:" + string(expression[startFrom]))
		}

		firstNonDigitPosition := startFrom
		for ; firstNonDigitPosition < len(expression) &&
			(expression[firstNonDigitPosition] == '.' ||
				expression[firstNonDigitPosition] >= '0' &&
					expression[firstNonDigitPosition] <= '9'); firstNonDigitPosition++ {
		}
		strToConvert := expression[startFrom:firstNonDigitPosition]
		parsedFloat, err := strconv.ParseFloat(strToConvert, 64)
		if err != nil {
			panic("not supported string:" + strToConvert)
		}
		return item{t: itemTypeNumber, value: parsedFloat}, firstNonDigitPosition
	}
}
