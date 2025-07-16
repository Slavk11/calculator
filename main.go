package main

import (
	"calculator/stack"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите выражение: ")

	result, err := calculate(os.Stdin)
	if err != nil {
		fmt.Println("Ошибка", err)
		os.Exit(1)
	}

	s := stack.New()

	s.Push("one")

	fmt.Println(result)
}
