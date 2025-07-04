package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate() {
	fmt.Println("Введите выражение: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expr := scanner.Text()
	expr = strings.ReplaceAll(expr, " ", "")

	priority := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'(': 0,
	}

	tokens := tokenize(expr)
	fmt.Println("Tokens: ", tokens)

	rpn := postfix(tokens, priority)
	fmt.Println("Постфикс", rpn)

	result := evalPostfix(rpn)
	fmt.Println("Результат:", result)
}

func pop(stack *[]string) string {
	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	return top
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
