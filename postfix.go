package main

func postfix(tokens []string, priority map[rune]int) []string {
	var output []string

	var stack []string

	for _, token := range tokens {
		switch {
		case isNumber(token):
			output = append(output, token)

		case token == "(":
			stack = append(stack, token)

		case token == ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, pop(&stack))
			}

			if len(stack) > 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			}

		default:
			for len(stack) > 0 && stack[len(stack)-1] != "(" &&
				priority[rune(token[0])] <= priority[rune(stack[len(stack)-1][0])] {
				output = append(output, pop(&stack))
			}

			stack = append(stack, token)
		}
	}

	for len(stack) > 0 {
		output = append(output, pop(&stack))
	}

	return output
}
