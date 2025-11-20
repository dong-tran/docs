package behavioral

import (
"fmt"
"strconv"
"strings"
)

// Interpreter Pattern - Defines a grammatical representation and an interpreter.

type Expression interface {
	Interpret() int
}

type NumberExpression struct {
	value int
}

func (n *NumberExpression) Interpret() int {
	return n.value
}

type AddExpression struct {
	left  Expression
	right Expression
}

func (a *AddExpression) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

type SubtractExpression struct {
	left  Expression
	right Expression
}

func (s *SubtractExpression) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}

type MultiplyExpression struct {
	left  Expression
	right Expression
}

func (m *MultiplyExpression) Interpret() int {
	return m.left.Interpret() * m.right.Interpret()
}

type DivideExpression struct {
	left  Expression
	right Expression
}

func (d *DivideExpression) Interpret() int {
	right := d.right.Interpret()
	if right == 0 {
		return 0
	}
	return d.left.Interpret() / right
}

func Parse(expression string) Expression {
	tokens := strings.Fields(expression)
	stack := []Expression{}
	for _, token := range tokens {
		switch token {
		case "+":
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, &AddExpression{left, right})
		case "-":
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, &SubtractExpression{left, right})
		case "*":
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, &MultiplyExpression{left, right})
		case "/":
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, &DivideExpression{left, right})
		default:
			val, _ := strconv.Atoi(token)
			stack = append(stack, &NumberExpression{val})
		}
	}
	return stack[0]
}

func DemoInterpreter() {
	fmt.Println("=== Interpreter Pattern Demo ===\n")
	expressions := []string{
		"5 3 +",       // 5 + 3 = 8
		"10 2 -",      // 10 - 2 = 8
		"4 5 *",       // 4 * 5 = 20
		"20 4 /",      // 20 / 4 = 5
		"5 3 + 2 *",   // (5 + 3) * 2 = 16
		"10 2 - 3 *",  // (10 - 2) * 3 = 24
	}
	for _, expr := range expressions {
		expression := Parse(expr)
		result := expression.Interpret()
		fmt.Printf("Expression: '%s' = %d\n", expr, result)
	}
}
