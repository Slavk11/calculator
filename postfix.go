package main

import "calculator/stack"

func postfix(tokens []string, priority map[rune]int) ([]string, error) {
	var output []string
	s := stack.New()

	for _, token := range tokens {
		switch {
		case isNumber(token):
			output = append(output, token)

		case token == "(":
			s.Push(token)

		case token == ")":
			for {
				val, _ := s.Pop()
				if val == "(" {
					break
				}
				output = append(output, val)
			}

		default:
			for {
				val, err := s.Pop()
				if err != nil {
					break
				}
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

	for {
		val, err := s.Pop()
		if err != nil {
			break
		}
		output = append(output, val)
	}

	return output, nil
}
