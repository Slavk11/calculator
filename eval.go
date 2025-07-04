package main

import "strconv"

func evalPostfix(rpn []string) int {
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
				res = a / b
			}

			stack = append(stack, res)
		}
	}

	return stack[0]
}
