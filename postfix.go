package main

import "calculator/stack"

func postfix(tokens []string, priority map[rune]int) []string {
	var output []string
	s := stack.New()

	for _, token := range tokens {
		switch {
		case isNumber(token):
			output = append(output, token)

		case token == "(":
			s.Push(token)

		case token == ")":
			for !s.Empty() {
				val := s.Pop()
				if val == "(" {
					break
				}
				output = append(output, val)
			}

		default:
			for !s.Empty() {
				val := s.Pop()
				if val == "(" {
					s.Push(val)
					break
				}
				if priority[rune(token[0])] <= priority[rune(val[0])] {
					output = append(output, val)
				} else {
					s.Push(val)
					break
				}
			}
			s.Push(token)
		}
	}

	for !s.Empty() {
		output = append(output, s.Pop())
	}

	return output
}
