package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func calculate(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	expr := scanner.Text()
	expr = strings.Join(strings.Fields(expr), "")
	tokens, err := tokenize(expr)
	if err != nil {
		return 0, fmt.Errorf("недопустимый символ в выражении %w", err)
	}
	priority := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'(': 0,
	}

	rpn, err := postfix(tokens, priority)
	if err != nil {
		return 0, fmt.Errorf("ошибка при работе с постфиксом %w", err)
	}

	result, err := evalPostfix(rpn)
	if err != nil {
		return 0, fmt.Errorf("ошибка при работе с постфиксом %w", err)
	}

	return result, nil
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
