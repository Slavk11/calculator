package main

import (
	"fmt"
	"unicode"
)

func tokenize(expr string) ([]string, error) {
	var tokens []string

	var buffer []rune

	prevToken := ""

	for i, ch := range expr {
		switch {
		case ch >= '0' && ch <= '9':
			buffer = append(buffer, ch)

		case ch == '+' || ch == '*' || ch == '/' || ch == ')':
			if len(buffer) > 0 {
				tokens = append(tokens, string(buffer))
				buffer = buffer[:0]
			}
			tokens = append(tokens, string(ch))
			prevToken = string(ch)

		case ch == '(':
			if len(buffer) > 0 {
				tokens = append(tokens, string(buffer))
				buffer = buffer[:0]
			} else {
				tokens = append(tokens, string(ch))
				prevToken = string(ch)
			}

		case ch == '-':
			if i == 0 || prevToken == "(" || prevToken == "+" || prevToken == "-" || prevToken == "*" || prevToken == "/" {
				buffer = append(buffer, ch)
			} else {
				if len(buffer) > 0 {
					tokens = append(tokens, string(buffer))
					buffer = buffer[:0]
				} else {
					tokens = append(tokens, string(ch))
					prevToken = string(ch)
				}
			}

		case unicode.IsSpace(ch):
			continue
		default:
			return nil, fmt.Errorf("Ошибка! Символ %s является недопустимым!\n", string(ch))
		}
	}

	if len(buffer) > 0 {
		tokens = append(tokens, string(buffer))
	}

	return tokens, nil
}
