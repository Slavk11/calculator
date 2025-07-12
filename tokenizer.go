package main

import (
	"fmt"
	"unicode"
)

func tokenize(expr string) []string {
	var tokens []string

	var buffer []rune

	for _, ch := range expr {
		switch {
		case ch >= '0' && ch <= '9':
			buffer = append(buffer, ch)
		case ch == '+' || ch == '-' || ch == '*' ||
			ch == '/' || ch == '(' || ch == ')':
			if len(buffer) > 0 {
				tokens = append(tokens, string(buffer))
				buffer = buffer[:0]
			}
			tokens = append(tokens, string(ch))
		case unicode.IsSpace(ch):
			continue
		default:
			fmt.Printf("Ошибка! Символ %s является недопустимым!\n", string(ch))
			continue
		}
	}

	if len(buffer) > 0 {
		tokens = append(tokens, string(buffer))
	}

	return tokens
}
