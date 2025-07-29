package main

import (
	"errors"
	"strconv"
)

func evalPostfix(rpn []string) (int, error) {
	var stack []int

	for _, token := range rpn {
		if isNumber(token) {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		} else {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var res int

			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				res = a / b
			default:
				return 0, errors.New("unknown operator: " + token)
			}

			stack = append(stack, res)
		}
	}

	return stack[0], nil
}
