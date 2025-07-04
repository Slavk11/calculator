package main

import "strings"

func tokenize(expr string) []string {
	var tokens []string

	var buffer strings.Builder

	for _, ch := range expr {
		if ch >= '0' && ch <= '9' {
			buffer.WriteRune(ch)
		} else {
			if buffer.Len() > 0 {
				tokens = append(tokens, buffer.String())
				buffer.Reset()
			}

			tokens = append(tokens, string(ch))
		}
	}

	if buffer.Len() > 0 {
		tokens = append(tokens, buffer.String())
	}

	return tokens
}
