package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите выражение: ")

	result, err := calculate(os.Stdin)
	if err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			panic("Ошибка " + writeErr.Error())
		}

		os.Exit(1)
	}

	_, _ = fmt.Fprintln(os.Stdout, result)
}
